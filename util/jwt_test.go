package util

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken("Yurishi", "http://localhost", "192.168.1.123")
	if err != nil {
		t.Errorf("生成Token失败: %v", err.Error())
	}
	fmt.Println("token:", token)
}

func TestVerifyToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkb21haW4iOiJodHRwOi8vbG9jYWxob3N0IiwiaXAiOiIxOTIuMTY4LjEuMTIzIiwidXNlcm5hbWUiOiJZdXJpc2hpIn0.pGVZoTCqMifmw5K73Q0U2gWV8P0NgRPck6ws6lg1Eas"

	ok := VerifyToken(token, "Yurishi", "http://localhost", "192.168.1.123")
	if !ok {
		t.Errorf("")
	}
}
