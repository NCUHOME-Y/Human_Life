package response

import (
	"Hack/define"
	"Hack/initDatabase"
	"Hack/myJWT"
	"Hack/zapLog"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strings"
	"time"
)

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
	tx := initDatabase.DB.Begin()
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

	err = tx.Model(&define.UserLogin{}).Where("user_name=?", User.UserName).Find(&UserExist).Error
	if err != nil {
		zapLog.SugarLogger.Debug(err)
		ctx.JSON(400, gin.H{
			"error": err,
		})
		tx.Rollback()
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
		err = tx.Model(&define.Token{}).Create(&Newtoken).Error
		if err != nil {
			zapLog.SugarLogger.Debug(err)
			ctx.JSON(500, gin.H{
				"message": "create token error",
			})
			tx.Rollback()
			return
		}

		ctx.JSON(200, gin.H{
			"message": "登录成功",
			"token":   Token,
			"UUID":    UserExist.UserName,
		})
		tx.Commit()
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
	tx := initDatabase.DB.Begin()
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
	err = tx.Model(&define.UserLogin{}).Where("user_name=?", NewUser.UserName).Find(&User).Error
	if err != nil {
		zapLog.SugarLogger.Debug("find error", err)
		ctx.JSON(500, err)
		return
	}
	if User.UserName == NewUser.UserName {
		ctx.JSON(500, gin.H{
			"message": "user is already exist",
		})
		return
	} else {
		err = tx.Model(&define.UserLogin{}).Create(&NewUser).Error
		if err != nil {
			zapLog.SugarLogger.Debug(err)
			ctx.JSON(500, gin.H{
				"message": "create error",
			})
			tx.Rollback()
			return
		}

	}
	tx.Commit()
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

// PostTimeRemind  公告板功能
func PostTimeRemind(ctx *gin.Context) {
	tx := initDatabase.DB.Begin()
	var bulletinBoard define.TimeRemind
	err := ctx.ShouldBind(&bulletinBoard)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	err = tx.Model(&define.TimeRemind{}).Create(&bulletinBoard).Error
	if err != nil {
		zapLog.SugarLogger.Debug(err)
		ctx.JSON(500, gin.H{
			"message": "create error",
		})
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(200, gin.H{
		"message": "Ok",
		"text":    bulletinBoard.Message,
	})
}

func GetTimeRemind(ctx *gin.Context) {
	var GetTimeRemind define.TimeRemind
	var ReadDB define.TimeRemind
	err := ctx.ShouldBind(&GetTimeRemind)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	err = initDatabase.DB.Model(&define.TimeRemind{}).Where("time=?,room_id=?)", GetTimeRemind.Time, GetTimeRemind.RoomId).Find(&ReadDB).Error
	if err != nil {
		zapLog.SugarLogger.Debug(err)
		ctx.JSON(500, gin.H{
			"message": "find error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"time":    ReadDB.Time,
		"message": ReadDB.Message,
	})
}

// PostDuty OnDuty 值日生功能
func PostDuty(ctx *gin.Context) {
	tx := initDatabase.DB.Begin()
	var Order define.GetDuty
	err := ctx.ShouldBind(&Order)
	people := []string{Order.P1, Order.P2, Order.P3, Order.P4}
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	for _, val := range people {
		var Duty = define.DutyOrder{
			Name:   val,
			RoomID: Order.RoomID,
		}
		err := tx.Model(&define.DutyOrder{}).Create(&Duty).Error
		if err != nil {
			zapLog.SugarLogger.Debug("create error", err)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	ctx.JSON(200, gin.H{
		"message": "query Ok",
	})
}

//func GetDuty(ctx *gin.Context) {
//     var Get struct{
//		 RoomID  string `json:"roomID"`
//		 Today   uint  `json:"today"`
//	 }
//
//}

// AddPartner  添加室友功能
func AddPartner(ctx *gin.Context) {
	tx := initDatabase.DB.Begin()
	var NewPartner define.AddPartner
	var Partner define.UserLogin
	var ExistUser define.UserLogin
	err := ctx.ShouldBind(&NewPartner)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	err = tx.Model(&define.UserLogin{}).Where("user_name=?", NewPartner.UserName).Find(&Partner).Error
	if err != nil {
		zapLog.SugarLogger.Debug("find error", err)
		ctx.JSON(500, err)
		return
	}
	if define.CheckAmt(Partner) {
		zapLog.SugarLogger.Info("no such user")
		ctx.JSON(400, gin.H{
			"message": "用户不存在",
		})
		return
	}
	err = tx.Model(&define.UserLogin{}).Where("user_name=?", NewPartner.UserName).Find(&ExistUser).Error
	if err != nil {
		zapLog.SugarLogger.Debug("find error", err)
		ctx.JSON(500, gin.H{
			"message": "server error",
			"error":   err,
		})
		return
	}
	var ID string
	if ExistUser.RoomID == "" {
		ID = GetRoomID(6)
		err := tx.Model(&ExistUser).Select("room_id").Updates(define.UserLogin{RoomID: ID}).Error
		if err != nil {
			zapLog.SugarLogger.Debug("Update error")
			ctx.JSON(500, err)
			tx.Rollback()
			return
		}

		err = tx.Model(&Partner).Select("room_id").Updates(define.UserLogin{RoomID: ID}).Error
		if err != nil {
			zapLog.SugarLogger.Debug("Update error")
			ctx.JSON(500, err)
			tx.Rollback()
			return
		}
	} else {
		err = tx.Model(&Partner).Select("room_id").Updates(define.UserLogin{RoomID: ExistUser.RoomID}).Error
		if err != nil {
			zapLog.SugarLogger.Debug("Update error")
			ctx.JSON(500, err)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	ctx.JSON(200, gin.H{
		"message": "添加成功",
		"roomId":  ID,
	})
}
func GetRoomID(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			zapLog.SugarLogger.Error(err)
		}
	}
	return sb.String()
}
func GetPartnerMessage(ctx *gin.Context) {
	var User define.UserLogin
	var UserExist define.UserLogin
	var Users []define.UserLogin
	err := ctx.ShouldBind(&User)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	err = initDatabase.DB.Model(&User).Find(&UserExist).Error
	if err != nil {
		zapLog.SugarLogger.Debug("find error", err)
		ctx.JSON(500, err)
		return

	}
	err = initDatabase.DB.Model(&define.UserLogin{}).Where("room_id=?", UserExist.RoomID).Find(&Users).Error
	if err != nil {

		zapLog.SugarLogger.Debug("find error", err)
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, Users)
}

//改变状态与获取状态

func ChangeStates(ctx *gin.Context) {
	tx := initDatabase.DB.Begin()
	var State struct {
		UserName uint   `json:"UserName"`
		State    string `json:"State"`
	}
	err := ctx.ShouldBind(&State)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	err = tx.Model(&define.UserLogin{}).Where("user_name=?", State.UserName).Update("state", State.State).Error
	if err != nil {
		zapLog.SugarLogger.Debug("update error", err)
		ctx.JSON(500, err)
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(200, gin.H{
		"message": "query OK",
	})
}

func GetStateMessage(ctx *gin.Context) {
	var User struct {
		UserName uint
	}
	var Message define.UserLogin
	err := ctx.ShouldBind(&User)
	if err != nil {
		zapLog.SugarLogger.Info(err)
		ctx.JSON(400, gin.H{
			"message": "bind register err please check your post json",
			"error":   err,
		})
		return
	}
	err = initDatabase.DB.Model(&define.UserLogin{}).Where("user_name=?", User.UserName).Find(&Message).Error
	if err != nil {
		zapLog.SugarLogger.Debug("find error", err)
		ctx.JSON(500, gin.H{
			"message": "server find error",
			"error":   err,
		})
	}
	ctx.JSON(200, Message.State)
}
