package sysinit

import (
	_ "book/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)


func dbinit(aliases ...string ){
	//如果是开发模式，显示信息
	isDev := beego.AppConfig.String("runmode") == "dev"
	if len(aliases) > 0 {
		for _, alias := range aliases {
			registerDatabase(alias)
			//主库自动建表
			if alias == "w" {
				orm.RunSyncdb("default", false, isDev)
			}
		}
	} else {
		registerDatabase("w")
		orm.RunSyncdb("default", false, isDev)
	}

	if isDev {
		orm.Debug = isDev
	}

}

func registerDatabase(alias string) {
	if len(alias) <= 0 {
		return
	}
	dbAlias := alias
	if dbAlias == "w" || dbAlias == "default" || len(dbAlias) <= 0 {
		dbAlias = "default"
		alias = "w"
	}
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	dbPassword := beego.AppConfig.String("db_" + alias + "_password")
	dbHost := beego.AppConfig.String("db_" + alias + "_host")
	dbPort := beego.AppConfig.String("db_" + alias + "_port")

	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8", 30)
}