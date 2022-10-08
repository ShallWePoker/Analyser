package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/configs"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/dao"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/models"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/requests"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/responses"
	"github.com/shallwepoker/ggpoker-hands-converter/internal/utils"
)

var (
	userDao = dao.SysUserDao{}
)

func GroupSysUsers(g *gin.RouterGroup) {
	group := g.Group(fmt.Sprintf("%s/users", configs.Config.UrlPrefix))

	group.POST("/register", wrapper(userRegister))
	group.POST("/login", wrapper(userLogin))
	group.POST("/change-password", wrapper(userChangePassword))
	group.POST("/set-self-info", wrapper(userSetSelfInfo))
}

func userRegister(c *gin.Context) error {
	var r requests.SysUserRegisterReq
	err := c.ShouldBindJSON(&r)
	if err != nil {
		return err
	}
	user := &models.SysUser{
		Username:  r.Username,
		Password:  r.Password,
		HeaderImg: r.HeaderImg,
		Email:     r.Email,
		Enable:    1,
	}
	userReturn, err := userDao.Register(*user)
	if err != nil {
		return err
	}
	return SuccessResp(c, responses.SysUserResponse{User: userReturn})
}

func userLogin(c *gin.Context) error {
	var l requests.SysUserLoginReq
	err := c.ShouldBindJSON(&l)
	if err != nil {
		return err
	}
	u := &models.SysUser{
		Username: l.Username,
		Password: l.Password,
	}
	user, err := userDao.Login(u)
	if err != nil {
		return err
	}
	return tokenNext(c, *user)
}

func tokenNext(c *gin.Context, user models.SysUser) error {
	j := &utils.JWT{SigningKey: []byte(configs.Config.JWT.SigningKey)}
	claims := j.CreateClaims(requests.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		logger.Errorf("create token for user %s err: %+v", user.Username, err)
		return err
	}
	return SuccessResp(c, responses.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	})
}

func userChangePassword(c *gin.Context) error {
	var req requests.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return err
	}
	userId := GetUserID(c)
	u := &models.SysUser{
		GormModel: models.GormModel{ID: userId},
		Password:  req.Password,
	}
	userReturn, err := userDao.ChangePassword(u, req.NewPassword)
	if err != nil {
		return err
	}
	return SuccessResp(c, responses.SysUserResponse{User: *userReturn})
}

func userSetSelfInfo(c *gin.Context) error {
	var req requests.ChangeUserInfoReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		return err
	}
	req.ID = GetUserID(c)
	err = userDao.SetUserInfo(models.SysUser{
		GormModel: models.GormModel{ID: req.ID},
		Email:     req.Email,
		HeaderImg: req.HeaderImg,
	})
	if err != nil {
		return err
	}
	return SuccessResp(c, nil)
}

func GetClaims(c *gin.Context) (*requests.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := utils.NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		logger.Errorf("parse jwt token x-token err: %+v", err)
	}
	return claims, err
}

func GetUserID(c *gin.Context) int {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.ID
		}
	} else {
		waitUse := claims.(*requests.CustomClaims)
		return waitUse.ID
	}
}
