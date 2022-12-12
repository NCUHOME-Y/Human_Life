package define

import (
	"reflect"
)

type Board struct {
	RoomID  string `json:"RoomID"`
	Message string `json:"Message" ,gorm:"varchar(200);not null"`
}

func (b Board) IsAmt() bool {
	return reflect.DeepEqual(b, Board{})
}

// UserLogin 登录模块 待修改
type UserLogin struct {
	UserName uint   `json:"UserName"`
	PassWord string `json:"PassWord"`
	RoomID   string
	State    string
}

type Token struct {
	Id    uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Token string `gorm:"type:varchar(1000);not null"`
	UID   uint
}

func (t Token) IsAmt() bool {
	return reflect.DeepEqual(t, Token{})
}

// TimeRemind 公告版面
type TimeRemind struct {
	Time    string `json:"time"`
	Message string ` json:"message" ,gorm:"varchar(1000)"`
	RoomId  string `json:"roomId"`
}

// DutyOrder OnDuty 值日生功能
type DutyOrder struct {
	id     uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name   string `gorm:"varchar(100);not null"`
	RoomID string
}
type NeedDuty struct {
	Mon  bool `json:"mon"`
	Tue  bool `json:"tue"`
	Wes  bool `json:"wes"`
	Thur bool `json:"thur"`
	Fri  bool `json:"fri"`
	Sat  bool `json:"sat"`
	Sun  bool `json:"sun"`
}
type PeopleOrder struct {
	P1 string `json:"p1"`
	P2 string `json:"p2"`
	P3 string `json:"p3"`
	P4 string `json:"p4"`
}

type GetDuty struct {
	Today  uint   `json:"today"`
	RoomID string `json:"roomID"`
	NeedDuty
	PeopleOrder
}

type AddPartner struct {
	UserName      uint `json:"UserName"`
	PartnerNumber uint `json:"partner_number"`
}

type loginConfig interface {
	IsAmt() bool
}

func (ul UserLogin) IsAmt() bool {
	return reflect.DeepEqual(ul, UserLogin{})
}

func CheckAmt(config loginConfig) bool {
	return config.IsAmt()
}
