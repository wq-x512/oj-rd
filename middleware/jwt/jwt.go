package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
	SignKey          = "xwq200505123337" // 签名信息应该设置成动态从库中获取
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			ReturnError(c, 401, "请求未携带token,无权限访问")
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParserToken(token)
		if err != nil {
			if errors.Is(err, TokenExpired) {
				ReturnError(c, 401, "token授权已过期,请重新申请授权")
				c.Abort()
				return
			}
			ReturnError(c, 401, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	UserID string `json:"userID"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

func GetSignKey() string {
	return SignKey
}

func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParserToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}

		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid

}

// UpdateToken 更新Token
func (j *JWT) UpdateToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", fmt.Errorf("token获取失败:%v", err)
}
