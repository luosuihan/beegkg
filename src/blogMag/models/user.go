package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//用户相关
type User struct {
	Id int
	MagUser string `orm:"size(20);unique"`
	MagPwd string	`orm:"size(70)"`
	MagEmail string `orm:"size(60)"`
	MagIphone string `orm:"size(60)"`
	Active bool	`orm:"default(false)"`  //是否激活
	Articles[]*Article `orm:"rel(m2m)"`
}
type Article struct {
	Id int
	Title string `orm:"size(60)"`
	Substance string `orm:"size(600)"`
	Img string `orm:"size(120)"`
	Utime time.Time `orm:type(datetime);auto_now_add`
	Users[]*User `orm:"reverse(many)"` //表示一对多
	ArticleType *ArticleType `orm:"rel(fk)"`  //外键
	Count int `orm:"default(0)"`
}
type ArticleType struct {
	Uid int `orm:"pk;auto"`
	TypeName string `orm:"size(20);unique"`
	Articles[]*Article `orm:"reverse(many)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@(127.0.0.1:3306)/blog2?charset=utf8&loc=Local")
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	orm.RunSyncdb("default",false,true)
}
