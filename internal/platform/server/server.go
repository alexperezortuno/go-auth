package server

import (
	"context"
	"fmt"
	"github.com/alexperezortuno/go-auth/common/environment"
	"github.com/alexperezortuno/go-auth/internal/platform/server/handler/auth"
	"github.com/alexperezortuno/go-auth/internal/platform/server/handler/health"
	"github.com/alexperezortuno/go-auth/internal/platform/server/middleware/logging"
	"github.com/alexperezortuno/go-auth/internal/platform/server/middleware/recovery"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/data_base"
	"github.com/alexperezortuno/go-auth/internal/platform/storage/redis"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpAddr        string
	engine          *gin.Engine
	shutdownTimeout time.Duration
}

func New(ctx context.Context, params environment.ServerValues) (context.Context, Server) {
	srv := Server{
		engine:          gin.New(),
		httpAddr:        fmt.Sprintf("%s:%d", params.Host, params.Port),
		shutdownTimeout: params.ShutdownTimeout,
	}

	log.Println(fmt.Sprintf("Check app in %s:%d/%s/%s", params.Host, params.Port, params.Context, "health"))
	srv.registerRoutes(params.Context)
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context, params environment.ServerValues) error {
	log.Println("Server running on", s.httpAddr)
	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	redis.InitializeStore()
	data_base.Init(params)
	defer data_base.CloseConnection()
	data_base.Migrate()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) registerRoutes(context string) {
	s.engine.Use(logging.Middleware(), gin.Logger(), recovery.Middleware())

	s.engine.GET(fmt.Sprintf("/%s/%s", context, "/health"), health.CheckHandler())
	s.engine.POST(fmt.Sprintf("/%s/%s", context, "/login"), auth.LoginHandler())
	s.engine.POST(fmt.Sprintf("/%s/%s", context, "/create"), auth.CreateUserHandler())
	s.engine.POST(fmt.Sprintf("/%s/%s", context, "/verify"), auth.ValidateTokenHandler())
	s.engine.POST(fmt.Sprintf("/%s/%s", context, "/refresh"), auth.RefreshTokenHandler())
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
