package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-restart/api/user/v1"
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

func (u *userRepo) CreateUser(ctx context.Context, in *biz.User) (*biz.User, error) {
	reply, err := u.data.uc.CreateUser(ctx, &v1.CreateUserRequest{
		Username: in.Username,
		Password: in.Password,
		Nickname: in.Nickname,
	})
	if err != nil {
		return nil, err
	}

	u.log.WithContext(ctx).Infof("Hello: %s", reply.Message)

	return &biz.User{
		Username: in.Username,
		Password: in.Password,
		Nickname: in.Nickname,
	}, nil
}
