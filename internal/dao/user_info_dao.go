package dao

import (
	"context"
	"gchat/internal/domain/model"
	"gorm.io/gorm"
)

type UserInfoDao struct {
	db *gorm.DB
}

func (dao *UserInfoDao) FindByTelephone(ctx context.Context, telephone string) (model.UserInfo, error) {
	var user model.UserInfo
	err := dao.db.WithContext(ctx).Where("telephone = ?", telephone).First(&user).Error
	return user, err
}
