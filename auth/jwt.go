package auth

import (
	"gin-admin/config"
	"gin-admin/models"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtConfig = config.InitJwtConfig()
var jwtSecret = []byte(jwtConfig.GetString("jwt.secret"))

type Claims struct {
	Id       string `json:"uid""`
	UserName string `json:"username"`
	//PassWord string `json:"password"`
	jwt.StandardClaims
}

// 生成token方法
func GenerateToken(user *models.User) (string, error) {
	nowTime := time.Now()
	// 设置token 过期时间 1天
	expireTime := nowTime.Add(time.Hour * 24).Unix()
	// 生成 claims
	claims := Claims{
		Id:       user.Id,
		UserName: user.UserName,
		//PassWord: user.PassWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,                      // 过期时间
			Issuer:    jwtConfig.GetString("jwt.sign"), //指定发行人
		},
	}
	// 该方法内部生成签名字符串，再用于获取完整、已签名的token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息(进而获取其中的用户名和密码)
func ParseToken(token string) (*Claims, error) {
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目结构体都是用指针传递，节省空间
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { // Valid()验证基于时间的声明
			return claims, nil
		}
	}
	return nil, err
}
