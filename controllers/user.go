package controllers

import (
	"errors"
	"fmt"
	jwtgo "pkg/middleware/jwt"
	"pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

func (UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")
	user, err := models.User{ID: id}.QueryUser()
	if err == nil {
		ReturnSuccess(c, 0, "suceess", user, 1)
		return
	}
	ReturnError(c, 1, err)
}
func (UserController) GetList(c *gin.Context) {
	Users, err := models.GetAllUsers()
	if err == nil {
		ReturnSuccess(c, 0, "success", Users, int64(len(Users)))
		return
	}
	ReturnError(c, 1, err)
}
func (UserController) CreateUser(c *gin.Context) {
	data := &models.User{}
	_ = c.BindJSON(&data)
	if err := models.CreateUser(*data); err != nil {
		ReturnError(c, 1, err)
		return
	}
	ReturnSuccess(c, 0, "success", data, 1)
}
func (UserController) UserLogin(c *gin.Context) {
	data := make(map[string]string)
	_ = c.BindJSON(&data)
	user, err := models.User{Account: data["username"], Password: data["password"]}.QueryUser()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ReturnError(c, 1, fmt.Sprintf("未找到名为 %s 的用户", data["username"]))
			return
		} else {
			ReturnError(c, 1, fmt.Sprintf("查询出错 %v", err))
			return
		}
	} else if user.ID != "" {
		token_access, _ := generateToken(c, user, 3600)
		token_refresh, _ := generateToken(c, user, 7*24*60*60)
		var tokens []string
		tokens = append(tokens, token_access)
		tokens = append(tokens, token_refresh)
		ReturnSuccess(c, 0, tokens, user, 1)
		return
	}
	ReturnError(c, 1, fmt.Sprintf("未找到名为 %s 的用户", data["username"]))
}
func (UserController) UserLogout(c *gin.Context) {
	ReturnSuccess(c, 0, "success", "", 1)
}
func (UserController) GetAccessToken(c *gin.Context) {
	data := make(map[string]string)
	_ = c.BindJSON(&data)
	user := models.User{Account: data["username"], Password: data["password"]}
	token_access, _ := generateToken(c, user, 3600)
	token_refresh, _ := generateToken(c, user, 7*24*60*60)
	var tokens []string
	tokens = append(tokens, token_access)
	tokens = append(tokens, token_refresh)
	ReturnSuccess(c, 0, "success", tokens, 1)
}

func generateToken(c *gin.Context, user models.User, Time int64) (string, error) {
	j := jwtgo.NewJWT()
	claims := jwtgo.CustomClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: time.Now().Unix() + Time, // 签名过期时间
			Issuer:    "xwq",                    // 签名颁发者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*jwtgo.CustomClaims)
	if claims != nil {
		ReturnSuccess(c, 0, "token有效", claims, 1)
	}
}
