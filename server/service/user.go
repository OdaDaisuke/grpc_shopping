package service

import (
	"context"
	"github.com/OdaDaisuke/grpc_shopping/pb/user"
)

type UserService struct {
}

func (s *UserService) Signup(ctx context.Context, in *user.SignupMessage) (*user.SignupResponse, error) {
	// TODO
	return nil, nil
}

func (s *UserService) Authenticate(ctx context.Context, in *user.AuthenticateMessage) (*user.AuthenticateResponse, error) {
	// TODO
	return nil, nil
}

func (s *UserService) Delete(ctx context.Context, in *user.DeleteMessage) (*user.DeleteResponse, error) {
	// TODO
	return nil, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, in *user.UpdateProfileMessage) (*user.UpdateProfileResponse, error) {
	// TODO
	return nil, nil
}