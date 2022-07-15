package model

import "time"

const (
	SEX_WOMEN  = "W"
	SEX_MEN    = "M"
	SEX_UNKNOW = "U"
)

type User struct {
	Id       int64     `gorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Mobile   string    `gorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd   string    `gorm:"varchar(40)" form:"passwd" json:"-"`
	Avatar   string    `gorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string    `gorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname string    `gorm:"varchar(20)" form:"nickname" json:"nickname"`
	Salt     string    `gorm:"varchar(10)" form:"salt" json:"-"`
	Online   int       `gorm:"int(10)" form:"online" json:"online"`
	Token    string    `gorm:"varchar(40)" form:"token" json:"token"`
	Memo     string    `gorm:"varchar(140)" form:"memo" json:"memo"`
	Createat time.Time `gorm:"datetime" form:"createat" json:"createat"`
}
