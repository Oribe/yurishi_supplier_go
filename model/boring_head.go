package model

import (
	"fmt"
)

// BoringHead 镗头
type BoringHead struct {
	Category     int    `db:"category" json:"category"`
	SubCategory  int    `db:"subCategory" json:"subCategory"`
	Manufacturer string `db:"manufacturer" json:"manufacturer"`
	OrderNumber  string `db:"orderNumber" json:"orderNumber"`
	BMC          string
	CCMS         string
	CSCMS        string
	CCTMS        string
	CZCMS        string
	CCCRT        string
	CSCCRT       string
	CASC         string
	IIC          string
	SC           string
	ANC          string
	ICDC         string
	THCKC        string
	HAND         string
	DCX          int
	DCN          int
	LF           int
	CSP          int
	INBORROUGH   int
	INBORFINE    int
	INBORSTEP    int
	INBORBK      int
	Submitter    string `db:"submitter" json:"submitter"`
}

// InsertToDB 插入数据库
func (b *BoringHead) InsertToDB() (int64, error) {
	sql := `INSERT INTO 
						boring_heads
						(category, subCategory, manufacturer, orderNumber, BMC, CCMS, CSCMS, CCTMS, CZCMS, CCCRT, CSCCRT, CASC, IIC, SC, ANC, ICDC, THCKC, HAND, DCX, DCN, LF, CSP, INBORROUGH, INBORFINE, INBORSTEP, INBORBK, submitter)
					VALUES
						(:category, :subCategory, :manufacturer, :orderNumber, :BMC, :CCMS, :CSCMS, :CCTMS, :CZCMS, :CCCRT, :CSCCRT, :CASC, :IIC, :SC, :ANC, :ICDC, :THCKC, :HAND, :DCX, :DCN, :LF, :CSP, :INBORROUGH, :INBORFINE, :INBORSTEP, :INBORBK, :submitter)`
	result, err := db.NamedExec(sql, b)
	if err != nil {
		error := fmt.Errorf("insert into table boring_head failed: %v", err.Error())
		return 0, error
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("get boring_head lastInsertId failed: %v", err.Error())
		return 0, error
	}
	return lastInsertID, nil
}

// BoringHeadQueryWithOrderNumber 查询
func BoringHeadQueryWithOrderNumber(boringHeadList *[]BoringHead, orderNumber string, userID string) error {
	sql := `SELECT 
						id, category, subCategory, manufacturer, orderNumber, BMC, CCMS, CSCMS, CCTMS, CZCMS, CCCRT, CSCCRT, CASC, IIC, SC, ANC, ICDC, THCKC, HAND, DCX, DCN, LF, CSP, INBORROUGH, INBORFINE, INBORSTEP, INBORBK, submitter
					FROM
						boring_heads
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						submitter = ?`
	stmt, err := db.Preparex(sql)

	if err != nil {
		error := fmt.Errorf("sql prepare failed: %v", err.Error())
		return error
	}

	err = stmt.Select(boringHeadList, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("boring_head query failed: %v", err.Error())
		return error
	}

	return nil
}
