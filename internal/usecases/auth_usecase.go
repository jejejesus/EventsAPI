package usecases

import (
	"context"
	"errors"
	"time"

	"EventsAPI/internal/config"
	"EventsAPI/internal/domain/entities"
	"EventsAPI/internal/domain/repositories"
	"EventsAPI/pkg/utils"

	"gorm.io/gorm"
)

type AuthUseCase struct {
	userRepo repositories.UserRepository
	config   *config.Config
}

func NewAuthUseCase(userRepo repositories.UserRepository, config *config.Config) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
		config:   config,
	}
}

func (uc *AuthUseCase) Register(ctx context.Context, req *entities.UserRequest) (*entities.UserResponse, error) {
	// Check if user exists
	existingUser, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &entities.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return &entities.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, req *entities.LoginRequest) (string, *entities.UserResponse, error) {
	// Get user by email
	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, errors.New("invalid credentials")
		}
		return "", nil, err
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	// Generate JWT
	expiration, _ := time.ParseDuration(uc.config.JWT.Expiration)
	token, err := utils.GenerateJWT(user.ID, user.Email, uc.config.JWT.Secret, expiration)
	if err != nil {
		return "", nil, err
	}

	userResponse := &entities.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
	}

	return token, userResponse, nil
}
