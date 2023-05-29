package main

import (
	_ "rime-shop-gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"rime-shop-gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
