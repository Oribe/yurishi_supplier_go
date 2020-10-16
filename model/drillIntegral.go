package model

import "fmt"

// DrillIntegral 整体式钻头
type DrillIntegral struct {
	CutterBase
	TPCODE string `db:"TPCODE"` // 刀具型号
	GRADE  string `db:"GRADE"`  // 材质等级代码
	SDTP   string `db:"SDTP"`   // 钻头类型
	ZEEP   int    `db:"ZEEP"`   // 切削刃数
	CZCMS  string `db:"CZCMS"`  // 机床侧接口代码
	TPCON  string `db:"TPCON"`  // 接口类型
	DCON   int    `db:"DCON"`   // 连接直径
	HAND   string `db:"HAND"`   // 切削方向
	DC     int    `db:"DC"`     // 切削直径
	OAL    int    `db:"OAL"`    // 刀具总长
	LU     int    `db:"LU"`     // 可用长度
	CSP    int    `db:"CSP"`    // 冷却液供给特征
	STN    int    `db:"STN"`    // 阶梯数
	SDL1   int    `db:"SDL_1"`  // 直径长度-1
	STA2   int    `db:"STA_2"`  // 阶梯倒角-2
	DC2    int    `db:"DC_2"`   // 阶梯直径-2
	SD2    int    `db:"SD_2"`   // 阶梯长度-2
	SDL2   int    `db:"SDL_2"`  // 直径长度-2
	STA3   int    `db:"STA_3"`  // 阶梯倒角-3
	DC3    int    `db:"DC_3"`   // 阶梯直径-3
	SD3    int    `db:"SD_3"`   // 阶梯长度-3
	P      int    `db:"P"`      // P
	M      int    `db:"M"`      // M
	K      int    `db:"K"`      // K
	N      int    `db:"N"`      // N
	S      int    `db:"S"`      // S
	H      int    `db:"H"`      // H
	O      int    `db:"O"`      // O
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
