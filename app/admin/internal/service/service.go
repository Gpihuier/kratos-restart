package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "kratos-restart/api/admin/v1"
	"kratos-restart/app/admin/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)

type AdminService struct {
	v1.UnimplementedAdminServer

	uc     *biz.UserUseCase
	logger *log.Helper
}

func NewAdminService(uc *biz.UserUseCase, logger log.Logger) *AdminService {
	return &AdminService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "service/admin")),
	}
}

func (s *AdminService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	s.logger.WithContext(ctx).Infow("login", "username", in.Username, "password", in.Password)
	tk := time.Now().Format(time.RFC1123)
	return &v1.LoginReply{
		Token:   tk,
		Message: fmt.Sprintf("Hello: %s, Welcome to you!", in.Username),
	}, nil
}

func (s *AdminService) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	cu, err := s.uc.CreateUser(ctx, &biz.User{
		Username: in.Username,
		Password: in.Password,
		Nickname: in.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Message: fmt.Sprintf("Hello: %s, Welcome to you!", cu.Nickname),
	}, err
}
