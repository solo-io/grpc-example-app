package main

import (
	"context"
	"fmt"
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/grpc-example-app/api/store"
	"github.com/solo-io/grpc-example-app/pkg/data"
	"github.com/solo-io/grpc-example-app/pkg/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	if err := cmd().Execute(); err != nil {
		log.Fatal(err)
	}

}

const (
	bookService   = "books"
	recordService = "records"
)

func cmd() *cobra.Command {
	var service string
	var port int
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start the dev-portal web server / kube controller",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := contextutils.WithLogger(context.Background(), service+"-store-grpc")
			return start(ctx, service, port)
		},
	}
	cmd.Flags().StringVar(&service, "service", "", "specify either books or records, determines which microservice to run as")
	cmd.Flags().IntVar(&port, "port", 8080, "grpc listener port")

	return cmd
}

func start(ctx context.Context, service string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.NewNop()),
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				log.Printf("%v", info.FullMethod)
				return handler(srv, ss)
			},
		)))

	switch service {
	case bookService:
		log.Printf("starting books service")
		store.RegisterBooksServer(grpcServer, server.NewBookServer(data.Books))
	case recordService:
		log.Printf("starting records service")
		store.RegisterRecordsServer(grpcServer, server.NewRecordServer(data.Records))
	default:
		return fmt.Errorf("--service flag must equal either %v or %v", bookService, recordService)
	}

	log.Printf("listening on %v", port)
	reflection.Register(grpcServer)
	return grpcServer.Serve(lis)
}
