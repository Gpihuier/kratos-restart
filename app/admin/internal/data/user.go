package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-restart/app/admin/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
