package users

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"mnc-users/apps/databases/repositories"
	"mnc-users/apps/users/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

type UserResponseDTO struct {
	UserID      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	CreatedDate     string `json:"created_date"`
}

type UpdateProfileResponse struct {
	UserID      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	UpdatedDate string `json:"updated_date"`
}

func (s *UserService) Register(dto dto.UserRegisterDTO) (*UserResponseDTO, error) {
	existingUser, _ := s.Repo.FindByPhoneNumber(dto.PhoneNumber)
	if existingUser != nil {
		return nil, errors.New("Phone Number already registered")
	}

	userID, createdAt, err := s.Repo.Create(dto.FirstName, dto.LastName, dto.PhoneNumber, dto.Address, dto.PIN)
	if err != nil {
		return nil, err
	}

	response := &UserResponseDTO{
		UserID:      userID,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		PhoneNumber: dto.PhoneNumber,
		Address:     dto.Address,
		CreatedDate: createdAt,
	}
	return response, nil
}

func generateToken(userID string, secret string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(expiry).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (s *UserService) Login(phoneNumber, pin string) (string, string, error) {
	user, err := s.Repo.FindByPhoneNumber(phoneNumber)
	if err != nil {
		return "", "", errors.New("Phone number not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PIN), []byte(pin)); err != nil {
		return "", "", errors.New("Phone number and PIN doesn’t match")
	}

	accessSecret := os.Getenv("JWT_SECRET")
	refreshSecret := os.Getenv("JWT_REFRESH_SECRET")

	accessToken, err := generateToken(user.ID.String(), accessSecret, time.Hour*24)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateToken(user.ID.String(), refreshSecret, time.Hour*24*7)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *UserService) UpdateProfile(userID string, req dto.UpdateProfileRequest) (*UpdateProfileResponse, error) {
	user, err := s.Repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("User not found")
	}

	if req.FirstName != nil {
		user.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		user.LastName = *req.LastName
	}
	if req.Address != nil {
		user.Address = *req.Address
	}
	updatedAt := time.Now()

	err = s.Repo.UpdateUserProfile(userID, user.FirstName, user.LastName, user.Address, updatedAt)
	if err != nil {
		return nil, err
	}

	return &UpdateProfileResponse{
		UserID:      user.ID.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Address:     user.Address,
		UpdatedDate: updatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
