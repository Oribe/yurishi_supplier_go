package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// DrillReplace 可更换钻削刀头
type DrillReplace struct {
	model.CutterBase
	BladeModel                       null.String `db:"bladeModel" json:"bladeModel"`                                             // 刀片型号
	MaterialLevel                    null.String `db:"materialLevel" json:"materialLevel"`                                       // 材质等级代码
	BladeInterfaceCode               null.String `db:"bladeInterfaceCode" json:"bladeInterfaceCode"`                             // 刀片接口代码
	ManufacturerLogo                 null.String `db:"manufacturerLogo" json:"manufacturerLogo"`                                 // 制造商标识
	Hand                             null.String `db:"hand" json:"hand"`                                                         // 切削方向
	BladeInterfaceConnectionSizeCode null.String `db:"bladeInterfaceConnectionSizeCode" json:"bladeInterfaceConnectionSizeCode"` // 连接尺寸代码
	ChipbreakerManufacturerLogo      null.String `db:"chipbreakerManufacturerLogo" json:"chipbreakerManufacturerLogo"`           // 断屑槽制造商标识
	CuttingDiameter                  null.String `db:"cuttingDiameter" json:"cuttingDiameter"`                                   // 切削直径
	P                                null.Int    `db:"P"`                                                                        // 非合金钢、合金钢、铁素体和马氏体不锈钢
	M                                null.Int    `db:"M"`                                                                        // 奥氏体不锈钢、双相不锈钢
	K                                null.Int    `db:"K"`                                                                        // 铸块
	N                                null.Int    `db:"N"`                                                                        // 有色金属
	S                                null.Int    `db:"S"`                                                                        // 钛合金和高温合金
	H                                null.Int    `db:"H"`                                                                        // 硬材料
	O                                null.Int    `db:"O"`
}

func drillReplaceRowInitial(body DrillReplace) model.DrillReplace {
	row := model.DrillReplace{
		BladeModel:                       body.BladeModel,
		MaterialLevel:                    body.MaterialLevel,
		BladeInterfaceCode:               body.BladeInterfaceCode,
		ManufacturerLogo:                 body.ManufacturerLogo,
		Hand:                             body.Hand,
		BladeInterfaceConnectionSizeCode: body.BladeInterfaceConnectionSizeCode,
		ChipbreakerManufacturerLogo:      body.ChipbreakerManufacturerLogo,
		CuttingDiameter:                  body.CuttingDiameter,
		P:                                body.P,
		M:                                body.M,
		K:                                body.K,
		N:                                body.N,
		S:                                body.S,
		H:                                body.H,
		O:                                body.O,
	}
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.Manufacturer = body.Manufacturer
	row.OrderNumber = body.OrderNumber
	row.Submitter = body.Submitter
	return row
}

// DrillReplaceInsertToDB 写入数据库
func DrillReplaceInsertToDB(body DrillReplace) (int64, error) {
	row := drillReplaceRowInitial(body)
	id, err := row.InsertToDB()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// DrillReplaceQueryWithOrderNumber 查询
func DrillReplaceQueryWithOrderNumber(orderNumber, userID string) ([]model.DrillReplace, error) {
	list := []model.DrillReplace{}
	err := model.DrillReplaceQueryWithOrderNumber(&list, orderNumber, userID)
	if err != nil {
		return nil, err
	}
	return list, nil
}
