package util

import (
	"errors"
	"fmt"
	"reflect"
)

// SwitchSQLNull 转换
func SwitchSQLNull(sqlNull interface{}) error {
	t := reflect.TypeOf(sqlNull)
	if t.Kind() != reflect.Struct {
		// err := errors.New("must be pass a struct, but get a %v", t.Kind())
		err := errors.New("must be pass a struct")
		return err
	}
	v := reflect.ValueOf(sqlNull)
	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i)
		fmt.Println(value)
	}
	fmt.Println(sqlNull)
	return nil
}
