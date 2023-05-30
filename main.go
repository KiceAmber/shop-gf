package main

import (
	_ "rime-shop-gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"rime-shop-gf/internal/cmd"

	_ "rime-shop-gf/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}
