package server

import (
	"testing"
)

func TestDrillGunQueryWithOrder(t *testing.T) {
	_, err := DrillGunQueryWithOrder("", 1)
	if err != nil {
		t.Errorf("枪钻查询失败: %v", err.Error())
	}
}
