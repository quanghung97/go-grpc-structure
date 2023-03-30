package authenticate

import (
	"context"

	PkgMiddleWare "github.com/bav-demo/micro1/pkg/gen_token"
	pbAuthen "github.com/bav-demo/micro1/proto/authenticate/go_pb"
)

func Login(ctx context.Context, in *pbAuthen.LoginRequest) (*pbAuthen.LoginResponse, error) {
	token, err := PkgMiddleWare.GenerateToken(in.Mnv, 2)
	if err != nil {
		return nil, err
	}
	return &pbAuthen.LoginResponse{
		Token: token,
	}, nil
}
