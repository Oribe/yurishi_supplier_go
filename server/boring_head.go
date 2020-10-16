package server

import (
	"manufacture_supplier_go/model"
)

// BoringHead 镗头
type BoringHead struct {
	Category     int    `json:"category"`
	SubCategory  int    `json:"subCategory"`
	Manufacturer string `json:"manufacturer"`
	OrderNumber  string `json:"orderNumber"`
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
	Submitter    string `json:"submitter"`
}

func dataInit(body *BoringHead) model.BoringHead {
	return model.BoringHead{
		Category:     body.Category,
		SubCategory:  body.SubCategory,
		Manufacturer: body.Manufacturer,
		OrderNumber:  body.OrderNumber,
		BMC:          body.BMC,
		CCMS:         body.CCMS,
		CSCMS:        body.CSCMS,
		CCTMS:        body.CCTMS,
		CZCMS:        body.CZCMS,
		CCCRT:        body.CCCRT,
		CSCCRT:       body.CSCCRT,
		CASC:         body.CASC,
		IIC:          body.IIC,
		SC:           body.SC,
		ANC:          body.ANC,
		ICDC:         body.ICDC,
		THCKC:        body.THCKC,
		HAND:         body.HAND,
		DCX:          body.DCX,
		DCN:          body.DCN,
		LF:           body.LF,
		CSP:          body.CSP,
		INBORROUGH:   body.INBORROUGH,
		INBORFINE:    body.INBORFINE,
		INBORSTEP:    body.INBORSTEP,
		INBORBK:      body.INBORBK,
		Submitter:    body.Submitter,
	}
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
