package main

import (
	"fmt"
	"pkg/conf"
	"pkg/migrations"
	"pkg/router"
)

func main() {
	migrations.Migrate()
	r := router.Router()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常", err)
		}
	}()
	if err := r.Run(string(":" + conf.HttpPort)); err != nil {
		fmt.Println("发生异常")
		return
	}
}
