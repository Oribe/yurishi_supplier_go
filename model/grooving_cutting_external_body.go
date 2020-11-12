package model

import (
	"fmt"

	"github.com/guregu/null"
)

// GroovingCuttingExternalBody 外圆切槽切断刀体
type GroovingCuttingExternalBody struct {
	CutterBase
	TSYC      null.String // 刀具形式代码
	BMC       null.String // 刀体材料代码
	CCMS      null.String // 机床侧接口代码
	CCTMS     null.String // 机床侧接口类型代码
	CZCMS     null.String // 机床侧接口尺寸代码
	IIC       null.String // 刀片接口代码
	IISC      null.String // 刀片接口标准代码
	ISC       null.String // 刀垫尺寸代码
	HAND      null.String // 切削方向
	LF        null.Int    // 功能长度
	CUTDIA    null.Int    // 最大工件切断直径
	CDX       null.Int    // 最大切削深度
	DAXX      null.Int    // 端面切槽外圈最大直径
	DAXN      null.Int    // 端面切槽外圈最小直径
	CSP       null.Int    // 冷却液供给特性
	DPC       null.Int    // 阻尼减振属性
	PART      null.Int    // 切断
	GRVEX     null.Int    // 外圆切槽
	GRVEXFC   null.Int    // 端面切槽（外部）
	GRVEXPR   null.Int    // 外圆轮廓
	GRVEXTR   null.Int    // 外圆车削
	GRVEXTRBK null.Int    // 外圆背车
	GRVEXUND  null.Int    // 外圆掏槽
}

// InsertToDB 写入数据库
func (g *GroovingCuttingExternalBody) InsertToDB() (int64, error) {
	sql := `INSERT INTO
						grooving_cutting_external_bodies
						(category,subCategory,manufacturer,orderNumber,TSYC,BMC,CCMS,CCTMS,CZCMS,IIC,IISC,ISC,HAND,LF,CUTDIA,CDX,DAXX,DAXN,CSP,DPC,PART,GRVEX,GRVEXFC,GRVEXPR,GRVEXTR,GRVEXTRBK,GRVEXUND,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:TSYC,:BMC,:CCMS,:CCTMS,:CZCMS,:IIC,:IISC,:ISC,:HAND,:LF,:CUTDIA,:CDX,:DAXX,:DAXN,:CSP,:DPC,:PART,:GRVEX,:GRVEXFC,:GRVEXPR,:GRVEXTR,:GRVEXTRBK,:GRVEXUND,:submitter)`
	result, err := db.NamedExec(sql, g)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingExternalBody insert failed: %v", err.Error())
		return 0, error
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("GroovingCuttingExternalBody get last id failed: %v", err.Error())
		return lastID, error
	}
	return lastID, nil
}

// GroovingCuttingExternalBodyQueryWithOrderNumber 查询
func GroovingCuttingExternalBodyQueryWithOrderNumber(list *[]GroovingCuttingExternalBody, orderNumber, userID string) error {
	sql := `SELECT
						category,subCategory,manufacturer,orderNumber,TSYC,BMC,CCMS,CCTMS,CZCMS,IIC,IISC,ISC,HAND,LF,CUTDIA,CDX,DAXX,DAXN,CSP,DPC,PART,GRVEX,GRVEXFC,GRVEXPR,GRVEXTR,GRVEXTRBK,GRVEXUND,submitter
					FROM
						grooving_cutting_external_bodies
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						submitter = ?`
	stmt, err := db.Preparex(sql)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingExternalBody prepare sql failed: %v", err.Error())
		return error
	}
	err = stmt.Select(list, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingExternalBody query failed: %v", err.Error())
		return error
	}
	return nil
}
