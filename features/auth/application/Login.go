package application

import (
	"errors"
	"fastcomp_api/core/security"
	"fastcomp_api/features/auth/domain"
	"fastcomp_api/features/auth/domain/entities"
	"strings"
)

type Login struct {
	userRepo domain.UserRepository
}

func NewLogin(userRepo domain.UserRepository) *Login {
	return &Login{userRepo: userRepo}
}

func (l *Login) Execute(email, password string) (*entities.User, error) {
	email = strings.TrimSpace(email)

	user, err := l.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("error al buscar usuario")
	}
	if user == nil {
		return nil, errors.New("credenciales inválidas")
	}
	if user.Password == nil {
		return nil, errors.New("usuario registrado con OAuth - usa login social")
	}
	if !security.CheckPassword(*user.Password, password) {
		return nil, errors.New("credenciales inválidas")
	}

	return user, nil
}
