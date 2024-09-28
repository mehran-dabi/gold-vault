package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"goldvault/asset-service/docs"
	"goldvault/asset-service/internal/config"
	"goldvault/asset-service/internal/server/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Server struct {
	External   *gin.Engine
	Internal   *gin.Engine
	healthFunc func(ctx *gin.Context)

	grpcServer   *grpc.Server
	grpcListener net.Listener
}

func NewServer(grpcServer *grpc.Server, grpcListener net.Listener) *Server {
	if !config.ServiceConfig.Server.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	s := &Server{
		External:     gin.Default(),
		Internal:     gin.Default(),
		grpcServer:   grpcServer,
		grpcListener: grpcListener,
	}
	s.External.Use(otelgin.Middleware(config.ServiceConfig.Name))
	s.External.Use(middlewares.CORS())
	s.Internal.Use(otelgin.Middleware(config.ServiceConfig.Name))
	if config.Environment(config.ServiceConfig.Env) != config.PROD {
		s.External.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		s.setDoc()
	}
	return s
}

func (s *Server) WithMiddlewares(middlewares ...gin.HandlerFunc) *Server {
	for _, mw := range middlewares {
		s.External.Use(mw)
	}
	return s
}

func (s *Server) SetHealthFunc(f func() error) *Server {
	s.healthFunc = func(ctx *gin.Context) {
		if err := f(); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
	return s
}

func (s *Server) SetupRoutes() {
	s.External.GET("/health", s.healthFunc)
	s.Internal.GET("/health", s.healthFunc)
}

func (s *Server) Run(port string) {
	err := s.External.Run(":" + port)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run web server")
	}
}

func (s *Server) setDoc() {
	docs.SwaggerInfo.Title = "Asset Service APIs"
	docs.SwaggerInfo.Version = "1.0"
}

func (s *Server) RunAsync(port string) {
	go s.Run(port)
}

func Run(lc fx.Lifecycle, s *Server) {
	external := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.ServiceConfig.Server.Ports.External),
		Handler: s.External,
	}
	internal := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.ServiceConfig.Server.Ports.Internal),
		Handler: s.Internal,
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("shutting down the server ...")
			err := external.Shutdown(ctx)
			if err != nil {
				return err
			}
			return internal.Shutdown(ctx)
		},
		OnStart: func(ctx context.Context) error {
			log.Info().Msg("running server ...")
			go func() {
				if err := external.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Err(err).Msg("failed to run external web server")
				}
			}()
			go func() {
				if err := internal.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Err(err).Msg("failed to run internal web server")
				}
			}()
			if s.grpcServer != nil && s.grpcListener != nil {
				go func() {
					if err := s.grpcServer.Serve(s.grpcListener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
						log.Error().Err(err).Msg("failed to run gRPC server")
					}
				}()
			}
			return nil
		}},
	)
}
