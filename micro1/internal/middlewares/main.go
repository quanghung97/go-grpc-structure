package middlewares

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bav-demo/internal/models/authenticate"
	PkgMiddleWare "github.com/bav-demo/pkg/gen_token"
	pb "github.com/bav-demo/proto/society/go_pb"
	initConnect "github.com/bav-demo/server/init"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var listLimit = map[string]bool{
	pb.SocietyService_GetAllPostRPC_FullMethodName: true,
}

func Limit(ctx context.Context) bool {
	tokenInfo := ctx.Value(authenticate.TokenAes{}).(*authenticate.TokenAes)
	limitCount, err := connection.Redis.Rdb.Get(ctx, tokenInfo.UserId).Int()
	if err != nil {
		connection.Redis.Rdb.Set(ctx, tokenInfo.UserId, 10, time.Duration(10*time.Second))
	}

	if limitCount <= 0 {
		return true
	}
	ttl, err := connection.Redis.Rdb.TTL(ctx, tokenInfo.UserId).Result()
	if err != nil || ttl == time.Duration(-1*time.Nanosecond) {
		fmt.Println(err)
		return false
	}

	connection.Redis.Rdb.Set(ctx, tokenInfo.UserId, limitCount-1, ttl)
	return false
}

var connection = initConnect.GetInstance()

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	fmt.Println("11111")
	// fmt.Println(ctx)

	fmt.Println('1', info.FullMethod)
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

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if strings.Contains(info.FullMethod, "AuthenticationService") || !listLimit[info.FullMethod] {
		return handler(ctx, req)
	}

	if Limit(ctx) {
		return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleware, please retry later.", info.FullMethod)
	}
	return handler(ctx, req)
}
