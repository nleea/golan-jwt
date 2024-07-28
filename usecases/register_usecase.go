package usecases

import (
	"context"
	"learn/jwt/domain"
)

type registerUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository domain.UserRepository) domain.RegisterUsecase {
	return &registerUseCase{
		userRepository: userRepository,
	}
}

// Create implements domain.RegisterUsecase.
func (r *registerUseCase) Create(c context.Context, user *domain.User) error {
	return r.userRepository.Create(c, user)
}

// CreateAccessToken implements domain.RegisterUsecase.
func (r *registerUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	panic("unimplemented")
}

// CreateRefreshToken implements domain.RegisterUsecase.
func (r *registerUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	panic("unimplemented")
}

// GetUserByEmail implements domain.RegisterUsecase.
func (r *registerUseCase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	return r.userRepository.GetByEmail(c, email)
}
