package controllers
//文章管理类
import (
	"blogMag/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"strconv"
	"time"
)

//文章列表页
type MainController struct {
	beego.Controller
}

func (this *MainController) ShowIndex() {
	pageIndex := this.GetString("pageIndex")
	pageIndex1, err := strconv.Atoi(pageIndex)
	if err != nil {
		pageIndex1 = 1
	}
	//var pageConut float64
	var pageM float64
	pageSize := 2
	logs.Info("当前页面pageIndex = ",pageIndex)
	o := orm.NewOrm()
	var articleList []models.Article
	count, _ := o.QueryTable("Article").RelatedSel("ArticleType").Count() //获取数据总条数
	pageM = float64(count) / 2
	pageM1 := math.Ceil(pageM)
	start := pageSize * (pageIndex1 - 1)
	//pageConut = float64(count) / float64(pageSize)
	//pageConut1 := math.Ceil(pageConut)
	logs.Info("当前页面pageM1 = ",pageM1)
	o.QueryTable("Article").RelatedSel("ArticleType").Limit(pageSize,start).All(&articleList)
	this.Data["page"] = 5 //page 表示当前页
	this.Data["pageM"] = pageM1
	this.Data["listLen"] = len(articleList) //总页数
	this.Data["articleList"] = articleList
	this.Layout = "pulicstyle.html"
	this.TplName = "index.html"
}
//添加文章
type AddArticleController struct {
	beego.Controller
}

func (this *AddArticleController)ShowAddArticle()  {
	//查询类型
	o := orm.NewOrm()
	var articleType []models.ArticleType
	o.QueryTable("ArticleType").All(&articleType)
	this.Data["articleType"] = articleType
	this.Layout = "pulicstyle.html"
	this.TplName = "add.html"
}
func (this *AddArticleController)HandleAddArticle()  {
	title := this.GetString("title")
	atype2 := this.GetString("atype2") //文章类型类型
	sub := this.GetString("substance")
	f, h, err := this.GetFile("upimg")
	if title == "" && sub == "" {
		logs.Info("标题与内容均不能为null")
		return
	}
	defer f.Close()
	if err != nil {
		logs.Info("文件获取失败")
	}
	//获取文件后缀名
	ext := path.Ext(h.Filename)
	logs.Info("文件后缀名 = ",ext)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg"{
		logs.Info("请上传正确图片")
		return
	}
	fileName := time.Now().Format("2006-01-02 15-04-05")
	err = this.SaveToFile("upimg", "./static/updateImg"+fileName+ext)
	if err != nil{
		logs.Info("文件上传服务器失败")
		return
	}
	//userName := this.GetSession("name")
	o := orm.NewOrm()
	article := models.Article{}
	article.Title = title
	article.Substance = sub
	article.Img = "./static/updateImg"+fileName+ext
	article.Utime = time.Now()
	//插入类型 start
	var articleType models.ArticleType
	articleType.TypeName = atype2
	err2 := o.Read(&articleType, "TypeName")
	if err2 != nil {
		logs.Info("数据类型查询为null")
		return
	}
	article.ArticleType = &articleType
	//end
	_, err1 := o.Insert(&article)
	if err1 != nil {
		logs.Info("文件插入失败 = ",err1)
		return
	}
	this.Redirect("addArticle",302)
}
//文章类型
type TypeArticleController struct {
	beego.Controller
}
func (this *TypeArticleController) ShowArticleType() {
	typeId, _ := this.GetInt("typeId")
	o := orm.NewOrm()
	var articleType []models.ArticleType
	o.QueryTable("ArticleType").All(&articleType)
	this.Data["article"] = articleType
	deleteArticle := models.ArticleType{Uid: typeId}
	o.Delete(&deleteArticle)
	this.Layout = "pulicstyle.html"
	this.TplName = "feilei.html"
}
func (this *TypeArticleController)HandleArticleType()  {
	s := this.GetString("atype")
	logs.Info("类型 = ",s)
	o := orm.NewOrm()
	articleType := models.ArticleType{}
	articleType.TypeName = s
	err := o.Read(&articleType, "TypeName")
	if err == nil {
		logs.Info("类型已经存在 = ",err)
		this.Redirect("articleType",302)
		return
	}
	_, err = o.Insert(&articleType)
	if err != nil {
		logs.Info("数据插入失败 ",err)
	}
	this.Redirect("articleType",302)
}