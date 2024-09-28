package server

import (
	"net"

	"goldvault/wallet-service/internal/config"
	grpchandler "goldvault/wallet-service/internal/interfaces/grpc"
	"goldvault/wallet-service/proto"

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

func NewGRPCServer(walletHandler *grpchandler.WalletGRPCHandler, assetHandler *grpchandler.AssetGRPCHandler) *grpc.Server {
	grpcServer := grpc.NewServer()
	// Register gRPC services
	proto.RegisterWalletServiceServer(grpcServer, walletHandler)
	proto.RegisterAssetServiceServer(grpcServer, assetHandler)

	// Register gRPC health service
	healthSrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthSrv)

	// Set the service status
	healthSrv.SetServingStatus(config.ServiceConfig.Name, grpc_health_v1.HealthCheckResponse_SERVING)

	return grpcServer

}
