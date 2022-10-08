package requests

type SysUserRegisterReq struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	HeaderImg string `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	Email     string `json:"email"`
}

type SysUserLoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	/*
		Captcha   string `json:"captcha"`   // 验证码
		CaptchaId string `json:"captchaId"` // 验证码ID
	*/
}

type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type ChangeUserInfoReq struct {
	ID        int
	Email     string `json:"email"`
	HeaderImg string `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
}
