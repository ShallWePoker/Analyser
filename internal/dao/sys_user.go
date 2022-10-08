package dao

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/models"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/utils"
	"gorm.io/gorm"
)

type SysUserDao struct {}

func (sysUserDao *SysUserDao) Register(u models.SysUser) (userInter models.SysUser, err error) {
	var user models.SysUser
	if !errors.Is(db.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("user name already exists")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = db.Create(&u).Error
	return u, err
}

func (sysUserDao *SysUserDao) Login(u *models.SysUser) (userInter *models.SysUser, err error) {
	var user models.SysUser
	err = db.Where("username = ?", u.Username).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("wrong password")
		}
	}
	return &user, err
}

func (sysUserDao *SysUserDao) ChangePassword(u *models.SysUser, newPassword string) (userInter *models.SysUser, err error) {
	var user models.SysUser
	if err = db.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("wrong original password")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = db.Save(&user).Error
	return &user, err
}

func (sysUserDao *SysUserDao) SetUserInfo(user models.SysUser) error {
	return db.Updates(&user).Error
}