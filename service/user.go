package service

import (
	"errors"
	"fmt"
	"gin-vue-admin-STL/entity"
	"gin-vue-admin-STL/global"
	"gin-vue-admin-STL/model/request"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"time"
)

type UserService struct{}

func (u *UserService) Register(entityUser entity.User) (userInter entity.User, err error) {
	var user entity.User
	if !errors.Is(global.GVA_DB.Where("username = ?", entityUser.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	entityUser.UUID = uuid.Must(uuid.NewV4())
	err = global.GVA_DB.Create(&entityUser).Error
	return entityUser, err
}

func (u *UserService) ChangePassWord(entityUser *entity.User, newPassWord string) (userInter *entity.User, err error) {
	var user entity.User
	if errors.Is(global.GVA_DB.Where("phone = ?", entityUser.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名或密码不正确")
	}
	user.Password = newPassWord
	err = global.GVA_DB.Save(&user).Error
	return &user, err
}

func (u *UserService) DeleteUser(id uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (u *UserService) SetUserInfo(entityUser entity.User) (err error) {
	result := global.GVA_DB.Model(&entity.User{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email").
		Where("id=?", entityUser.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  entityUser.NickName,
			"header_img": entityUser.HeaderIMG,
			"phone":      entityUser.Phone,
			"email":      entityUser.Email,
		})
	fmt.Printf("Updated %v rows\n", result.RowsAffected)
	return result.Error
}

func (u *UserService) ResetPassword(id uint) (err error) {
	return global.GVA_DB.Model(&entity.User{}).Where("id = ?", id).Update("password", "123456").Error
}

func (u *UserService) GetUserInfoList(pageInfo request.PageInfo) (list interface{}, total int64, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := global.GVA_DB.Model(&entity.User{})
	var userList []entity.User
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

func (u *UserService) GetUserInfo(id uint) (entityUser entity.User, err error) {
	var user entity.User
	err = global.GVA_DB.First(&user, "id = ?", id).Error
	return user, err
}
