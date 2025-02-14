package dto

type UserRegisterDTO struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	PIN         string `json:"pin"`
}

type UserLoginDTO struct {
	PhoneNumber string `json:"phone_number"`
	PIN         string `json:"pin"`
}

type UpdateProfileRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Address   *string `json:"address,omitempty"`
}
