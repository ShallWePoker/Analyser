package responses

import "github.com/shallwepoker/ggpoker-hands-converter/internal/models"

type SysUserResponse struct {
	User models.SysUser `json:"user"`
}

type LoginResponse struct {
	User      models.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}
