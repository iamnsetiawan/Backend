package user

import (
	"context"

	"github.com/TrinityKnights/Backend/internal/domain/model"
)

type UserService interface {
	Register(ctx context.Context, request *model.RegisterRequest) (*model.UserResponse, error)
	Login(ctx context.Context, request *model.LoginRequest) (*model.TokenResponse, error)
	Profile(ctx context.Context) (*model.UserResponse, error)
	Update(ctx context.Context, request *model.UpdateRequest) (*model.UserResponse, error)
	RefreshToken(ctx context.Context, request *model.RefreshTokenRequest) (*model.TokenResponse, error)
	RequestReset(ctx context.Context, request *model.RequestReset) error
	ResetPassword(ctx context.Context, request *model.ResetPassword) error
}
