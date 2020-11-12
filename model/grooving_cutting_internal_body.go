package model

import (
	"fmt"
	"manufacture_supplier_go/util"

	"github.com/guregu/null"
)

// GroovingCuttingInternalBody 內圆切槽切断刀体
type GroovingCuttingInternalBody struct {
	CutterBase
	TSYC       null.String `db:"TSYC"`       // 刀具形式代码
	BMC        null.String `db:"BMC"`        // 刀体材料代码
	CCMS       null.String `db:"CCMS"`       // 机床侧接口代码
	CCTMS      null.String `db:"CCTMS"`      // 机床侧接口类型代码
	CZCMS      null.String `db:"CZCMS"`      // 机床侧接口尺寸代码
	IIC        null.String `db:"IIC"`        // 刀片接口代码
	IISC       null.String `db:"IISC"`       // 刀片接口标准代码
	ISC        null.String `db:"ISC"`        // 刀片尺寸代码
	HAND       null.String `db:"HAND"`       // 切削方向
	LF         null.Int    `db:"LF"`         // 功能长度
	CDX        null.Int    `db:"CDX"`        // 最大切削深度
	DAXX       null.Int    `db:"DAXX"`       // 端面切槽外圈最大直径
	DAXN       null.Int    `db:"DAXN"`       // 端面切槽外圈最小直径
	DMIN       null.Int    `db:"DMIN"`       // 允许最小切削直径
	CSP        null.Int    `db:"CSP"`        // 冷却液供给特性
	DPC        null.Int    `db:"DPC"`        // 阻尼减振属性
	GRVIN      null.Int    `db:"GRVIN"`      // 内圆切槽
	GRVINFC    null.Int    `db:"GRVINFC"`    // 端面切槽（内部）
	GRVINPR    null.Int    `db:"GRVINPR"`    // 内圆轮廓
	GRVINTR    null.Int    `db:"GRVINTR"`    // 内圆车削
	GRVINBORBK null.Int    `db:"GRVINBORBK"` // 内圆背镗
	GRVINUND   null.Int    `db:"GRVINUND"`   // 内圆掏槽
}

// InsertToDB 插入数据库
func (g *GroovingCuttingInternalBody) insertToDB() error {
	sql := `INSERT INTO 
						grooving_cutting_internal_bodies
						(category,subCategory,manufacturer,orderNumber,TSYC,BMC,CCMS,CCTMS,CZCMS,IIC,IISC,HAND,LF,CDX,DAXX,DAXN,DMIN,CSP,DPC,GRVIN,GRVINFC,GRVINPR,GRVINTR,GRVINBORBK,GRVINUND,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:TSYC,:BMC,:CCMS,:CCTMS,:CZCMS,:IIC,:IISC,:HAND,:LF,:CDX,:DAXX,:DAXN,:DMIN,:CSP,:DPC,:GRVIN,:GRVINFC,:GRVINPR,:GRVINTR,:GRVINBORBK,:GRVINUND,:submitter)
	`
	_, err := db.NamedExec(sql, g)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingInternalBody insert failed: %v", err.Error())
		return error
	}
	return nil
}

// UpdateToDB 更新到数据库
func (g *GroovingCuttingInternalBody) updateToDB() error {
	sql := `UPDATE 
						grooving_cutting_internal_bodies
					SET
						category=:category, subCategory=:subCategory, manufacturer=:manufacturer, orderNumber=:orderNumber, TSYC=:TSYC,
						BMC=:BMC, CCMS=:CCMS, CCTMS=:CCTMS,CZCMS=:CZCMS,IIC=:IIC,IISC=:IISC,HAND=:HAND,LF=:LF,CDX=:CDX,DAXX=:DAXX,DAXN=:DAXN,DMIN=:DMIN,
						CSP=:CSP,DPC=:DPC,GRVIN=:GRVIN,GRVINFC=:GRVINFC,GRVINPR=:GRVINPR,GRVINTR=:GRVINTR,GRVINBORBK=:GRVINBORBK,GRVINUND=:GRVINUND,submitter=:submitter
					WHERE
						id = :id
					`
	_, err := db.NamedExec(sql, g)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingInternalBody update failed: %v", err.Error())
		return error
	}
	return nil
}

func groovingCuttingInternalBodyQueryByID(id int64) (bool, error) {
	if id <= 0 {
		return false, nil
	}
	row := GroovingCuttingInternalBody{}
	sql := `SELECT * FROM grooving_cutting_internal_bodies WHERE id = ? LIMIT 1`
	err := db.Get(&row, sql, id)
	if err != nil {
		return false, err
	}
	if row.ID != id {
		return false, nil
	}
	return true, nil
}

// GroovingCuttingInternalBodyUpsert 更新或者新增
func (g *GroovingCuttingInternalBody) GroovingCuttingInternalBodyUpsert() (int, error) {
	b, err := groovingCuttingInternalBodyQueryByID(g.ID)
	if err != nil {
		return -1, err
	}
	if b { // 更新
		return util.UPDATE, g.updateToDB()
	}
	// 新增
	return util.INSERT, g.insertToDB()
}

// GroovingCuttingInternalBodyQueryWithOrderNumber 搜索
func GroovingCuttingInternalBodyQueryWithOrderNumber(rows *[]GroovingCuttingInternalBody, orderNumber, userID string) error {
	sql := `SELECT 
						id,category,subCategory,manufacturer,orderNumber,TSYC,BMC,CCMS,CCTMS,CZCMS,IIC,IISC,HAND,LF,CDX,DAXX,DAXN,DMIN,CSP,DPC,GRVIN,GRVINFC,GRVINPR,GRVINTR,GRVINBORBK,GRVINUND,submitter
					FROM
						grooving_cutting_internal_bodies
					WHERE
						orderNumber LIKE CONCAT('%',?,'%')
					AND
						submitter = ?
					`
	err := db.Select(rows, sql, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingInternalBody query failed: %v", err.Error())
		return error
	}
	return nil
}
