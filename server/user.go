package server

import (
	"errors"
	"fmt"
	"manufacture_supplier_go/cache"
	"manufacture_supplier_go/model"
	"manufacture_supplier_go/util"
	"manufacture_supplier_go/util/crypto"
	"strings"
)

// UserLoginInfo 用户登录信息（用户存到缓存中）
type UserLoginInfo struct {
	ID int
	IP string
}

// Login 登陆
func Login(username, password, ip string) (*model.UserModel, string, error) {
	var user model.UserModel

	err := model.UserQueryRow(&user, username)
	if err != nil {
		return nil, "", errors.New("当前账户不存在")
	}

	username = strings.TrimSpace(username)
	signPassword := crypto.SignPassword(username, password)

	pwd, err := user.Password.MarshalText()

	if err != nil { // 用户输入密码加密失败
		return nil, "", errors.New("密码或者用户名错误")
	}

	// 验证数据库中的用户密码是否与用户输入的密码相同
	// if !hmac.Equal(signPassword, pwd) {
	if signPassword != string(pwd) {
		return nil, "", errors.New("密码或者用户名错误")
	}

	// 登陆成功，生成token
	// token, err := jwt.CreateToken(username, ip, user.ID)
	// if err != nil {
	// 	return nil, "", fmt.Errorf("生成Token失败：%v", err)
	// }

	token := util.CreateToken()

	fmt.Println(token)

	loginInfo := UserLoginInfo{
		user.ID,
		ip,
	}

	// 保存到缓存
	ok := cache.Set(token, loginInfo, cache.DefaultExpiration)
	if !ok {
		return nil, "", fmt.Errorf("token保存失败")
	}

	return &user, token, nil
}

// LogOut 登出
func LogOut(token string) (ok bool) {
	ok = cache.Delete(token)
	return
}

// UserEdit 修改用户信息
func UserEdit(user model.User) error {
	err := model.UserUpdate(user)
	return err
}
