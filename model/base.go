package model

import (
	"manufacture_supplier_go/database"
	"manufacture_supplier_go/util"
)

var db = database.NewDB()

var log = util.NewLog()

// CutterBase 刀具基本字段
type CutterBase struct {
	Category     int    `db:"category" json:"category"`
	SubCategory  int    `db:"subCategory" json:"subCategory"`
	Manufacturer string `db:"manufacturer" json:"manufacturer"`
	OrderNumber  string `db:"orderNumber" json:"orderNumber"`
	Submitter    string `db:"submitter" json:"submitter"`
}

type queryParam struct {
	orderNumber string
	userID      int
}
