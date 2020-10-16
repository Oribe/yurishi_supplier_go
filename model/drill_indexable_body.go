package model

import (
	"fmt"
)

// DrillIndexableBody 可转位钻头刀体
type DrillIndexableBody struct {
	CutterBase
	BMC   string `db:"BMC"`   // 刀体材料代码
	CCMS  string `db:"CCMS"`  // 机床侧接口代码
	CCTMS string `db:"CCTMS"` // 机床侧接口类型代码
	CZCMS string `db:"CZCMS"` // 机床侧接口尺寸代码
	IICP  string `db:"IICP"`  // 周边刀片接口代码
	IISCP string `db:"IISCP"` // 周边刀片接口标准代码
	ISCP  string `db:"ISCP"`  // 周边刀片尺寸代码
	IICC  string `db:"IICC"`  // 中心刀片接口代码
	IISCC string `db:"IISCC"` // 中心刀片接口标准代码
	ISCC  string `db:"ISCC"`  // 中心刀片尺寸代码
	HAND  string `db:"HAND"`  // 切削方向
	DC    int    `db:"DC"`    // 切削直径
	LF    int    `db:"LF"`    // 功能长度
	LU    int    `db:"LU"`    // 可用长度
	CSP   int    `db:"CSP"`   // 冷却液供给特性
	DPC   int    `db:"DPC"`   // 阻尼减振属性
}

// InsertToDB 插入到数据库
func (d *DrillIndexableBody) InsertToDB() (int64, error) {
	sql := `INSERT INTO 
						drill_indexable_bodies
						(category,subCategory,manufacturer,orderNumber,BMC,CCMS,CCTMS,CZCMS,IICP,IISCP,ISCP,IICC,IISCC,ISCC,HAND,DC,LF,LU,CSP,DPC,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:BMC,:CCMS,:CCTMS,:CZCMS,:IICP,:IISCP,:ISCP,:IICC,:IISCC,:ISCC,:HAND,:DC,:LF,:LU,:CSP,:DPC,:submitter)`
	result, err := db.NamedExec(sql, d)
	if err != nil {
		error := fmt.Errorf("drill_indexable_body insert failed: %v", err.Error())
		return 0, error
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("drill_indexable_body get last id failed: %v", err.Error())
		return 0, error
	}
	return lastInsertID, nil
}

// DrillIndexableBodyQueryWithOrderNumber 查询
func DrillIndexableBodyQueryWithOrderNumber(data *[]DrillIndexableBody, orderNumber, userID string) error {
	sql := `SELECT 
						category,subCategory,manufacturer,orderNumber,BMC,CCMS,CCTMS,CZCMS,IICP,IISCP,ISCP,IICC,IISCC,ISCC,HAND,DC,LF,LU,CSP,DPC,submitter
					FROM
						drill_indexable_bodies
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						submitter = ?`
	stmt, err := db.Preparex(sql)
	if err != nil {
		error := fmt.Errorf("drill_indexable_body prepare sql failed: %v", err.Error())
		return error
	}

	err = stmt.Select(data, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("drill_indexable_body query failed: %v", err.Error())
		return error
	}

	return nil
}
