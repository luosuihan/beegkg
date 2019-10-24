package routers

import (
	"blogMag/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    //index
    beego.Router("/", &controllers.MainController{},"get:ShowIndex")
    //注册
    beego.Router("register", &controllers.RegController{},"get:ShowReg;post:HandleReg")
	//登录
	beego.Router("login",&controllers.LoginController{},"get:ShowLogin;post:HandleLogin")
	//文章添加
	beego.Router("addArticle",&controllers.AddArticleController{},"get:ShowAddArticle;post:HandleAddArticle")
	//文章类型
    beego.Router("articleType",&controllers.TypeArticleController{},"get:ShowArticleType;post:HandleArticleType")
}
