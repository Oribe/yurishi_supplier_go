package model

import (
	"fmt"
)

// DrillGun 枪钻
type DrillGun struct {
	CutterBase
	BMC       string // 刀体材料代码
	SUBSTRATE string // 刀具基体材料
	CCMS      string // 机床侧接口代码
	CCTMS     string // 机床侧接口类型代码
	CZCMS     string // 机床侧接口尺寸代码
	HAND      string // 切削方向
	DC        string // 切削直径
	LF        string // 功能长度
	LU        string // 可用长度
	CSP       string // 冷却液供给特性
	DPC       string // 阻尼减振属性
	AMCP      string // 适用加工材料大类P
	AMCM      string // 适用加工材料大类M
	AMCK      string // 适用加工材料大类K
	AMCN      string // 适用加工材料大类N
	AMCS      string // 适用加工材料大类S
	AMCH      string // 适用加工材料大类H
	AMCO      string // 适用加工材料大类O
}

// DrillGunInsertToDB 新增
func (d *DrillGun) DrillGunInsertToDB() (int64, error) {
	sql := `INSERT INTO 
					drill_guns
					(category,subCategory,manufacturer,orderNumber,BMC,SUBSTRATE,CCMS,CCTMS,CZCMS,HAND,DC,LF,LU,CSP,DPC,AMCP,AMCM,AMCK,AMCN,AMCS,AMCH,AMCO)
				 VALUES
					(:category,:subCategory,:manufacturer,:orderNumber,:BMC,:SUBSTRATE,:CCMS,:CCTMS,:CZCMS,:HAND,:DC,:LF,:LU,:CSP,:DPC,:AMCP,:AMCM,:AMCK,:AMCN,:AMCS,:AMCH,:AMCO)`
	result, err := db.NamedExec(sql, d)
	if err != nil {
		error := fmt.Errorf("drill_gun query failed: %v", err.Error())
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
func DrillGunQueryWithOrderNumber(data *[]DrillGun, orderNumber string, userID int) error {
	sql := `SELECT
						category,subCategory,manufacturer,orderNumber,BMC,SUBSTRATE,CCMS,CCTMS,CZCMS,HAND,DC,LF,LU,CSP,DPC,AMCP,AMCM,AMCK,AMCN,AMCS,AMCH,AMCO
					FROM
						drill_guns
					WHERE
						orderNumber like '%?%'
					AND
						submitter = ?`
	stmt, err := db.Prepare(sql)
	if err != nil {
		error := fmt.Errorf("drill_gun sql error: %v", err.Error())
		return error
	}
	rows, err := stmt.Query(orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("drill_gun query failed: %v", err.Error())
		return error
	}
	for rows.Next() {
		err := rows.Scan(data)
		if err != nil {
			error := fmt.Errorf("drill_gun scan failed: %v", err.Error())
			return error
		}
	}
	return nil
}
