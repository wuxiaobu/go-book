package sysinit

import (
	"book/models"
	"book/utils"
	"github.com/astaxie/beego"
	"path/filepath"
	"strings"
)

func sysinit() {
	uploads := filepath.Join("./",  "uploads")
	beego.BConfig.WebConfig.StaticDir["/uploads"] = uploads

	//前端函数注册
	RegisterFunc()
}

//RegisterFunc
func RegisterFunc() {
	beego.AddFuncMap("cdnjs", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdnjs", "")
		if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:])
		}
		if !strings.HasPrefix(p, "/") && !strings.HasSuffix(cdn, "/") {
			return cdn + "/" + p
		}
		return cdn + p
	})
	beego.AddFuncMap("cdncss", func(p string) string {
		cdn := beego.AppConfig.DefaultString("cdncss", "")
		if strings.HasPrefix(p, "/") && strings.HasSuffix(cdn, "/") {
			return cdn + string(p[1:])
		}
		if !strings.HasPrefix(p, "/") && !strings.HasSuffix(cdn, "/") {
			return cdn + "/" + p
		}
		return cdn + p
	})
	beego.AddFuncMap("getUsernameByUid", func(id interface{}) string {
		return new(models.Member).GetUsernameByUid(id)
	})
	beego.AddFuncMap("getNicknameByUid", func(id interface{}) string {
		return new(models.Member).GetNicknameByUid(id)
	})
	beego.AddFuncMap("inMap", utils.InMap)

	//	//用户是否收藏了文档
	beego.AddFuncMap("doesCollection", new(models.Collection).DoesCollection)
	//	beego.AddFuncMap("scoreFloat", utils.ScoreFloat)
	beego.AddFuncMap("showImg", utils.ShowImg)
	beego.AddFuncMap("IsFollow", new(models.Fans).Relation)
	beego.AddFuncMap("isubstr", utils.Substr)
}
