package repository

import (
	"authentication/domain/entity"
	"context"
)

type User interface {
	FindUser(ctx context.Context, id uint) (entity.User, error)

	DeleteUser(ctx context.Context, id uint) error
}
