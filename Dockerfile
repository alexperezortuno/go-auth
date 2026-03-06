FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /go-auth ./cmd/api

FROM alpine:3.22

# Crear usuario no root con UID/GID específicos
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata && \
    # Establecer permisos seguros para el directorio
    chown -R appuser:appgroup /app && \
    chmod 750 /app

COPY --from=builder --chown=appuser:appgroup /go-auth /app/go-auth

# Permisos específicos: dueño puede leer/ejecutar, grupo solo ejecutar, otros ninguno
RUN chmod 750 /app/go-auth

# Cambiar al usuario no root
USER appuser

CMD ["/app/go-auth"]
