package jwt

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("Yurishi", "192.168.1.123", 1)
	if err != nil {
		t.Errorf("生成Token失败: %v", err.Error())
	}
	fmt.Println("token:", token)
}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBdWRpZW5jZSI6Ill1cmlzaGkiLCJJZCI6MSwiSXNzdWVkQXQiOjE2MDAyNDQ0NjYsIklzc3VlciI6ImJhaXFpYW8iLCJOb3RCZWZvcmUiOjE2MDAyNDQ0NjYsIlN1YmplY3QiOiJsb2dpbiIsImlwIjoiMTkyLjE2OC4xLjEyMyJ9.NJY1Ai83yiyjDrSl8H8w7k66nE9GIJWyym3sdyUj7bQ"

	ok := VerifyToken(token, "192.168.1.123")
	if !ok {
		t.Errorf("token验证失败")
	}
}
