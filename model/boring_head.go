package model

import (
	"fmt"

	"github.com/guregu/null"
)

// BoringHead 镗头
type BoringHead struct {
	CutterBase
	BMC        null.String
	CCMS       null.String
	CSCMS      null.String
	CCTMS      null.String
	CZCMS      null.String
	CCCRT      null.String
	CSCCRT     null.String
	CASC       null.String
	IIC        null.String
	SC         null.String
	ANC        null.String
	ICDC       null.String
	THCKC      null.String
	HAND       null.String
	DCX        null.Int
	DCN        null.Int
	LF         null.Int
	CSP        null.Int
	INBORROUGH null.Int
	INBORFINE  null.Int
	INBORSTEP  null.Int
	INBORBK    null.Int
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
