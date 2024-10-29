package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-restart/api/user/v1"
)

type UserRepo interface {
	// v1.UserClient
}

type UserUseCase struct {
	repo UserRepo
	uc   v1.UserClient
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, uc v1.UserClient, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, uc: uc, log: log.NewHelper(log.With(logger, "module", "biz/user"))}
}
