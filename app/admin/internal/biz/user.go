package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-restart/api/user/v1"
)

type User struct {
	Username string
	Password string
	Nickname string
}

type UserRepo interface {
	CreateUser(ctx context.Context, in *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	uc   v1.UserClient
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, uc v1.UserClient, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, uc: uc, log: log.NewHelper(log.With(logger, "module", "biz/user"))}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, in *User) (*User, error) {
	return uc.repo.CreateUser(ctx, in)
}
