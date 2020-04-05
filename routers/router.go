package routers

import (
	"book/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/hi", &controllers.MainController{}, "get:Hi")
	beego.Router("/explore", &controllers.ExploreController{}, "get:Index")
	beego.Router("/books/:key", &controllers.DocumentController{}, "*:Index")

	//读书
	beego.Router("/read/:key/:id", &controllers.DocumentController{}, "*:Read")
	beego.Router("/read/:key/search", &controllers.DocumentController{}, "post:Search")

	//login
	beego.Router("/login", &controllers.AccountController{}, "*:Login")
	beego.Router("/regist", &controllers.AccountController{}, "get:Regist")
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")
	beego.Router("/regist", &controllers.AccountController{}, "post:RegistPost")

	//编辑
	beego.Router("/api/:key/edit/?:id", &controllers.DocumentController{}, "*:Edit")
	beego.Router("/api/:key/content/?:id", &controllers.DocumentController{}, "*:Content")
	beego.Router("/api/upload", &controllers.DocumentController{}, "post:Upload")
	beego.Router("/api/:key/create", &controllers.DocumentController{}, "post:Create")
	beego.Router("/api/:key/delete", &controllers.DocumentController{}, "post:Delete")

	//搜索
	beego.Router("/search", &controllers.SearchController{}, "get:Search")
	beego.Router("/search/result", &controllers.SearchController{}, "get:Result")

	beego.Router("/book/score/:id", &controllers.BookController{}, "*:Score")                   //评分
	beego.Router("/book/comment/:id", &controllers.BookController{}, "post:Comment")
}
