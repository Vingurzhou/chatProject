package service

import (
	"chatProject/model"
	"chatProject/util"
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
	if !util.ValidatePasswd(plainpwd, tmp.Salt, tmp.Passwd) {
		return tmp, errors.New("密码不正确")
	}
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)

	DB.Exec("UPDATE users SET token = ? WHERE id= ? ", token, tmp.Id)
	return tmp, nil
}
func (s *UserService) Register(
	mobile,
	plainpwd,
	nickname,
	avatar, sex string) (user model.User, err error) {

	tmp := model.User{}
	err = DB.Raw("select *from users where mobile = ?", mobile).Scan(&tmp).Error
	if err != nil {
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已经注册")
	}

	salt := fmt.Sprintf("%06d", rand.Int31n(10000))
	passwd := util.MakePasswd(plainpwd, salt)
	creatat := time.Now()
	token := fmt.Sprintf("%08d", rand.Int31())
	err = DB.Exec("INSERT INTO users (mobile,passwd,avatar,sex,nickname,salt,online,token,memo,createat)  VALUES  (?,?,?,?,?,?,?,?,?,?)", mobile, passwd, avatar, sex, nickname, salt, tmp.Online, token, tmp.Memo, creatat).Error
	//curl http://127.0.0.1:8000/user/register -d "mobile=19952429930&passwd=123456"
	return tmp, err
}
