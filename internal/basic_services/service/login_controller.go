package service

import (
	"context"

	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/biz"
	"github.com/AlgernonGuo/tiktok-micro/internal/basic_services/data"
	"github.com/AlgernonGuo/tiktok-micro/internal/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type UserResponse struct {
	utils.Response
	User data.User `json:"user"`
}

func Login(ctx context.Context, c *app.RequestContext) (interface{}, error) {
	var login struct {
		Username string `query:"username,required" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
		Password string `query:"password,required" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
	}
	if err := c.BindAndValidate(&login); err != nil {
		return nil, err
	}

	user, err := biz.NewLoginService().Login(login.Username, utils.MD5(login.Password))
	if err != nil {
		return nil, err
	}
	c.Set("user_id", user.Id)
	return user.Id, nil
}

func Register(ctx context.Context, c *app.RequestContext) error {
	var register struct {
		Username string `query:"username,required" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
		Password string `query:"password,required" vd:"(len($) > 0 && len($) < 128); msg:'Illegal format'"`
	}
	if err := c.BindAndValidate(&register); err != nil {
		hlog.CtxWarnf(ctx, "Bind failed, err: %v", err)
		return err
	}
	user := &data.User{
		Id:       utils.GenID(),
		Name:     register.Username,
		Password: utils.MD5(register.Password),
	}
	err := biz.NewLoginService().Register(user)
	if err != nil {
		hlog.CtxWarnf(ctx, "Register failed")
		return err
	}
	return nil
}
