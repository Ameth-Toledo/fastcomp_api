package application

import (
	"errors"
	"fastcomp_api/features/auth/domain"
	"fastcomp_api/features/auth/domain/entities"
)

type GetUserByID struct {
	userRepo domain.UserRepository
}

func NewGetUserByID(userRepo domain.UserRepository) *GetUserByID {
	return &GetUserByID{userRepo: userRepo}
}

func (g *GetUserByID) Execute(id int) (*entities.User, error) {
	user, err := g.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}
	return user, nil
}
