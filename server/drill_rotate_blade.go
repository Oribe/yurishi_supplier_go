package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// DrillRotateBlade 可转位钻头刀片
type DrillRotateBlade struct {
	model.CutterBase
	BladeModel                  null.String `json:"bladeModel"`                  // 刀片型号
	MaterialLevel               null.String `json:"materialLevel"`               // 材质等级代码
	BladeInterfaceCode          null.String `json:"bladeInterfaceCode"`          // 刀片接口代码
	ManufacturerLogo            null.String `json:"manufacturerLogo"`            // 制造商标识
	Hand                        null.String `json:"hand"`                        // 切削方向
	BladeSizeCode               null.String `json:"bladeSizeCode"`               // 刀片尺寸代码
	BladeUsageCode              null.String `json:"bladeUsageCode"`              // 刀片用途代码
	BladeEdgeRadius             null.String `json:"bladeEdgeRadius"`             // 刀尖圆角半径
	ChipbreakerManufacturerLogo null.String `json:"chipbreakerManufacturerLogo"` // 断屑槽制造商标识
	CuttingEdgeCount            null.Int    `json:"cuttingEdgeCount"`            // 刀片可转位次数
	WiperBladeLogo              null.Int    `json:"wiperBladeLogo"`              // 修光刃标识
	P                           null.Int    `json:"P"`                           // 非合金钢、合金钢、铁素体和马氏体不锈钢
	M                           null.Int    `json:"M"`                           // 奥氏体不锈钢、双相不锈钢
	K                           null.Int    `json:"K"`                           // 铸块
	N                           null.Int    `json:"N"`                           // 有色金属
	S                           null.Int    `json:"S"`                           // 钛合金和高温合金
	H                           null.Int    `json:"H"`                           // 硬材料
	O                           null.Int    `json:"O"`                           // 非金属和符合材料
}

func drillRotateBladeRowInitial(body DrillRotateBlade) model.DrillRotateBlade {
	row := model.DrillRotateBlade{
		BladeModel:                  body.BladeModel,                  // 刀片型号
		MaterialLevel:               body.MaterialLevel,               // 材质等级代码
		BladeInterfaceCode:          body.BladeInterfaceCode,          // 刀片接口代码
		ManufacturerLogo:            body.ManufacturerLogo,            // 制造商标识
		Hand:                        body.Hand,                        // 切削方向
		BladeSizeCode:               body.BladeSizeCode,               // 刀片尺寸代码
		BladeUsageCode:              body.BladeUsageCode,              // 刀片用途代码
		BladeEdgeRadius:             body.BladeEdgeRadius,             // 刀尖圆角半径
		ChipbreakerManufacturerLogo: body.ChipbreakerManufacturerLogo, // 断屑槽制造商标识
		CuttingEdgeCount:            body.CuttingEdgeCount,            // 刀片可转位次数
		WiperBladeLogo:              body.WiperBladeLogo,              // 修光刃标识
		P:                           body.P,                           // 非合金钢、合金钢、铁素体和马氏体不锈钢
		M:                           body.M,                           // 奥氏体不锈钢、双相不锈钢
		K:                           body.K,                           // 铸块
		N:                           body.N,                           // 有色金属
		S:                           body.S,                           // 钛合金和高温合金
		H:                           body.H,                           // 硬材料
		O:                           body.O,                           // 非金属和符合材料
	}

	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.Manufacturer = body.Manufacturer
	row.OrderNumber = body.OrderNumber
	row.Submitter = body.Submitter
	return row
}

// DrillRotateBladeInsert 添加
func DrillRotateBladeInsert(body DrillRotateBlade) (int64, error) {
	row := drillRotateBladeRowInitial(body)
	return row.InsertToDB()
}

// DrillRotateBladeQueryWithOrderNumber 查询
func DrillRotateBladeQueryWithOrderNumber(orderNumber, userID string) ([]model.DrillRotateBlade, error) {
	rows := []model.DrillRotateBlade{}
	err := model.DrillRotateBladeQueryWithOrderNumber(&rows, orderNumber, userID)
	return rows, err
}
