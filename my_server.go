package main
import (
        "context"
        "net"
        "log"
        "github.com/MarcusNelhams/part3/auth_service"
        "google.golang.org/grpc"
)

type authServiceClient struct {
        auth_service.UnimplementedAuthServiceServer
}

func (s authServiceClient) Authenticate(ctx context.Context, in *auth_service.AuthRequest) (*auth_service.AuthResponse, error) {
        log.Printf("received authRequest of\n\tusername: %s\n\tpassword: %s\n", in.Username, in.Password)
        if in.Username == "admin" && in.Password == "password123" {
                log.Println("Success!")
                return &auth_service.AuthResponse {
                        Success: true,
                }, nil
        }
        log.Println("Failure!")
        return &auth_service.AuthResponse {
                Success: false,
        }, nil
        
}

func (s authServiceClient) Logout(ctx context.Context, in *auth_service.LogoutRequest) (*auth_service.LogoutResponse, error) {
        log.Printf("received LogoutRequest with token: %s", in.Token)
        return &auth_service.LogoutResponse {
                Success: true,
        }, nil
}

func (s authServiceClient) Ping(ctx context.Context, in *auth_service.PingRequest) (*auth_service.PingResponse, error) {
        log.Printf("Received Ping int: %d", in.Ping)
        return &auth_service.PingResponse {
                Response: 1,
        }, nil
}

func (s authServiceClient) RefreshToken(ctx context.Context, in *auth_service.RefreshTokenRequest) (*auth_service.RefreshTokenResponse, error) {
        log.Printf("received RefreshTokenRequest")
        return &auth_service.RefreshTokenResponse {
                Token: "1",
        }, nil
}

func (s authServiceClient) RegisterOTPSeed(ctx context.Context, in *auth_service.RegisterOTPSeedRequest) (*auth_service.RegisterOTPSeedResponse, error) {
        log.Printf("received RegisterOPTSeed")
        return &auth_service.RegisterOTPSeedResponse {
                Success: true,
        }, nil
}

func (s authServiceClient) VerifyOTP(ctx context.Context, in *auth_service.VerifyOTPRequest) (*auth_service.VerifyOTPResponse, error) {
        log.Printf("received VerifyOTPRequest")
        return &auth_service.VerifyOTPResponse {
                Success: true,
                Token: "1",
        }, nil
}

func main() {
        lis, err := net.Listen("tcp", ":50052")
        if err != nil {
                log.Fatalf("cannot create listener: %s", err)
        }

        serverRegistrar := grpc.NewServer()
        service := &authServiceClient{}

        auth_service.RegisterAuthServiceServer(serverRegistrar, service)

        err = serverRegistrar.Serve(lis)
        if err != nil {
                log.Fatalf("cannot serve server: %s", err)
        }
}
