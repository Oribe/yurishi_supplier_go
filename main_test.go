package main

import (
	"encoding/json"
	"manufacture_supplier_go/route"
	"manufacture_supplier_go/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

var router = route.NewRouter()
var w = httptest.NewRecorder()

// TestLogin 登陆测试
func TestLogin(t *testing.T) {

	type user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	body := user{"15172545243", "123456"}

	str, _ := json.Marshal(body)
	r := strings.NewReader(string(str))

	req, _ := http.NewRequest("POST", "/tool/interface/login", r)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// TestOrderSearch 订单查询
func TestOrderSearch(t *testing.T) {
	req, _ := http.NewRequest("GET", "/tool/interface/order/query", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// TestDrillIndexableBodyQueryWidthOrderNumber 可转位钻头刀体查询
func TestDrillIndexableBodyQueryWidthOrderNumber(t *testing.T) {
	_, err := server.DrillIndexableBodyQueryWidthOrderNumber("", "1")
	if err != nil {
		t.Errorf("可转位钻头刀体查询失败：%v", err.Error())
	}
}
