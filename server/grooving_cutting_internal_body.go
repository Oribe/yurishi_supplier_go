package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// GroovingCuttingInternalBody 內圆切槽切断刀体
type GroovingCuttingInternalBody struct {
	model.CutterBase
	TSYC       null.String // 刀具形式代码
	BMC        null.String // 刀体材料代码
	CCMS       null.String // 机床侧接口代码
	CCTMS      null.String // 机床侧接口类型代码
	CZCMS      null.String // 机床侧接口尺寸代码
	IIC        null.String // 刀片接口代码
	IISC       null.String // 刀片接口标准代码
	ISC        null.String // 刀片尺寸代码
	HAND       null.String // 切削方向
	LF         null.Int    // 功能长度
	CDX        null.Int    // 最大切削深度
	DAXX       null.Int    // 端面切槽外圈最大直径
	DAXN       null.Int    // 端面切槽外圈最小直径
	DMIN       null.Int    // 允许最小切削直径
	CSP        null.Int    // 冷却液供给特性
	DPC        null.Int    // 阻尼减振属性
	GRVIN      null.Int    // 内圆切槽
	GRVINFC    null.Int    // 端面切槽（内部）
	GRVINPR    null.Int    // 内圆轮廓
	GRVINTR    null.Int    // 内圆车削
	GRVINBORBK null.Int    // 内圆背镗
	GRVINUND   null.Int    // 内圆掏槽
}

func groovingCuttingInternalBodyRowInitial(body GroovingCuttingInternalBody) *model.GroovingCuttingInternalBody {
	row := &model.GroovingCuttingInternalBody{
		TSYC:       body.TSYC,       // 刀具形式代码
		BMC:        body.BMC,        // 刀体材料代码
		CCMS:       body.CCMS,       // 机床侧接口代码
		CCTMS:      body.CCTMS,      // 机床侧接口类型代码
		CZCMS:      body.CZCMS,      // 机床侧接口尺寸代码
		IIC:        body.IIC,        // 刀片接口代码
		IISC:       body.IISC,       // 刀片接口标准代码
		ISC:        body.ISC,        // 刀片尺寸代码
		HAND:       body.HAND,       // 切削方向
		LF:         body.LF,         // 功能长度
		CDX:        body.CDX,        // 最大切削深度
		DAXX:       body.DAXX,       // 端面切槽外圈最大直径
		DAXN:       body.DAXN,       // 端面切槽外圈最小直径
		DMIN:       body.DMIN,       // 允许最小切削直径
		CSP:        body.CSP,        // 冷却液供给特性
		DPC:        body.DPC,        // 阻尼减振属性
		GRVIN:      body.GRVIN,      // 内圆切槽
		GRVINFC:    body.GRVINFC,    // 端面切槽（内部）
		GRVINPR:    body.GRVINPR,    // 内圆轮廓
		GRVINTR:    body.GRVINTR,    // 内圆车削
		GRVINBORBK: body.GRVINBORBK, // 内圆背镗
		GRVINUND:   body.GRVINUND,   // 内圆掏槽
	}
	row.ID = body.ID
	row.OrderNumber = body.OrderNumber
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.Manufacturer = body.Manufacturer
	row.Submitter = body.Submitter
	return row
}

// GroovingCuttingInternalBodyQueryWithOrderNumber 查询
func GroovingCuttingInternalBodyQueryWithOrderNumber(orderNumber, userID string) ([]model.GroovingCuttingInternalBody, error) {
	rows := []model.GroovingCuttingInternalBody{}
	error := model.GroovingCuttingInternalBodyQueryWithOrderNumber(&rows, orderNumber, userID)
	return rows, error
}

// GroovingCuttingInternalBodyUpsert 更新或者新增
func GroovingCuttingInternalBodyUpsert(body GroovingCuttingInternalBody) (int, error) {
	row := groovingCuttingInternalBodyRowInitial(body)
	return row.GroovingCuttingInternalBodyUpsert()
}
