package define

import "reflect"

// UserLogin 登录模块 待修改
type UserLogin struct {
	UserName uint   `json:"UserName"`
	PassWord string `json:"PassWord"`
	RoomID   int
}

type Token struct {
	Id    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Token string `gorm:"type:varchar(1000);not null"`
	UID   uint
}

type UserMessage struct {
	PeopleNumber uint
	Name         string `json:"Name"`
	State        string `json:"State"`
}

// BulletinBoard 公告版面
type BulletinBoard struct {
	Time    string
	Message string `gorm:"varchar(1000)"`
}

type loginConfig interface {
	IsAmt() bool
}

func (ul UserLogin) IsAmt() bool {
	return reflect.DeepEqual(ul, UserLogin{})
}

func (um UserMessage) IsAmt() bool {
	return reflect.DeepEqual(um, UserMessage{})
}
func CheckAmt(config loginConfig) bool {
	return config.IsAmt()
}
