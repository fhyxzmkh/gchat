package repository

import (
	"context"
	"gchat/internal/dao"
	"gchat/internal/domain/model"
)

type UserInfoRepository struct {
	dao *dao.UserInfoDao
}

func (r *UserInfoRepository) FindByTelephone(ctx context.Context, telephone string) (model.UserInfo, error) {
	u, err := r.dao.FindByTelephone(ctx, telephone)

	if err != nil {
		return model.UserInfo{}, err
	}

	return u, nil
}
