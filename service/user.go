package service

import "chatProject/model"

type UserService struct {
}

func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	// tmp := model.User{}
	// DB.Where("mobile = ?", mobile).Find(&tmp)
}
