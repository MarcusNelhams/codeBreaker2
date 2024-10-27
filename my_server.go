package main
import (
        "context"
        "net"
        "log"
        "github.com/MarcusNelhams/part3/auth"
        "google.golang.org/grpc"
)

type myAuthServer struct {
        auth.UnimplementedAuthServiceServer
}

func (s myAuthServer) Authenticate(ctx context.Context, in *auth.AuthRequest) (*auth.AuthResponse, error) {
        log.Printf("received authRequest of\n\tusername: %s\n\tpassword: %s\n", in.Username, in.Password)
        if in.Username == "admin" && in.Password == "password123" {
                log.Println("Success!")
                return &auth.AuthResponse {
                        Success: true,
                }, nil
        }
        log.Println("Failure!")
        return &auth.AuthResponse {
                Success: false,
        }, nil
        
}

func (s myAuthServer) Logout(ctx context.Context, in *auth.LogoutRequest) (*auth.LogoutResponse, error) {
        log.Printf("received LogoutRequest with token: %s", in.Token)
        return &auth.LogoutResponse {
                Success: true,
        }, nil
}

func (s myAuthServer) Ping(ctx context.Context, in *auth.PingRequest) (*auth.PingResponse, error) {
        log.Printf("Received Ping int: %d", in.Ping)
        return &auth.PingResponse {
                Response: 1,
        }, nil
}

func (s myAuthServer) RefreshToken(ctx context.Context, in *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
        log.Printf("received RefreshTokenRequest")
        return &auth.RefreshTokenResponse {
                Token: "1",
        }, nil
}

func (s myAuthServer) RegisterOTPSeed(ctx context.Context, in *auth.RegisterOTPSeedRequest) (*auth.RegisterOTPSeedResponse, error) {
        log.Printf("received RegisterOPTSeed")
        return &auth.RegisterOTPSeedResponse {
                Success: true,
        }, nil
}

func (s myAuthServer) VerifyOTP(ctx context.Context, in *auth.VerifyOTPRequest) (*auth.VerifyOTPResponse, error) {
        log.Printf("received VerifyOTPRequest")
        return &auth.VerifyOTPResponse {
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
        service := &myAuthServer{}

        auth.RegisterAuthServiceServer(serverRegistrar, service)

        err = serverRegistrar.Serve(lis)
        if err != nil {
                log.Fatalf("cannot serve server: %s", err)
        }
}
