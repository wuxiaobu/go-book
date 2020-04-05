package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:wzx3410@tcp(127.0.0.1:3306)/mbook?charset=utf8", 30)
	orm.RegisterModel(new(Company))
}

type Company struct {
	Id      int
	Name    string
	Email   string
	Phone   string
	Website string
}

func PrintCompanyByOrm() {
	o := orm.NewOrm()
	user := Company{Id: 557}
	o.Read(&user)
	fmt.Println(user)
}
