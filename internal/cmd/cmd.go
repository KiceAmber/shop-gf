package cmd

import (
	"context"
	"rime-shop-gf/internal/controller"
	"rime-shop-gf/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Rotation, // 轮播图
					controller.Position, // 手工位
					controller.Admin,    // 管理员
					controller.Login,    // 后台用户登录
					controller.Data,     // 数据大屏
				)
			})
			s.Run()
			return nil
		},
	}
)
