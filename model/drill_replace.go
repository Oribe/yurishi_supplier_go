package model

import (
	"fmt"

	"github.com/guregu/null"
)

// DrillReplace 可更换钻削刀头
type DrillReplace struct {
	CutterBase
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
	O                                null.Int    `db:"O"`                                                                        // 非金属和符合材料
}

// InsertToDB 添加
func (d *DrillReplace) InsertToDB() (int64, error) {
	sql := `INSERT INTO
						drill_replace_suppliers
						(category,subCategory,manufacturer,orderNumber,bladeModel,materialLevel,bladeInterfaceCode,manufacturerLogo,hand,bladeInterfaceConnectionSizeCode,chipbreakerManufacturerLogo,cuttingDiameter,P,M,K,N,S,H,O,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:bladeModel,:materialLevel,:bladeInterfaceCode,:manufacturerLogo,:hand,:bladeInterfaceConnectionSizeCode,:chipbreakerManufacturerLogo,:cuttingDiameter,:P,:M,:K,:N,:S,:H,:O,:submitter)`
	result, err := db.NamedExec(sql, d)
	if err != nil {
		error := fmt.Errorf("drill_replace insert failed: %v", err.Error())
		return 0, error
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("drill_replace get last id failed: %v", err.Error())
		return 0, error
	}
	return lastID, nil
}

// DrillReplaceQueryWithOrderNumber 查询
func DrillReplaceQueryWithOrderNumber(list *[]DrillReplace, orderNumber, userID string) error {
	sql := `SELECT 
						category,subCategory,manufacturer,orderNumber,bladeModel,materialLevel,bladeInterfaceCode,manufacturerLogo,hand,bladeInterfaceConnectionSizeCode,chipbreakerManufacturerLogo,cuttingDiameter,P,M,K,N,S,H,O,submitter
					FROM
						drill_replace_suppliers
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						userId = ?`
	stmt, err := db.Preparex(sql)
	if err != nil {
		error := fmt.Errorf("drill_replace select sql prepare failed: %v", err.Error())
		return error
	}
	err = stmt.Select(list, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("drill_replace selected failed: %v", err.Error())
		return error
	}
	return nil
}
