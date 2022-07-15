package service

import (
	"chatProject/model"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct {
}

func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	tmp := model.User{}
	DB.Raw("select *from users where mobile = ?", mobile).Scan(&tmp)
	if tmp.Id == 0 {
		return tmp, errors.New("该用户不存在")
	}
	return tmp, nil
}
func (s *UserService) Register(
	mobile,
	plainpwd,
	nickname,
	avatar, sex string) (user model.User, err error) {

	tmp := model.User{}
	DB.Raw("select *from users where mobile = ?", mobile).Scan(&tmp)
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}

	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickname
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Createat = time.Now()
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	DB.Exec("")
	return tmp, err
}
