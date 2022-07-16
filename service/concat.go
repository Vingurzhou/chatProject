package service

import (
	"chatProject/model"
	"errors"
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
	return coms
}

func (service *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	// num, err := DB.Count(&com)

	// if num > 5 {
	// 	err = errors.New("一个用户最多只能创见5个群")
	// 	return com, err
	// } else {
	comm.Createat = time.Now()
	session := DB.Session(&gorm.Session{})
	session.Begin()
	err = DB.Exec("INSERT INTO contacts (ownerid,dstobj,cate,createat)  VALUES  (?,?,?,?)", comm.Ownerid, comm.Id, model.CONCAT_CATE_COMUNITY, time.Now()).Error
	if err != nil {
		session.Rollback()
		return com, err
	}
	err = DB.Exec("INSERT INTO contacts (ownerid,dstobj,cate,createat)  VALUES  (?,?,?,?)", comm.Ownerid, comm.Id, model.CONCAT_CATE_COMUNITY, time.Now()).Error
	if err != nil {
		session.Rollback()
	} else {
		session.Commit()
	}
	return com, err
	// }
}
