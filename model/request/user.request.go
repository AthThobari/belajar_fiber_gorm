package request


type UserCreateRequest struct {
	Name      string `json:"name" validate:"required"` 
	Email     string `json:"email" validate:"required,email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Password  string `json:"password" validate:"required,min=4"`
}
type UserUpdateRequest struct {
	Name      string `json:"name"` 
	Address   string `json:"address"`
	Phone     string `json:"phone"`
}

type UserUpdateEmailRequest struct {
		Email     string `json:"email" validate:"required"`
}