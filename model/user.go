package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// UserBrief 用户信息（无密码）
type UserBrief struct {
	ID         int            `json:"id" db:"id"`                 // 用户编号
	UserName   string         `json:"userName" db:"userName"`     // 用户名
	Email      sql.NullString `json:"email" db:"email"`           //	邮箱
	Mobile     sql.NullString `json:"mobile" db:"mobile"`         // 手机号
	Supplier   sql.NullString `json:"supplier" db:"supplier"`     // 供应商
	Contact    sql.NullString `json:"contact" db:"contact"`       // 联系人
	SupplierID sql.NullString `json:"supplierId" db:"supplierId"` // 供应商编号
	Remark     sql.NullString `json:"remark" db:"remark"`         // 备注
}

// User 用户类型
type User struct {
	UserBrief
	Password string `json:"passwrod" db:"password"` // 密码
}

// Brief ...
func (u *User) Brief() *UserBrief {
	userByte, _ := json.Marshal(u)
	var userbRrief UserBrief
	json.Unmarshal(userByte, &userbRrief)
	return &userbRrief
}

// UserQueryRow 查询单条用户信息
func UserQueryRow(user *User, userName string) error {
	sql := "SELECT id, userName, email, mobile, supplier, contact, supplierId, password, remark FROM users WHERE userName = ?"
	err := db.Get(user, sql, userName)
	if err != nil {
		fmt.Println("model.user:38", err)
		return err
	}
	return nil
}
