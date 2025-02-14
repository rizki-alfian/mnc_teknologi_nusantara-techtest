package repositories

import (
	"mnc-users/apps/databases/models"
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(firstName, lastName, phoneNumber, address, pin string) (string, string, error) {

	hashedPin, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	user := models.Users{
		ID:          uuid.New(),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Address:     address,
		PIN:         string(hashedPin),
		CreatedAt:   time.Now(),
	}

	err = r.db.Create(&user).Error
	if err != nil {
		return "", "", err
	}

	return user.ID.String(), user.CreatedAt.Format("2006-01-02 15:04:05"), nil
}

func (r *UserRepository) FindByPhoneNumber(phone string) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("phone_number = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(userID string) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUserProfile(userID, firstName, lastName, address string, updatedAt time.Time) error {
	return r.db.Model(&models.Users{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"first_name":   firstName,
			"last_name":    lastName,
			"address":      address,
			"updated_at":   updatedAt,
		}).Error
}