package data

import (
	"context"

	"kratos-restart/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data   *Data
	logger *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	r.logger.WithContext(ctx).Infof("username: %s,999", u.Username)
	return nil, nil
}
