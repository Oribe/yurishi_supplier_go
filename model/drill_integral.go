package model

import (
	"fmt"

	"github.com/guregu/null"
)

// DrillIntegral 整体式钻头
type DrillIntegral struct {
	CutterBase
	TPCODE null.String `db:"TPCODE"`             // 刀具型号
	GRADE  null.String `db:"GRADE"`              // 材质等级代码
	SDTP   null.String `db:"SDTP"`               // 钻头类型
	ZEEP   null.Int    `db:"ZEEP"`               // 切削刃数
	CZCMS  null.String `db:"CZCMS"`              // 机床侧接口代码
	TPCON  null.String `db:"TPCON"`              // 接口类型
	DCON   null.Int    `db:"DCON"`               // 连接直径
	HAND   null.String `db:"HAND"`               // 切削方向
	DC     null.Int    `db:"DC"`                 // 切削直径
	OAL    null.Int    `db:"OAL"`                // 刀具总长
	LU     null.Int    `db:"LU"`                 // 可用长度
	CSP    null.Int    `db:"CSP"`                // 冷却液供给特征
	STN    null.Int    `db:"STN"`                // 阶梯数
	SDL1   null.Int    `db:"SDL_1" json:"SDL_1"` // 直径长度-1
	STA2   null.Int    `db:"STA_2" json:"STA_2"` // 阶梯倒角-2
	DC2    null.Int    `db:"DC_2" json:"DC_2"`   // 阶梯直径-2
	SD2    null.Int    `db:"SD_2" json:"SD_2"`   // 阶梯长度-2
	SDL2   null.Int    `db:"SDL_2" json:"SDL_2"` // 直径长度-2
	STA3   null.Int    `db:"STA_3" json:"STA_3"` // 阶梯倒角-3
	DC3    null.Int    `db:"DC_3" json:"DC_3"`   // 阶梯直径-3
	SD3    null.Int    `db:"SD_3" json:"SD_3"`   // 阶梯长度-3
	P      null.Int    `db:"P"`                  // P
	M      null.Int    `db:"M"`                  // M
	K      null.Int    `db:"K"`                  // K
	N      null.Int    `db:"N"`                  // N
	S      null.Int    `db:"S"`                  // S
	H      null.Int    `db:"H"`                  // H
	O      null.Int    `db:"O"`                  // O
}

// InsertToDB 插入到数据库
func (d *DrillIntegral) InsertToDB() (int64, error) {
	sql := `INSERT INTO
						drill_integrals
						(category,subCategory,manufacturer,orderNumber,TPCODE,GRADE,SDTP,ZEEP,CZCMS,TPCON,DCON,HAND,DC,OAL,LU,CSP,STN,SDL_1,STA_2,DC_2,SD_2,SDL_2,STA_3,DC_3,SD_3,P,M,K,N,S,H,O,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:TPCODE,:GRADE,:SDTP,:ZEEP,:CZCMS,:TPCON,:DCON,:HAND,:DC,:OAL,:LU,:CSP,:STN,:SDL_1,:STA_2,:DC_2,:SD_2,:SDL_2,:STA_3,:DC_3,:SD_3,:P,:M,:K,:N,:S,:H,:O,:submitter)`
	result, err := db.NamedExec(sql, d)
	if err != nil {
		error := fmt.Errorf("drill_integral insert failed: %v", err.Error())
		return 0, error
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("drill_integral get last id failed: %v", err.Error())
		return 0, error
	}

	return lastInsertID, nil
}

// DrillIntegralQueryWithOrderNumber 更具订单查询
func DrillIntegralQueryWithOrderNumber(data *[]DrillIntegral, orderNumber, userID string) error {
	sql := `SELECT 
						category,subCategory,manufacturer,orderNumber,TPCODE,GRADE,SDTP,ZEEP,CZCMS,TPCON,DCON,HAND,DC,OAL,LU,CSP,STN,SDL_1,STA_2,DC_2,SD_2,SDL_2,STA_3,DC_3,SD_3,P,M,K,N,S,H,O,submitter
					FROM
						drill_integrals
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						submitter = ?`
	stmt, err := db.Preparex(sql)
	if err != nil {
		error := fmt.Errorf("drill_integral prepare sql failed: %v", err.Error())
		return error
	}
	err = stmt.Select(data, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("drill_integral query failed: %v", err.Error())
		return error
	}
	return nil
}
