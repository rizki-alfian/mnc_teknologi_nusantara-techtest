package dto

type UserResponseDTO struct {
    UserID      string `json:"user_id"`
    FirstName   string `json:"first_name"`
    LastName    string `json:"last_name"`
    PhoneNumber string `json:"phone_number"`
    Address     string `json:"address"`
    CreatedDate string `json:"created_date"`
}

type UpdateProfileResponse struct {
	UserID      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Address     string `json:"address"`
	UpdatedDate string `json:"updated_date"`
}