package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// DrillIntegral 整体式钻削
type DrillIntegral struct {
	model.CutterBase
	TPCODE null.String `json:"TPCODE"` // 刀具型号
	GRADE  null.String `json:"GRADE"`  // 材质等级代码
	SDTP   null.String `json:"SDTP"`   // 钻头类型
	ZEEP   null.Int    `json:"ZEEP"`   // 切削刃数
	CZCMS  null.String `json:"CZCMS"`  // 机床侧接口代码
	TPCON  null.String `json:"TPCON"`  // 接口类型
	DCON   null.Int    `json:"DCON"`   // 连接直径
	HAND   null.String `json:"HAND"`   // 切削方向
	DC     null.Int    `json:"DC"`     // 切削直径
	OAL    null.Int    `json:"OAL"`    // 刀具总长
	LU     null.Int    `json:"LU"`     // 可用长度
	CSP    null.Int    `json:"CSP"`    // 冷却液供给特征
	STN    null.Int    `json:"STN"`    // 阶梯数
	SDL1   null.Int    `json:"SDL_1"`  // 直径长度-1
	STA2   null.Int    `json:"STA_2"`  // 阶梯倒角-2
	DC2    null.Int    `json:"DC_2"`   // 阶梯直径-   2
	SD2    null.Int    `json:"SD_2"`   // 阶梯长度-2
	SDL2   null.Int    `json:"SDL_2"`  // 直径长度-2
	STA3   null.Int    `json:"STA_3"`  // 阶梯倒角-3
	DC3    null.Int    `json:"DC_3"`   // 阶梯直径-3
	SD3    null.Int    `json:"SD_3"`   // 阶梯长度-3
	P      null.Int    `json:"P"`      // P
	M      null.Int    `json:"M"`      // M
	K      null.Int    `json:"K"`      // K
	N      null.Int    `json:"N"`      // N
	S      null.Int    `json:"S"`      // S
	H      null.Int    `json:"H"`      // H
	O      null.Int    `json:"O"`      // O
}

func drillIntegralRowInitail(body DrillIntegral) model.DrillIntegral {
	row := model.DrillIntegral{
		TPCODE: body.TPCODE,
		GRADE:  body.GRADE,
		SDTP:   body.SDTP,
		ZEEP:   body.ZEEP,
		CZCMS:  body.CZCMS,
		TPCON:  body.TPCON,
		DCON:   body.DCON,
		HAND:   body.HAND,
		DC:     body.DC,
		OAL:    body.OAL,
		LU:     body.LU,
		CSP:    body.CSP,
		STN:    body.STN,
		SDL1:   body.SDL1,
		STA2:   body.STA2,
		DC2:    body.DC2,
		SD2:    body.SD2,
		SDL2:   body.SDL2,
		STA3:   body.STA3,
		DC3:    body.DC3,
		SD3:    body.SD3,
		P:      body.P,
		M:      body.M,
		K:      body.K,
		N:      body.N,
		S:      body.S,
		H:      body.H,
		O:      body.O,
	}
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.OrderNumber = body.OrderNumber
	row.Manufacturer = body.Manufacturer
	row.Submitter = body.Submitter
	return row
}

// DrillIntegralQueryWithOrderNumber 查询
func DrillIntegralQueryWithOrderNumber(orderNumber, userID string) ([]model.DrillIntegral, error) {
	list := []model.DrillIntegral{}
	err := model.DrillIntegralQueryWithOrderNumber(&list, orderNumber, userID)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// DrillIntegralInsertToDB 新增
func DrillIntegralInsertToDB(body DrillIntegral) (int64, error) {
	row := drillIntegralRowInitail(body)
	lastID, err := row.InsertToDB()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}
