package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"rime-shop-gf/internal/dao"
	"rime-shop-gf/internal/model"
	"rime-shop-gf/internal/model/entity"
	"rime-shop-gf/internal/service"
	"rime-shop-gf/utility"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// Login 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	//验证账号密码是否正确
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return err
	}

	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("账号或者密码不正确")
	}

	// 登录后，需要在 Session 设置用户信息
	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线 for session
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: uint8(adminInfo.IsAdmin),
	})
	return nil
}

// Logout 注销
func (s *sLogin) Logout(ctx context.Context) error {
	return service.Session().RemoveUser(ctx)
}
