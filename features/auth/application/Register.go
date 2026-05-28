package application

import (
	"errors"
	"fastcomp_api/core/security"
	"fastcomp_api/features/auth/domain"
	"fastcomp_api/features/auth/domain/entities"
	"time"
)

type Register struct {
	userRepo domain.UserRepository
}

func NewRegister(userRepo domain.UserRepository) *Register {
	return &Register{userRepo: userRepo}
}

func (r *Register) Execute(user *entities.User) (*entities.User, error) {
	existing, _ := r.userRepo.GetByEmail(user.Email)
	if existing != nil {
		return nil, errors.New("el email ya está registrado")
	}
	if user.Password == nil || *user.Password == "" {
		return nil, errors.New("la contraseña es requerida")
	}

	hashed, err := security.HashPassword(*user.Password)
	if err != nil {
		return nil, errors.New("error al procesar la contraseña")
	}

	user.Password = &hashed
	user.RegistrationDate = time.Now()
	if user.RoleID == 0 {
		user.RoleID = 2
	}

	return r.userRepo.Save(user)
}
