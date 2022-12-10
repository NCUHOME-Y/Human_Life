package router

import (
	"Hack/response"
	"Hack/zapLog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func InitRouter() {
	e := gin.Default()
	e.Use(cors.Default())
	//contextTimeOut, cancel = context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//登录注册鉴权
	Logined := e.Group("/Logined", response.AuthLogined())
	e.POST("/hack/register", response.Register)
	e.POST("/hack/login", response.Login)
	//时间提醒
	Logined.POST("/BulletinBoardPost", response.PostTimeRemind)
	Logined.GET("/BulletinBoardGet", response.GetTimeRemind)
	//添加室友
	Logined.POST("/addPartner", response.AddPartner)
	Logined.GET("/PartnerMessage", response.GetPartnerMessage)
	//状态
	Logined.POST("/remindTime", response.ChangeStates)
	Logined.GET("/remindTime/GetMessage", response.GetStateMessage)
	//值日生
	Logined.POST("/Duty/add", response.PostDuty)
	//Logined.GET("/Duty/get", response.GetDuty)
	err := e.Run(":5500")
	if err != nil {
		zapLog.SugarLogger.Fatalln(err)
		os.Exit(1)
	}
}
