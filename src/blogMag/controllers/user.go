package controllers

import (
	"blogMag/models"
	_ "encoding/base64"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"

)

const (
	//mnsDomain = "1943695596114318.mns.cn-hangzhou.aliyuncs.com"
	mnsDomain = "aaa.com@1252171207239820.onaliyun.com"
)

type RegController struct {
	beego.Controller
}

//注册
func (this *RegController) ShowReg() {
	this.TplName = "register.html"
}

func (this *RegController)HandleReg()  {
	userName := this.GetString("user")
	rpwd := this.GetString("rpwd")
	rpwdq := this.GetString("rpwdq")
	email := this.GetString("email")
	iphone := this.GetString("iphone")
	vfcode := this.GetString("vfcode")

	if userName == "" && rpwd == "" && rpwdq == "" && email == "" && iphone == "" && vfcode == ""{
		logs.Info("值不能为null")
		return
	}
	fmt.Printf("第一次输入的密码%d与第二次输入的密码%d",rpwd,rpwdq)
	if rpwd != rpwdq{
		logs.Info("两次密码不一致，请重新输入")
		return
	}
	//数据插入
	o := orm.NewOrm()
	user := models.User{}
	user.MagUser = userName
	user.MagPwd = rpwd
	user.MagEmail = email
	user.MagIphone = iphone
	o.Insert(&user)
	this.Redirect("/login",302)
}

type LoginController struct {
	beego.Controller
}
//登录
func (this *LoginController)ShowLogin()  {
	this.TplName = "login.html"
}
func (this *LoginController)HandleLogin()  {
	userName := this.GetString("user")
	userPwd := this.GetString("pwd")
	o := orm.NewOrm()
	user := models.User{MagUser:userName}
	err := o.Read(&user,"MagUser")
	if err != nil{
		logs.Info("没有用户")
		return
	}
	if user.MagPwd != userPwd {
		logs.Info("密码不一致")
		return
	}
	this.SetSession("name",userName)
	this.Redirect("/",302)
}