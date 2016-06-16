package models

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var BeegoOrm orm.Ormer
var Step int64

func init() {
	mariadbInit()
}

func mariadbInit() {
	var err error

	orm.RegisterModel(
		new(MaxSeqInfo),
	)
	mode := beego.BConfig.RunMode
	user := beego.AppConfig.String(mode + "::user")
	pwd := beego.AppConfig.String(mode + "::password")
	host := beego.AppConfig.String(mode + "::host")
	dbname := beego.AppConfig.String(mode + "::dbname")

	Step, _ = beego.AppConfig.Int64("maxSeqNum")

	if mode == "dev" {
		orm.Debug = true
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)
	err = orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s)/?loc=Local", user, pwd, host))
	if err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(0)
	}

	BeegoOrm = orm.NewOrm()
	BeegoOrm.Raw("CREATE DATABASE IF NOT EXISTS `" + dbname + "` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;").Exec()
	err = orm.RegisterDataBase(dbname, "mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?loc=Local", user, pwd, host, dbname))
	if err != nil {
		fmt.Println("数据库连接失败")
		os.Exit(0)
	}
	orm.RunSyncdb(dbname, false, false)
	BeegoOrm.Using(dbname)
}
