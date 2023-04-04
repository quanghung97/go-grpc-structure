package routing

import (
	"context"

	AuthenticateController "github.com/bav-demo/internal/controllers/authenticate"
	pbAuthen "github.com/bav-demo/proto/authenticate/go_pb"
)

func (s *Route) GetTokenDemo(ctx context.Context, in *pbAuthen.LoginRequest) (*pbAuthen.LoginResponse, error) {

	return AuthenticateController.Login(ctx, in)
}
