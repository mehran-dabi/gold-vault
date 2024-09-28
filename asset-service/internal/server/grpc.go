package server

import (
	"net"

	"goldvault/asset-service/internal/config"
	grpchandler "goldvault/asset-service/internal/interfaces/grpc"
	"goldvault/asset-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func NewGRPCListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", ":"+config.ServiceConfig.Server.Ports.GRPC)
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func NewGRPCServer(assetHandler *grpchandler.AssetPriceGRPCHandler) *grpc.Server {
	grpcServer := grpc.NewServer()
	// Register Grpc services
	proto.RegisterAssetServiceServer(grpcServer, assetHandler)

	// Register gRPC health service
	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthSrv)

	// Set the service status
	healthSrv.SetServingStatus(config.ServiceConfig.Name, grpc_health_v1.HealthCheckResponse_SERVING)

	return grpcServer

}
