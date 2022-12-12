package initDatabase

import (
	"Hack/define"
	"Hack/zapLog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:@tcp(43.143.227.115:3306)/Hackweek?charset=utf8&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		zapLog.SugarLogger.Info(err)
	}
	err = DB.AutoMigrate(&define.UserLogin{}, &define.Token{}, &define.TimeRemind{}, &define.DutyOrder{}, define.Board{})
	if err != nil {
		zapLog.SugarLogger.Fatal(err)
	}
}
