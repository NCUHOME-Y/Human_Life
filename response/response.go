package response

import (
	"Hack/define"
	"Hack/initDatabase"
	"Hack/myJWT"
	"Hack/zapLog"
	"github.com/gin-gonic/gin"
	"time"
)

func UserInOneRoom() {

}

// ForgetPassword 忘记密码部分 暂时不用 后期迭代会用
//func ForgetPassword(ctx *gin.Context) {
//	var user define.UserLogin
//	err := ctx.ShouldBind(&user)
//	if err != nil {
//		zapLog.SugarLogger.Info("bind error", err)
//		ctx.JSON(400, gin.H{
//			"message": "输入的格式有误",
//			"error":   err,
//		})
//		return
//	}
//	err = initDatabase.DB.Model(&define.UserLogin{}).Where("pass_word=?", user.PassWord).Updates(&user).Error
//	if err != nil {
//		zapLog.SugarLogger.Debug("update error", err)
//		ctx.JSON(500, gin.H{
//			"message": "server error",
//			"error":   err,
//		})
//		return
//	}
//}

// Login 登录注册鉴权部分 包括了JWT的获取
func Login(ctx *gin.Context) {
	var User define.UserLogin
	var UserExist define.UserLogin
	err := ctx.ShouldBind(&User)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(500, gin.H{
			"message": "用户名或密码格式有误",
			"error":   err,
		})
		return
	}

	err = initDatabase.DB.Model(&define.UserLogin{}).Where("user_name=?", User.UserName).Find(&UserExist).Error
	if err != nil {
		zapLog.SugarLogger.Debug(err)
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	if define.CheckAmt(UserExist) {
		zapLog.SugarLogger.Info("empty Login message")
		ctx.JSON(400, gin.H{
			"message": "empty Login message",
		})
		return
	}
	if UserExist.UserName == User.UserName && UserExist.PassWord == User.PassWord {
		Token, err := GetJwt(UserExist.UserName)
		if err != nil {
			zapLog.SugarLogger.Debug(err)
			ctx.JSON(500, gin.H{
				"message": err,
			})
			return
		}
		Newtoken := define.Token{
			Token: Token,
			UID:   UserExist.UserName,
		}
		err = initDatabase.DB.Model(&define.Token{}).Create(&Newtoken).Error
		if err != nil {
			zapLog.SugarLogger.Debug(err)
			ctx.JSON(500, gin.H{
				"message": "create token error",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "登录成功",
			"token":   Token,
			"UUID":    UserExist.UserName,
		})
	} else {
		zapLog.Logger.Error("用户名和密码错误")
		ctx.JSON(400, gin.H{
			"message": "用户名和密码错误",
		})
	}

}

func AuthLogined() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var UserLogined define.UserLogin
		var usertoken define.Token
		err := ctx.ShouldBind(&UserLogined)
		Token := ctx.Request.Header.Get("token")
		if err != nil {
			zapLog.SugarLogger.Info("bind error", err)
			ctx.JSON(400, gin.H{
				"message": "输入的格式有误",
				"error":   err,
			})
			return
		}
		err = initDatabase.DB.Model(&define.Token{}).Where("uid=?", UserLogined.UserName).Find(&usertoken).Error
		if err != nil {
			zapLog.SugarLogger.Debug("Find error", err)
			ctx.JSON(500, gin.H{
				"message": "server error",
				"error":   err,
			})
			return
		}
		if usertoken.Token != Token {
			zapLog.SugarLogger.Info("token 不正确")
			ctx.JSON(400, gin.H{
				"Message": "token error 无权访问",
			})
		}
		ctx.Next()
	}
}

func Register(ctx *gin.Context) {
	var NewUser define.UserLogin
	var User define.UserLogin
	err := ctx.ShouldBind(&NewUser)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	initDatabase.DB.Model(&define.UserLogin{}).Where("user_name=?", NewUser.UserName).Find(&User)
	if User.UserName == NewUser.UserName {
		ctx.JSON(500, gin.H{
			"message": "user is already exist",
		})
		return
	} else {
		err = initDatabase.DB.Model(&define.UserLogin{}).Create(&NewUser).Error
		if err != nil {
			zapLog.SugarLogger.Debug(err)
			ctx.JSON(500, gin.H{
				"message": "create error",
			})
			return
		}

	}
	ctx.JSON(200, gin.H{
		"message": "注册成功",
		"UserID":  NewUser.UserName,
	})
}

func GetJwt(usr uint) (string, error) {
	newJwtHeader := myJWT.JwtHeader{
		Alg: "HS256",
		Typ: "JWT",
	}
	timeNow := time.Now()
	newJwtPayLoad := myJWT.JWTPayLoad{
		Iss: "server Manna. 43.143.227.115",
		Iat: timeNow.Format("2006-01-02"),
		Jti: usr,
	}
	keygen := "zcf304422"
	Header, err := myJWT.NewMarshaler(newJwtHeader)
	if err != nil {
		return "", err
	}
	PayLoad, err := myJWT.NewMarshaler(newJwtPayLoad)
	if err != nil {
		return "", err
	}
	JWT := myJWT.Base64Encode(PayLoad, Header, keygen)
	return JWT, nil
}

// PostBulletinBoard 公告板功能
func PostBulletinBoard(ctx *gin.Context) {
	var bulletinBoard define.BulletinBoard
	err := ctx.ShouldBind(&bulletinBoard)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	err = initDatabase.DB.Model(&define.BulletinBoard{}).Create(&bulletinBoard).Error
	if err != nil {
		zapLog.SugarLogger.Debug(err)
		ctx.JSON(500, gin.H{
			"message": "create error",
		})
		return
	}
}

func GetBulletinBoard(ctx *gin.Context) {
	var GetBulletinBoard define.BulletinBoard
	var ReadDB define.BulletinBoard
	err := ctx.ShouldBind(&GetBulletinBoard)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	err = initDatabase.DB.Model(&define.BulletinBoard{}).Where("time=?", GetBulletinBoard.Time).Find(&ReadDB).Error
	if err != nil {
		zapLog.SugarLogger.Debug(err)
		ctx.JSON(500, gin.H{
			"message": "create error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"time":    ReadDB.Time,
		"message": ReadDB.Message,
	})

}
