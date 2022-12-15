package bootstrap

import (
	"context"
	"github.com/alexperezortuno/go-auth/common/environment"
	"github.com/alexperezortuno/go-auth/internal/platform/server"
)

var params = environment.Server()

func Run() error {
	ctx, srv := server.New(context.Background(), params)
	return srv.Run(ctx, params)
}
