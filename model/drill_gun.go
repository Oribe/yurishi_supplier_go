package model

import (
	"fmt"

	"github.com/guregu/null"
)

// DrillGun 枪钻
type DrillGun struct {
	CutterBase
	BMC       null.String `db:"BMC" json:"BMC"`             // 刀体材料代码
	SUBSTRATE null.String `db:"SUBSTRATE" json:"SUBSTRATE"` // 刀具基体材料
	CCMS      null.String `db:"CCMS" json:"CCMS"`           // 机床侧接口代码
	CCTMS     null.String `db:"CCTMS" json:"CCTMS"`         // 机床侧接口类型代码
	CZCMS     null.String `db:"CZCMS" json:"CZCMS"`         // 机床侧接口尺寸代码
	HAND      null.String `db:"HAND" json:"HAND"`           // 切削方向
	DC        null.Int    `db:"DC" json:"DC"`               // 切削直径
	LF        null.Int    `db:"LF" json:"LF"`               // 功能长度
	LU        null.Int    `db:"LU" json:"LU"`               // 可用长度
	CSP       null.Int    `db:"CSP" json:"CSP"`             // 冷却液供给特性
	DPC       null.Int    `db:"DPC" json:"DPC"`             // 阻尼减振属性
	AMCP      null.Int    `db:"AMCP" json:"AMCP"`           // 适用加工材料大类P
	AMCM      null.Int    `db:"AMCM" json:"AMCM"`           // 适用加工材料大类M
	AMCK      null.Int    `db:"AMCK" json:"AMCK"`           // 适用加工材料大类K
	AMCN      null.Int    `db:"AMCN" json:"AMCN"`           // 适用加工材料大类N
	AMCS      null.Int    `db:"AMCS" json:"AMCS"`           // 适用加工材料大类S
	AMCH      null.Int    `db:"AMCH" json:"AMCH"`           // 适用加工材料大类H
	AMCO      null.Int    `db:"AMCO" json:"AMCO"`           // 适用加工材料大类O
}

// DrillGunInsertToDB 新增
func (d *DrillGun) DrillGunInsertToDB() (int64, error) {
	sql := `INSERT INTO 
					drill_guns
					(category,subCategory,manufacturer,orderNumber,BMC,SUBSTRATE,CCMS,CCTMS,CZCMS,HAND,DC,LF,LU,CSP,DPC,AMCP,AMCM,AMCK,AMCN,AMCS,AMCH,AMCO,submitter)
				 VALUES
					(:category,:subCategory,:manufacturer,:orderNumber,:BMC,:SUBSTRATE,:CCMS,:CCTMS,:CZCMS,:HAND,:DC,:LF,:LU,:CSP,:DPC,:AMCP,:AMCM,:AMCK,:AMCN,:AMCS,:AMCH,:AMCO,:submitter)`
	result, err := db.NamedExec(sql, d)
	if err != nil {
		error := fmt.Errorf("drill_gun insert failed: %v", err.Error())
		return 0, error
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("drill_gun get lastId failed: %v", err.Error())
		return lastInsertID, error
	}
	return lastInsertID, nil
}

// DrillGunQueryWithOrderNumber 查询
func DrillGunQueryWithOrderNumber(data *[]DrillGun, orderNumber string, userID string) error {
	sql := `SELECT
						category,subCategory,manufacturer,orderNumber,BMC,SUBSTRATE,CCMS,CCTMS,CZCMS,HAND,DC,LF,LU,CSP,DPC,AMCP,AMCM,AMCK,AMCN,AMCS,AMCH,AMCO,submitter
					FROM
						drill_guns
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						submitter = ?`
	stmt, err := db.Preparex(sql)
	if err != nil {
		error := fmt.Errorf("drill_gun sql error: %v", err.Error())
		return error
	}
	err = stmt.Select(data, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("drill_gun query failed: %v", err.Error())
		return error
	}
	return nil
}
