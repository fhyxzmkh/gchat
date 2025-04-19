package api

import (
	"errors"
	"gchat/internal/service"
	"gchat/pkg/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "gchat/internal/domain/dto/request/account"
import "github.com/sirupsen/logrus"

type UserInfoHandler struct {
	svc *service.UserInfoService
}

func (u *UserInfoHandler) Login(ctx *gin.Context) {
	var loginReq account.LoginRequest

	// 1. 参数绑定
	if err := ctx.BindJSON(&loginReq); err != nil {
		logrus.WithContext(ctx).Errorf("Bind failed: %v", err)
		ctx.JSON(http.StatusInternalServerError, "Invalid request format")
		return
	}

	// 2. 调用 Service
	user, err := u.svc.Login(ctx, loginReq)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidUserOrPassword):
			ctx.JSON(http.StatusUnauthorized, "账号或密码错误")
		default:
			logrus.WithContext(ctx).Errorf("Login failed: %v", err)
			ctx.JSON(http.StatusInternalServerError, constants.SYSTEM_ERROR)
		}
		return
	}

	// 3. 成功响应
	ctx.JSON(http.StatusOK, user)
}
