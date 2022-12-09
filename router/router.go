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
	Logined := e.Group("/Logined", response.AuthLogined())
	e.POST("/hack/register", response.Register)
	e.POST("/hack/login", response.Login)
	Logined.POST("/BulletinBoardPost", response.PostBulletinBoard)
	Logined.GET("/BulletinBoardGet", response.GetBulletinBoard)
	err := e.Run(":5500")
	if err != nil {
		zapLog.SugarLogger.Fatalln(err)
		os.Exit(1)
	}
}
