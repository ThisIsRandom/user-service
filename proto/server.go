package proto

import "context"

type Service struct {
	UnimplementedTokenServer
}

func (s *Service) CreateToken(ctx context.Context, in *CreateTokenRequest) (*CreateTokenResponse, error) {
	return &CreateTokenResponse{
		Token: "tester",
	}, nil
}
