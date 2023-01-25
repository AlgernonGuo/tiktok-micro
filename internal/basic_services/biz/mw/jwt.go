package mw

import (
	"context"
	"net/http"
	"time"

	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/service"

	utils2 "github.com/AlgernonGuo/tiktok-micro/internal/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "identity"
)

type UserLoginResponse struct {
	utils2.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour * 24 * 15, // 15 days
		MaxRefresh:    time.Hour * 24 * 15, // token can be refreshed for 15 days
		TokenLookup:   "param: token, query: token",
		TokenHeadName: "Bearer",
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: utils2.Response{StatusCode: 0, StatusMsg: "login success"},
				UserId:   c.GetInt64("user_id"),
				Token:    token,
			})
		},
		Authenticator: service.Login,
		IdentityKey:   IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			// I don't know why the value here will become a float64 type
			// I set the value to int64 type in the PayloadFunc function
			// But when I get the value here, it becomes a float64 type, I will check the source code later
			// After the conversion process there are no problems at the moment
			if claims[IdentityKey] == nil {
				hlog.CtxErrorf(ctx, "claims[IdentityKey] is nil")
				return int64(0)
			}
			hlog.Debugf("claims[IdentityKey] = %+v", int64(claims[IdentityKey].(float64)))
			return int64(claims[IdentityKey].(float64))
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// data is user id which type is int64
			// This value comes from the return value of Authenticator call to the service.Login function
			// We set the user id as the identity key in the token
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			hlog.CtxErrorf(ctx, "jwt err = %+v", e.Error())
			return e.Error()
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: utils2.Response{StatusCode: 400, StatusMsg: message},
			})
		},
		RefreshResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			//c.JSON(http.StatusOK, map[string]interface{}{})
			// Do nothing font-end not implement
		},
	})
	if err != nil {
		panic(err)
	}
}
