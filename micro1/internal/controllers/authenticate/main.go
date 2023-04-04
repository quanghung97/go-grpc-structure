package authenticate

import (
	"context"
	"fmt"

	PkgMiddleWare "github.com/bav-demo/pkg/gen_token"
	pbAuthen "github.com/bav-demo/proto/authenticate/go_pb"
	initConnect "github.com/bav-demo/server/init"
)

var connection = initConnect.GetInstance()
var limit = 10

func Login(ctx context.Context, in *pbAuthen.LoginRequest) (*pbAuthen.LoginResponse, error) {
	fmt.Println("login---authen")
	token, err := PkgMiddleWare.GenerateToken(in.Mnv, 2)

	// connection.Redis.Rdb.Set(ctx, in.Mnv, limit, time.Duration(10*time.Second))

	if err != nil {
		return nil, err
	}
	return &pbAuthen.LoginResponse{
		Token: token,
	}, nil
}
