package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// GroovingCuttingExternalBody 外圆切槽切断刀体
type GroovingCuttingExternalBody struct {
	model.CutterBase
	TSYC      null.String // 刀具形式代码
	BMC       null.String // 刀体材料代码
	CCMS      null.String // 机床侧接口代码
	CCTMS     null.String // 机床侧接口类型代码
	CZCMS     null.String // 机床侧接口尺寸代码
	IIC       null.String // 刀片接口代码
	IISC      null.String // 刀片接口标准代码
	ISC       null.String // 刀垫尺寸代码
	HAND      null.String // 切削方向
	LF        null.Int    // 功能长度
	CUTDIA    null.Int    // 最大工件切断直径
	CDX       null.Int    // 最大切削深度
	DAXX      null.Int    // 端面切槽外圈最大直径
	DAXN      null.Int    // 端面切槽外圈最小直径
	CSP       null.Int    // 冷却液供给特性
	DPC       null.Int    // 阻尼减振属性
	PART      null.Int    // 切断
	GRVEX     null.Int    // 外圆切槽
	GRVEXFC   null.Int    // 端面切槽（外部）
	GRVEXPR   null.Int    // 外圆轮廓
	GRVEXTR   null.Int    // 外圆车削
	GRVEXTRBK null.Int    // 外圆背车
	GRVEXUND  null.Int    // 外圆掏槽
}

// GroovingCuttingExternalBodyQueryWithOrderNumber 查询
func GroovingCuttingExternalBodyQueryWithOrderNumber(orderNumber, userID string) ([]model.GroovingCuttingExternalBody, error) {
	list := []model.GroovingCuttingExternalBody{}
	err := model.GroovingCuttingExternalBodyQueryWithOrderNumber(&list, orderNumber, userID)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func groovingCuttingExternalBodyRowInitial(body GroovingCuttingExternalBody) model.GroovingCuttingExternalBody {
	row := model.GroovingCuttingExternalBody{
		TSYC:      body.TSYC,      // 刀具形式代码
		BMC:       body.BMC,       // 刀体材料代码
		CCMS:      body.CCMS,      // 机床侧接口代码
		CCTMS:     body.CCTMS,     // 机床侧接口类型代码
		CZCMS:     body.CZCMS,     // 机床侧接口尺寸代码
		IIC:       body.IIC,       // 刀片接口代码
		IISC:      body.IISC,      // 刀片接口标准代码
		ISC:       body.ISC,       // 刀垫尺寸代码
		HAND:      body.HAND,      // 切削方向
		LF:        body.LF,        // 功能长度
		CUTDIA:    body.CUTDIA,    // 最大工件切断直径
		CDX:       body.CDX,       // 最大切削深度
		DAXX:      body.DAXX,      // 端面切槽外圈最大直径
		DAXN:      body.DAXN,      // 端面切槽外圈最小直径
		CSP:       body.CSP,       // 冷却液供给特性
		DPC:       body.DPC,       // 阻尼减振属性
		PART:      body.PART,      // 切断
		GRVEX:     body.GRVEX,     // 外圆切槽
		GRVEXFC:   body.GRVEXFC,   // 端面切槽（外部）
		GRVEXPR:   body.GRVEXPR,   // 外圆轮廓
		GRVEXTR:   body.GRVEXTR,   // 外圆车削
		GRVEXTRBK: body.GRVEXTRBK, // 外圆背车
		GRVEXUND:  body.GRVEXUND,  // 外圆掏槽
	}
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.OrderNumber = body.OrderNumber
	row.Manufacturer = body.Manufacturer
	row.Submitter = body.Submitter
	return row
}

// GroovingCuttingExternalBodyInsertToDB 写入数据库
func GroovingCuttingExternalBodyInsertToDB(body GroovingCuttingExternalBody) (int64, error) {
	row := groovingCuttingExternalBodyRowInitial(body)
	return row.InsertToDB()
}
