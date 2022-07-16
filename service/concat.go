package service

import (
	"chatProject/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ConcatService struct{}

func (service *ConcatService) AddFriend(userid, dstid int64) error {
	if userid == dstid {
		return errors.New("不能添加自己为好友啊")
	}
	tmp := model.Contact{}
	DB.Raw("select *from users Where ownerid = ? And dstid = ? And cate = ?", userid, dstid, model.CONCAT_CATE_USER).Scan(&tmp)
	if tmp.Id > 0 {
		return errors.New("该用户已经被添加过啦")
	}
	session := DB.Session(&gorm.Session{})
	session.Begin()

	e2 := DB.Exec("INSERT INTO users (ownerid,dstobj,cate,memo,createat)  VALUES  (?,?,?,?,?,?,?,?,?,?)", userid, dstid, model.CONCAT_CATE_USER, time.Now()).Error
	e3 := DB.Exec("INSERT INTO users (ownerid,dstobj,cate,memo,createat)  VALUES  (?,?,?,?,?,?,?,?,?,?)", dstid, userid, model.CONCAT_CATE_USER, time.Now()).Error
	if e2 == nil && e3 == nil {
		//提交
		session.Commit()
		return nil
	} else {
		//回滚
		session.Rollback()
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}
