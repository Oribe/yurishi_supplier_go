package model

import (
	"fmt"

	"github.com/guregu/null"
)

// DrillRotateBlade 可转位钻头刀片
type DrillRotateBlade struct {
	CutterBase
	BladeModel                  null.String `db:"bladeModel"`                  // 刀片型号
	MaterialLevel               null.String `db:"materialLevel"`               // 材质等级代码
	BladeInterfaceCode          null.String `db:"bladeInterfaceCode"`          // 刀片接口代码
	ManufacturerLogo            null.String `db:"manufacturerLogo"`            // 制造商标识
	Hand                        null.String `db:"hand"`                        // 切削方向
	BladeSizeCode               null.String `db:"bladeSizeCode"`               // 刀片尺寸代码
	BladeUsageCode              null.String `db:"bladeUsageCode"`              // 刀片用途代码
	BladeEdgeRadius             null.String `db:"bladeEdgeRadius"`             // 刀尖圆角半径
	ChipbreakerManufacturerLogo null.String `db:"chipbreakerManufacturerLogo"` // 断屑槽制造商标识
	CuttingEdgeCount            null.Int    `db:"cuttingEdgeCount"`            // 刀片可转位次数
	WiperBladeLogo              null.Int    `db:"wiperBladeLogo"`              // 修光刃标识
	P                           null.Int    `db:"P"`                           // 非合金钢、合金钢、铁素体和马氏体不锈钢
	M                           null.Int    `db:"M"`                           // 奥氏体不锈钢、双相不锈钢
	K                           null.Int    `db:"K"`                           // 铸块
	N                           null.Int    `db:"N"`                           // 有色金属
	S                           null.Int    `db:"S"`                           // 钛合金和高温合金
	H                           null.Int    `db:"H"`                           // 硬材料
	O                           null.Int    `db:"O"`                           // 非金属和符合材料
}

// InsertToDB 写入数据库
func (d *DrillRotateBlade) InsertToDB() (int64, error) {
	sql := `INSERT INTO
						drill_rotate_blade_suppliers
						(category,subCategory,manufacturer,orderNumber,bladeModel,materialLevel,bladeInterfaceCode,manufacturerLogo,hand,bladeSizeCode,bladeUsageCode,bladeEdgeRadius,chipbreakerManufacturerLogo,cuttingEdgeCount,wiperBladeLogo,P,M,K,N,S,H,O,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:bladeModel,:materialLevel,:bladeInterfaceCode,:manufacturerLogo,:hand,:bladeSizeCode,:bladeUsageCode,:bladeEdgeRadius,:chipbreakerManufacturerLogo,:cuttingEdgeCount,:wiperBladeLogo,:P,:M,:K,:N,:S,:H,:O,:submitter)`
	result, err := db.NamedExec(sql, d)
	if err != nil {
		error := fmt.Errorf("drill_rotate_blade insert failed: %v", err.Error())
		return 0, error
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("drill_rotate_blade get last id failed: %v", err.Error())
		return lastID, error
	}
	return lastID, nil
}

// DrillRotateBladeQueryWithOrderNumber 查询
func DrillRotateBladeQueryWithOrderNumber(rows *[]DrillRotateBlade, orderNumber, userID string) error {
	sql := `SELECT
						category,subCategory,manufacturer,orderNumber,bladeModel,materialLevel,bladeInterfaceCode,manufacturerLogo,hand,bladeSizeCode,bladeUsageCode,bladeEdgeRadius,chipbreakerManufacturerLogo,cuttingEdgeCount,wiperBladeLogo,P,M,K,N,S,H,O,submitter
					FROM
						drill_rotate_blade_suppliers
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						userId = ?`
	stmt, err := db.Preparex(sql)
	if err != nil {
		error := fmt.Errorf("drill_rotate_blade prepare sql failed: %v", err.Error())
		return error
	}
	err = stmt.Select(rows, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("drill_rotate_blade query failed: %v", err.Error())
		return error
	}
	return nil
}
