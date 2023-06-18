package session

import (
	"context"
	"rime-shop-gf/internal/model/entity"
	"rime-shop-gf/internal/service"
)

const (
	sessionKeyUser = "SessionKeyUser" // 存放用户信息的 Key
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return &sSession{}
}

// SetUser 设置用户 Session
func (s *sSession) SetUser(ctx context.Context, user *entity.AdminInfo) error {
	return service.BizCtx().Get(ctx).Session.Set(sessionKeyUser, user)
}

// GetUser 获取当前登录的用户信息对象，如果用户未登录返回nil。
func (s *sSession) GetUser(ctx context.Context) *entity.AdminInfo {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(sessionKeyUser)
		if !v.IsNil() {
			var user *entity.AdminInfo
			_ = v.Struct(&user)
			return user
		}
	}
	return &entity.AdminInfo{}
}

// RemoveUser 删除用户Session。
func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(sessionKeyUser)
	}
	return nil
}
