package router

import (
	"Hack/response"
	"Hack/zapLog"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		if method := c.Request.Method; method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}

func InitRouter() {
	e := gin.Default()
	e.Use(CORS())
	//contextTimeOut, cancel = context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//登录注册鉴权
	Logined := e.Group("/Logined", response.AuthLogined()) //ok
	Logined.Use(CORS())
	e.POST("/hack/register", response.Register) //ok
	e.POST("/hack/login", response.Login)       //ok
	//时间提醒
	Logined.POST("/remindTime", response.PostTimeRemind)          //ok
	Logined.GET("/remindTime/GetMessage", response.GetTimeRemind) //ok
	//添加室友
	Logined.POST("/addPartner", response.AddPartner)           //ok
	Logined.GET("/PartnerMessage", response.GetPartnerMessage) //ok
	//状态
	Logined.POST("/State/Change", response.ChangeStates)       //ok
	Logined.GET("/State/GetMessage", response.GetStateMessage) //ok
	//值日生
	Logined.POST("/Duty/add", response.PostDuty) //ok
	Logined.GET("/Duty/Get", response.GetDuty)   //ok
	//公告板
	Logined.POST("/Board/post", response.PostBoard) //ok
	Logined.GET("/Board/get", response.GetBoard)    //ok
	err := e.Run(":5500")
	if err != nil {
		zapLog.SugarLogger.Fatalln(err)
		os.Exit(1)
	}
}
