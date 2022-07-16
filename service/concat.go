package service

import (
	"chatProject/model"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ContactService struct{}

func (service *ContactService) AddFriend(userid, dstid int64) error {
	if userid == dstid {
		return errors.New("不能添加自己为好友啊")
	}
	tmp := model.Contact{}
	DB.Raw("select *from contacts Where ownerid = ? And dstobj = ? And cate = ?", userid, dstid, model.CONCAT_CATE_USER).Scan(&tmp)
	if tmp.Id > 0 {
		return errors.New("该用户已经被添加过啦")
	}
	session := DB.Session(&gorm.Session{})
	session.Begin()

	e2 := DB.Exec("INSERT INTO contacts (ownerid,dstobj,cate,createat)  VALUES  (?,?,?,?)", userid, dstid, model.CONCAT_CATE_USER, time.Now()).Error
	e3 := DB.Exec("INSERT INTO contacts (ownerid,dstobj,cate,createat)  VALUES  (?,?,?,?)", dstid, userid, model.CONCAT_CATE_USER, time.Now()).Error
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
func (service *ContactService) SearchFriend(userId int64) []model.User {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)

	DB.Raw("select *from contacts Where ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Scan(&conconts)
	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DB.Raw("select *from contacts where id in ?", objIds).Scan(&coms)
	fmt.Println(coms)
	return coms
}
