package model

import (
	"fmt"

	"github.com/guregu/null"
)

// DrillReplaceBody 可更换钻削刀头刀体
type DrillReplaceBody struct {
	CutterBase
	BMC   null.String // 刀体材料代码
	CCMS  null.String // 机床侧接口代码
	CCTMS null.String // 机床侧接口类型代码
	CZCMS null.String // 机床侧接口尺寸代码
	IIC   null.String // 刀片接口代码
	IISC  null.String // 刀片接口标准代码
	ISC   null.String // 刀片尺寸代码
	LF    null.Int    // 功能长度
	LU    null.Int    // 可用长度
	CSP   null.Int    // 冷却液供给特性
	DPC   null.Int    // 阻尼减振属性
}

// InsertToDB 写入数据库
func (d *DrillReplaceBody) InsertToDB() (int64, error) {
	sql := `INSERT INTO
						drill_replace_bodies
						(category,subCategory,manufacturer,orderNumber,BMC,CCMS,CCTMS,CZCMS,IIC,IISC,ISC,LF,LU,CSP,DPC,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:BMC,:CCMS,:CCTMS,:CZCMS,:IIC,:IISC,:ISC,:LF,:LU,:CSP,:DPC,:submitter)`
	result, err := db.NamedExec(sql, d)
	if err != nil {
		error := fmt.Errorf("drill_replace_bodies insert failed: %v", err.Error())
		return 0, error
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("drill_replace_bodies get last id failed: %v", err.Error())
		return lastID, error
	}
	return lastID, nil
}

// DrillReplaceBodyQueryWithOrderNumber 查询
func DrillReplaceBodyQueryWithOrderNumber(list *[]DrillReplaceBody, orderNumber, userID string) error {
	sql := `SELECT 
						category,subCategory,manufacturer,orderNumber,BMC,CCMS,CCTMS,CZCMS,IIC,IISC,ISC,LF,LU,CSP,DPC,submitter
					FROM
						drill_replace_bodies
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						userID = ?`
	stmt, err := db.Preparex(sql)
	if err != nil {
		error := fmt.Errorf("drill_replace_bodies prepare sql failed: %v", err.Error())
		return error
	}
	err = stmt.Select(list, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("drill_replace_bodies select failed: %v", err.Error())
		return error
	}
	return nil
}
