package server

import (
	"crypto/hmac"
	"errors"
	"fmt"
	"manufacture_supplier_go/model"
	"manufacture_supplier_go/util"
	"manufacture_supplier_go/util/jwt"
	"strings"

	"github.com/patrickmn/go-cache"
)

// Login 登陆
func Login(username, password, ip string) (*model.UserModel, string, error) {
	var user model.UserModel

	err := model.UserQueryRow(&user, username)
	if err != nil {
		return nil, "", errors.New("当前账户不存在")
	}

	username = strings.TrimSpace(username)
	signPassword := util.SignPassword(username, password)

	pwd, err := user.Password.MarshalText()
	if err != nil { // 用户输入密码加密失败
		return nil, "", errors.New("密码或者用户名错误")
	}

	// 验证数据库中的用户密码是否与用户输入的密码相同
	if hmac.Equal(signPassword, pwd) {
		return nil, "", errors.New("密码或者用户名错误")
	}

	// 登陆成功，生成token
	token, err := jwt.CreateToken(username, ip, user.ID)

	// 加入缓存
	cache.Cache.Get()

	if err != nil {
		return nil, "", fmt.Errorf("生成Token失败：%v", err)
	}
	return &user, token, nil
}
