package service

import (
	"errors"
	"time"

	"github.com/B6137151/GDZ-Commerce/internal/domain/model"
	"github.com/B6137151/GDZ-Commerce/internal/domain/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	LoginAdmin(email, password string) (string, string, error) // Return both access and refresh tokens
	LoginUser(email, password string) (string, string, error)  // Return both access and refresh tokens
	RegisterAdmin(email, password string) error
	RegisterUser(email, password string) error
	RefreshToken(refreshToken string) (string, error) // Refresh access token using refresh token
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

// Admin login with both access and refresh token generation
func (s *authService) LoginAdmin(email, password string) (string, string, error) {
	admin, err := s.repo.FindAdminByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := generateJWTToken(admin.ID.String(), admin.Role, 15*time.Minute) // Short-lived access token
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateJWTToken(admin.ID.String(), admin.Role, 7*24*time.Hour) // Long-lived refresh token
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// User login with both access and refresh token generation
func (s *authService) LoginUser(email, password string) (string, string, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, err := generateJWTToken(user.ID.String(), user.Role, 15*time.Minute) // Short-lived access token
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateJWTToken(user.ID.String(), user.Role, 7*24*time.Hour) // Long-lived refresh token
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Register new admin
func (s *authService) RegisterAdmin(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	admin := &model.Admin{
		Email:    email,
		Password: string(hashedPassword),
		Role:     "admin",
	}
	return s.repo.CreateAdmin(admin)
}

// Register new user
func (s *authService) RegisterUser(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user",
	}
	return s.repo.CreateUser(user)
}

// Generate JWT token with expiration duration
func generateJWTToken(userID, role string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(expiration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key"))
}

// Handle refresh token logic to issue new access token
func (s *authService) RefreshToken(refreshToken string) (string, error) {
	// Parse the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	// Generate a new access token
	newAccessToken, err := generateJWTToken(userID, role, 15*time.Minute)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
