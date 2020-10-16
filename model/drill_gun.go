package model

import (
	"fmt"
)

// DrillGun 枪钻
type DrillGun struct {
	CutterBase
	BMC       string `db:"BMC"`       // 刀体材料代码
	SUBSTRATE string `db:"SUBSTRATE"` // 刀具基体材料
	CCMS      string `db:"CCMS"`      // 机床侧接口代码
	CCTMS     string `db:"CCTMS"`     // 机床侧接口类型代码
	CZCMS     string `db:"CZCMS"`     // 机床侧接口尺寸代码
	HAND      string `db:"HAND"`      // 切削方向
	DC        int    `db:"DC"`        // 切削直径
	LF        int    `db:"LF"`        // 功能长度
	LU        int    `db:"LU"`        // 可用长度
	CSP       int    `db:"CSP"`       // 冷却液供给特性
	DPC       int    `db:"DPC"`       // 阻尼减振属性
	AMCP      int    `db:"AMCP"`      // 适用加工材料大类P
	AMCM      int    `db:"AMCM"`      // 适用加工材料大类M
	AMCK      int    `db:"AMCK"`      // 适用加工材料大类K
	AMCN      int    `db:"AMCN"`      // 适用加工材料大类N
	AMCS      int    `db:"AMCS"`      // 适用加工材料大类S
	AMCH      int    `db:"AMCH"`      // 适用加工材料大类H
	AMCO      int    `db:"AMCO"`      // 适用加工材料大类O
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
