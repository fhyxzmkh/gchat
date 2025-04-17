package api

import (
	"gchat/pkg/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "gchat/internal/domain/dto/request/account"
import "github.com/sirupsen/logrus"

func Login(ctx *gin.Context) {
	var loginReq account.LoginRequest

	if err := ctx.BindJSON(&loginReq); err != nil {
		logrus.Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})

		return
	}
}
