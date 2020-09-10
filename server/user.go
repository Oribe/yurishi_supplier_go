package server

import (
	"crypto/hmac"
	"errors"
	"manufacture_supplier_go/model"
	"manufacture_supplier_go/util"
	"strings"
)

// Login 登陆
func Login(username string, password string) (*model.User, error) {
	var user model.User

	err := model.UserQueryRow(&user, username)
	if err != nil {
		return nil, err
	}

	username = strings.TrimSpace(username)
	signPassword := util.SignPassword(username, password)
	if hmac.Equal(signPassword, []byte(user.Password)) {
		return nil, errors.New("密码或者用户名错误")
	}
	return &user, nil
}
