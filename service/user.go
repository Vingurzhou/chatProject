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

	salt := fmt.Sprintf("%06d", rand.Int31n(10000))
	creatat := time.Now()
	token := fmt.Sprintf("%08d", rand.Int31())

	err = DB.Exec("INSERT INTO users (mobile,passwd,avatar,sex,nickname,salt,online,token,memo,createat)  VALUES  (?,?,?,?,?,?,?,?,?,?)", mobile, tmp.Passwd, avatar, sex, nickname, salt, tmp.Online, token, tmp.Memo, creatat).Error
	//curl http://127.0.0.1:8000/user/register -d "mobile=19952429930&passwd=123456"
	return tmp, err
}
