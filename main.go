package main

import (
	_ "book/routers"
	_ "book/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

