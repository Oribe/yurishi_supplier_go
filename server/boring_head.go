package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// BoringHead 镗头
type BoringHead struct {
	model.CutterBase
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

func dataInit(body *BoringHead) model.BoringHead {
	row := model.BoringHead{
		BMC:        body.BMC,
		CCMS:       body.CCMS,
		CSCMS:      body.CSCMS,
		CCTMS:      body.CCTMS,
		CZCMS:      body.CZCMS,
		CCCRT:      body.CCCRT,
		CSCCRT:     body.CSCCRT,
		CASC:       body.CASC,
		IIC:        body.IIC,
		SC:         body.SC,
		ANC:        body.ANC,
		ICDC:       body.ICDC,
		THCKC:      body.THCKC,
		HAND:       body.HAND,
		DCX:        body.DCX,
		DCN:        body.DCN,
		LF:         body.LF,
		CSP:        body.CSP,
		INBORROUGH: body.INBORROUGH,
		INBORFINE:  body.INBORFINE,
		INBORSTEP:  body.INBORSTEP,
		INBORBK:    body.INBORBK,
	}
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.Manufacturer = body.Manufacturer
	row.OrderNumber = body.OrderNumber
	row.Submitter = body.Submitter
	return row
}

// BoringHeadInsert 新增
func BoringHeadInsert(body *BoringHead) (int64, error) {
	boringHead := dataInit(body)
	lastInsertID, err := boringHead.InsertToDB()
	if err != nil {
		return 0, err
	}
	return lastInsertID, err
}

// BoringHeadQuery 查询
func BoringHeadQuery(orderNumber string, userID string) ([]model.BoringHead, error) {
	boringHeadList := []model.BoringHead{}
	err := model.BoringHeadQueryWithOrderNumber(&boringHeadList, orderNumber, userID)
	if err != nil {
		return nil, err
	}
	return boringHeadList, nil
}
