package service

import (
	"context"
	"rime-shop-gf/internal/consts"
	"rime-shop-gf/internal/model"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "rime-shop",
		Key:             []byte(consts.SecretKey),
		Timeout:         time.Hour * 3,
		MaxRefresh:      time.Hour * 3,
		IdentityKey:     "id",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

// Authenticator 负责身份认证的函数，必须携带身份验证的键值对，比如 IdentityKey:"id",则身份认证回调函数应当包含{"id": 1}
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in model.UserLoginInput
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}
	if user := Admin().GetUserByUserNamePassword(ctx, in); user != nil {
		return user, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// Unauthorized 授权失败(账号密码错误，刷新失败，token过期等)，返回的信息，可以自定义
func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})

	r.ExitAll()
}

// PayloadFunc 用于自定义私有的 Payload，也就是校验时用于识别的一些自定义字段
// 默认是将 Authenticator 函数返回的 map[string]interface 循环放入载荷中
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler 用于从 JWT 中解析获取 Payload
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}
