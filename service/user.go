package service

import (
	"errors"
	"gin-vue-admin-STL/entity"
	"gin-vue-admin-STL/global"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type UserService struct{}

func (u *UserService) Register(entityUser entity.User) (userInter entity.User, err error) {
	var user entity.User
	if !errors.Is(global.GVA_DB.Where("username = ?", entityUser.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	entityUser.UUID = uuid.Must(uuid.NewV4())
	err = global.GVA_DB.Create(&u).Error
	return entityUser, err
}
