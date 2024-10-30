package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Username string
	Password string
	Nickname string
}

type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}
