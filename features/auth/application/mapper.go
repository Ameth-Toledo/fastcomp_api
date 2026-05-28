package application

import (
	"fastcomp_api/features/auth/domain/dto"
	"fastcomp_api/features/auth/domain/entities"
)

func ToUserResponse(u *entities.User) dto.UserResponse {
	return dto.UserResponse{
		ID:               u.ID,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		BusinessName:     u.BusinessName,
		Email:            u.Email,
		Phone:            u.Phone,
		Website:          u.Website,
		ProfilePhoto:     u.ProfilePhoto,
		RegistrationDate: u.RegistrationDate,
		RoleID:           u.RoleID,
	}
}
