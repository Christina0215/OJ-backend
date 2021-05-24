package main

import (
	"qkcode/boot/config"
	"qkcode/boot/http"
	"qkcode/boot/log"
	"qkcode/boot/orm"
	"qkcode/route"
)

func _init() {
	config.InitConfig()
	log.InitLog()
	log.InitTimer()
	orm.InitOrm()

	http.InitHttp()

	route.AddRoute()
	route.AddStaticRoute()
}

func _end() {
	orm.EndOrm()
}

func main() {
	_init()
	http.Run()
	defer _end()
}
