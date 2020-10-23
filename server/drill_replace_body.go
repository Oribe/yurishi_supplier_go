package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// DrillReplaceBody 可更换钻削刀头刀体
type DrillReplaceBody struct {
	model.CutterBase
	BMC   null.String // 刀体材料代码
	CCMS  null.String // 机床侧接口代码
	CCTMS null.String // 机床侧接口类型代码
	CZCMS null.String // 机床侧接口尺寸代码
	IIC   null.String // 刀片接口代码
	IISC  null.String // 刀片接口标准代码
	ISC   null.String // 刀片尺寸代码
	LF    null.Int    // 功能长度
	LU    null.Int    // 可用长度
	CSP   null.Int    // 冷却液供给特性
	DPC   null.Int    // 阻尼减振属性
}

func drillReplaceBodyRowInitial(body DrillReplaceBody) model.DrillReplaceBody {
	row := model.DrillReplaceBody{
		BMC:   body.BMC,   // 刀体材料代码
		CCMS:  body.CCMS,  // 机床侧接口代码
		CCTMS: body.CCTMS, // 机床侧接口类型代码
		CZCMS: body.CZCMS, // 机床侧接口尺寸代码
		IIC:   body.IIC,   // 刀片接口代码
		IISC:  body.IISC,  // 刀片接口标准代码
		ISC:   body.ISC,   // 刀片尺寸代码
		LF:    body.LF,    // 功能长度
		LU:    body.LU,    // 可用长度
		CSP:   body.CSP,   // 冷却液供给特性
		DPC:   body.DPC,   // 阻尼减振属性
	}
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.Manufacturer = body.Manufacturer
	row.OrderNumber = body.OrderNumber
	row.Submitter = body.Submitter
	return row
}

// DrillReplaceBodyInsert 添加
func DrillReplaceBodyInsert(body DrillReplaceBody) (int64, error) {
	row := drillReplaceBodyRowInitial(body)
	lastID, err := row.InsertToDB()
	if err != nil {
		return lastID, err
	}
	return lastID, nil
}

// DrillReplaceBodyQueryWithOrderNumber 查询
func DrillReplaceBodyQueryWithOrderNumber(orderNumber, userID string) ([]model.DrillReplaceBody, error) {
	rows := []model.DrillReplaceBody{}
	err := model.DrillReplaceBodyQueryWithOrderNumber(&rows, orderNumber, userID)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
