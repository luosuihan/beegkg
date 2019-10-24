package main

import (
	_ "blogMag/routers"
	"github.com/astaxie/beego"
	_ "blogMag/models"
)


func main() {
	beego.AddFuncMap("pre",prePage)
	beego.Run()
}
func prePage(pre int)(preout int)  {
	preout = pre - 1
	return preout
}
func middlePage(pre int)  {

}
