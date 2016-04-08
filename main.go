package main

import (
	"github.com/astaxie/beego"
	_ "moonlord/models"
	_ "moonlord/routers"
)

func main() {
	beego.Run()
}
