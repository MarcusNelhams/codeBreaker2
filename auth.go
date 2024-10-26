package main
import (
	"context"
	"net"
	"log"
	"github.com/MarcusNelhams/part3_go/auth"
	"google.golang.org/grpc"
)

type myAuthServiceServer struct {
	auth.UnimplementedAuthServiceServer
}

func (s myAuthServiceServer) Create(context.Context, *auth.PingRequest) (*auth.PingResponse, err) {
	return &auth.PingResponse {
		Message: "Pong"
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	
	serverRegistrar := grpc.NewServer()
	service := &myAuthServiceServer{}

	auth.RegisterAuthServiceServer(serverRegistrar, service)

	serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("cannot serve server: %s", err)
	}
}
