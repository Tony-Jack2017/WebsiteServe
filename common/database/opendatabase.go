package database

import (
	"Serve/tools/color_print"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func ConnectDatabase() *gorm.DB  {
	dsn := "user:password@tcp(127.0.0.1:3306)/apis-admin?charset=utf8&parseTime=True&loc=Local&timeout=1000ms"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		color_print.ColorPrintln("数据库连接错误，Error:", color_print.Fail(err))
		panic(err)
	}
	return db
}
