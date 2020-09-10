package util

import (
	"errors"
	"fmt"
	"reflect"
)

// SwitchSqlNull 转换
func SwitchSqlNull(sqlNull interface{}) error {
	t := reflect.TypeOf(sqlNull)
	if t.Kind() != reflect.Struct {
		// err := errors.New("must be pass a struct, but get a %v", t.Kind())
		err := errors.New("must be pass a struct, but get a %v")
		return err
	}
	fmt.Println(sqlNull)
	return nil
}
