package service

import (
	"authentication/domain/entity"
	"authentication/domain/repository"
	"context"
)

type UserService struct {
	repo repository.User
}

type CreateUserDto struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

func (u *UserService) FindUser(ctx context.Context, id uint) (entity.User, error) {
	var user entity.User
	user, err := u.repo.FindUser(ctx, id)
	if err != nil {
		return user, err
	}
	return user, nil
}
