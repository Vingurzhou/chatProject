package model

import "time"

type Community struct {
	Id       int64     `gorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Name     string    `gorm:"varchar(30)" form:"name" json:"name"`
	Ownerid  int64     `gorm:"bigint(20)" form:"ownerid" json:"ownerid"`
	Icon     string    `gorm:"varchar(250)" form:"icon" json:"icon"`
	Cate     int       `gorm:"int(11)" form:"cate" json:"cate"`
	Memo     string    `gorm:"varchar(120)" form:"memo" json:"memo"`
	Createat time.Time `gorm:"datetime" form:"createat" json:"createat"`
}

const (
	COMMUNITY_CATE_COM = 0x01
)
