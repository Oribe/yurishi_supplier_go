package server

import (
	"crypto/hmac"
	"errors"
	"fmt"
	"manufacture_supplier_go/model"
	"manufacture_supplier_go/util"
	"strings"
)

// Login 登陆
func Login(username string, password string) (*model.UserModel, error) {
	var user model.UserModel

	err := model.UserQueryRow(&user, username)
	if err != nil {
		return nil, err
	}

	username = strings.TrimSpace(username)
	signPassword := util.SignPassword(username, password)
	pwd, err := user.Password.MarshalText()
	if err != nil {
		return nil, err
	}
	if hmac.Equal(signPassword, pwd) {
		return nil, errors.New("密码或者用户名错误")
	}
	fmt.Println(user.Email.Value())
	return &user, nil
}
