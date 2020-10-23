package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// DrillIndexableBody 可转位钻头刀体
type DrillIndexableBody struct {
	model.CutterBase
	BMC   null.String `json:"BMC"`   // 刀体材料代码
	CCMS  null.String `json:"CCMS"`  // 机床侧接口代码
	CCTMS null.String `json:"CCTMS"` // 机床侧接口类型代码
	CZCMS null.String `json:"CZCMS"` // 机床侧接口尺寸代码
	IICP  null.String `json:"IICP"`  // 周边刀片接口代码
	IISCP null.String `json:"IISCP"` // 周边刀片接口标准代码
	ISCP  null.String `json:"ISCP"`  // 周边刀片尺寸代码
	IICC  null.String `json:"IICC"`  // 中心刀片接口代码
	IISCC null.String `json:"IISCC"` // 中心刀片接口标准代码
	ISCC  null.String `json:"ISCC"`  // 中心刀片尺寸代码
	HAND  null.String `json:"HAND"`  // 切削方向
	DC    null.Int    `json:"DC"`    // 切削直径
	LF    null.Int    `json:"LF"`    // 功能长度
	LU    null.Int    `json:"LU"`    // 可用长度
	CSP   null.Int    `json:"CSP"`   // 冷却液供给特性
	DPC   null.Int    `json:"DPC"`   // 阻尼减振属性
}

func drillIndexableBodyDataInitial(body *DrillIndexableBody) *model.DrillIndexableBody {
	row := &model.DrillIndexableBody{
		BMC:   body.BMC,
		CCMS:  body.CCMS,
		CCTMS: body.CCTMS,
		CZCMS: body.CZCMS,
		IICP:  body.IICP,
		IISCP: body.IISCP,
		ISCP:  body.ISCP,
		IICC:  body.IICC,
		IISCC: body.IISCC,
		ISCC:  body.ISCC,
		HAND:  body.HAND,
		DC:    body.DC,
		LF:    body.LF,
		LU:    body.LU,
		CSP:   body.CSP,
		DPC:   body.DPC,
	}
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.OrderNumber = body.OrderNumber
	row.Manufacturer = body.Manufacturer
	row.Submitter = body.Submitter
	return row
}

// DrillIndexableBodyQueryWidthOrderNumber 查询
func DrillIndexableBodyQueryWidthOrderNumber(orderNumber, userID string) ([]model.DrillIndexableBody, error) {
	rows := []model.DrillIndexableBody{}
	err := model.DrillIndexableBodyQueryWithOrderNumber(&rows, orderNumber, userID)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// DrillIndexableBodyInsert 新增
func DrillIndexableBodyInsert(body *DrillIndexableBody) (int64, error) {
	rows := drillIndexableBodyDataInitial(body)
	lastInsertID, err := rows.InsertToDB()
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}
