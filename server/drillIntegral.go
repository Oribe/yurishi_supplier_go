package server

import "manufacture_supplier_go/model"

// DrillIntegral
type DrillIntegral struct {
	model.CutterBase
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

// DrillIntegralQueryWithOrderNumber 查询
func DrillIntegralQueryWithOrderNumber() {

}
