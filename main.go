package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/novikk/medbay/routers"
)

func main() {
	beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(os.Getenv("PORT"))
	beego.Run()
}
