package dao

import (
	"errors"
	"gchat/internal/domain/dto/request/account"
	"gchat/internal/domain/model"
	"gchat/pkg/constants"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserInfoDao struct {
	db *gorm.DB
}

func (dao *UserInfoDao) GetUserInfoByTelephone(loginReq account.LoginRequest) (string, *model.UserInfo, int) {
	var user model.UserInfo
	res := dao.db.First(&user, "telephone = ?", loginReq.Telephone)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			message := "用户不存在，请注册"
			logrus.Error(message)
			return message, nil, -2
		}
		logrus.Error(res.Error.Error())
		return constants.SYSTEM_ERROR, nil, -1
	}
	return "", &user, 0
}
