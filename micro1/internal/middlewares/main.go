package middlewares

import (
	"context"
	"strings"

	"github.com/bav-demo/micro1/internal/models/authenticate"
	PkgMiddleWare "github.com/bav-demo/micro1/pkg/gen_token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// fmt.Println(ok)
	// fmt.Println(ctx)

	if strings.Contains(info.FullMethod, "AuthenticationService") {
		return handler(ctx, req)
	} else {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "No token found")
		}
		token := md["authorization"]
		if len(token) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "No token found")
		}
		data, err := PkgMiddleWare.VerifyToken(token[0])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
		}
		ctx = context.WithValue(ctx, authenticate.TokenAes{}, data)
		return handler(ctx, req)
	}
	// token, ok := md["authorization"]
	// fmt.Println(md)
	// fmt.Println(token)

}
