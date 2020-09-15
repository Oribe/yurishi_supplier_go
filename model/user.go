package model

import (
	"encoding/json"

	"github.com/guregu/null"
)

// UserModel 用户类型
type UserModel struct {
	ID         int         `json:"id" db:"id"`                 // 用户编号
	UserName   null.String `json:"userName" db:"userName"`     // 用户名
	Email      null.String `json:"email" db:"email"`           //	邮箱
	Mobile     null.String `json:"mobile" db:"mobile"`         // 手机号
	Supplier   null.String `json:"supplier" db:"supplier"`     // 供应商
	Contact    null.String `json:"contact" db:"contact"`       // 联系人
	SupplierID null.String `json:"supplierId" db:"supplierId"` // 供应商编号
	Remark     null.String `json:"remark" db:"remark"`         // 备注
	Password   null.String `json:"passwrod" db:"password"`     // 密码
}

// User 用户信息
type User struct {
	ID         int    `json:"id" db:"id"`                 // 用户编号
	UserName   string `json:"userName" db:"userName"`     // 用户名
	Email      string `json:"email" db:"email"`           //	邮箱
	Mobile     string `json:"mobile" db:"mobile"`         // 手机号
	Supplier   string `json:"supplier" db:"supplier"`     // 供应商
	Contact    string `json:"contact" db:"contact"`       // 联系人
	SupplierID string `json:"supplierId" db:"supplierId"` // 供应商编号
	Remark     string `json:"remark" db:"remark"`         // 备注
}

// Brief ...
func (u *UserModel) Brief() *User {
	userByte, _ := json.Marshal(u)
	var user User
	json.Unmarshal(userByte, &user)
	return &user
}

// UserQueryRow 查询单条用户信息
func UserQueryRow(user *UserModel, userName string) error {
	sql := `SELECT 
						id, userName, email, mobile, supplier, contact, supplierId, password, remark 
					FROM 
						users
					WHERE 
						userName = ?`
	err := db.Get(user, sql, userName)
	if err != nil {
		return err
	}
	user.Email.Value()
	return nil
}
