package service

import (
	"context"
	"errors"
	"gchat/internal/domain/dto/request/account"
	"gchat/internal/domain/dto/response"
	"gchat/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrInvalidUserOrPassword = errors.New("账号或密码错误")

type UserInfoService struct {
	repo *repository.UserInfoRepository
}

func (svc *UserInfoService) Login(ctx context.Context, loginReq account.LoginRequest) (response.LoginResponse, error) {
	telephone := loginReq.Telephone
	password := loginReq.Password

	u, err := svc.repo.FindByTelephone(ctx, telephone)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.LoginResponse{}, ErrInvalidUserOrPassword
	}

	if err != nil {
		return response.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// DEBUG 打日志
		return response.LoginResponse{}, ErrInvalidUserOrPassword
	}
	return response.LoginResponse{
		Uuid:      u.Uuid,
		Nickname:  u.Nickname,
		Telephone: u.Telephone,
		Avatar:    u.Avatar,
		Email:     u.Email,
		Gender:    u.Gender,
		Birthday:  u.Birthday,
		Signature: u.Signature,
		CreatedAt: u.CreatedAt,
		IsAdmin:   u.IsAdmin,
		Status:    u.Status,
	}, nil
}
