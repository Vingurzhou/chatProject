package model

import "time"

const (
	CONCAT_CATE_USER     = 0x01
	CONCAT_CATE_COMUNITY = 0x02
)

type Contact struct {
	Id       int64     `gorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Ownerid  int64     `gorm:"bigint(20)" form:"ownerid" json:"ownerid"` // 记录是谁的
	Dstobj   int64     `gorm:"bigint(20)" form:"dstobj" json:"dstobj"`   // 对端信息
	Cate     int       `gorm:"int(11)" form:"cate" json:"cate"`          // 什么类型
	Memo     string    `gorm:"varchar(120)" form:"memo" json:"memo"`     // 备注
	Createat time.Time `gorm:"datetime" form:"createat" json:"createat"` // 创建时间
}
