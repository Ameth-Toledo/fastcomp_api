package dto

type CreateUserRequest struct {
	FirstName    string  `json:"firstName"    binding:"required"`
	LastName     string  `json:"lastName"     binding:"required"`
	BusinessName string  `json:"businessName" binding:"required"`
	Email        string  `json:"email"        binding:"required,email"`
	Password     string  `json:"password"     binding:"required,min=8"`
	Phone        string  `json:"phone"        binding:"required,len=10"`
	Website      *string `json:"website,omitempty"`
	ProfilePhoto *string `json:"profilePhoto,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	FirstName    string  `json:"firstName"    binding:"required"`
	LastName     string  `json:"lastName"     binding:"required"`
	BusinessName string  `json:"businessName" binding:"required"`
	Email        string  `json:"email"        binding:"required,email"`
	Phone        string  `json:"phone"        binding:"required,len=10"`
	Website      *string `json:"website,omitempty"`
	ProfilePhoto *string `json:"profilePhoto,omitempty"`
}
