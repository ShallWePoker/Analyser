package models

import (
	"github.com/satori/go.uuid"
)

type SysUser struct {
	GormModel
	UUID      uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	Username  string    `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password  string    `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	HeaderImg string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Email     string    `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	Enable    int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                      //用户是否被冻结 1正常 2冻结
}
