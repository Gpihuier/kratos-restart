package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-restart/api/user/v1"
	"kratos-restart/app/user/internal/biz"
)

type UserService struct {
	v1.UnimplementedUserServer

	uc     *biz.UserUseCase
	logger *log.Helper
}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "service/admin")),
	}
}

func (s *UserService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserReply, error) {

	return nil, nil
	//u, err := s.uc.CreateUser(ctx, &biz.User{
	//	Username: in.Username,
	//	Password: in.Password,
	//	Nickname: in.Nickname,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//return &v1.CreateUserReply{Message: "Congratulations on your successful registration: " + u.Username}, nil
}
