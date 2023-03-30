package routing

import (
	"context"

	AuthenticateContrroler "github.com/bav-demo/micro1/internal/controllers/authenticate"
	pbAuthen "github.com/bav-demo/micro1/proto/authenticate/go_pb"
)

func (s *Route) GetTokenDemo(ctx context.Context, in *pbAuthen.LoginRequest) (*pbAuthen.LoginResponse, error) {

	return AuthenticateContrroler.Login(ctx, in)
}
