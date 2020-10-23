package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// DrillGunQueryWithOrder 查询
func DrillGunQueryWithOrder(orderNumber string, userID string) ([]model.DrillGun, error) {

	dataList := []model.DrillGun{}
	err := model.DrillGunQueryWithOrderNumber(&dataList, orderNumber, userID)
	if err != nil {
		return nil, err
	}
	return dataList, nil
}

// DrillGun 枪钻
type DrillGun struct {
	model.CutterBase
	BMC       null.String //`json:"BMC"`       // 刀体材料代码
	SUBSTRATE null.String //`json:"SUBSTRATE"` // 刀具基体材料
	CCMS      null.String //`json:"CCMS"`      // 机床侧接口代码
	CCTMS     null.String //`json:"CCTMS"`     // 机床侧接口类型代码
	CZCMS     null.String //`json:"CZCMS"`     // 机床侧接口尺寸代码
	HAND      null.String //`json:"HAND"`      // 切削方向
	DC        null.Int    //`json:"DC"`        // 切削直径
	LF        null.Int    //`json:"LF"`        // 功能长度
	LU        null.Int    //`json:"LU"`        // 可用长度
	CSP       null.Int    //`json:"CSP"`       // 冷却液供给特性
	DPC       null.Int    //`json:"DPC"`       // 阻尼减振属性
	AMCP      null.Int    //`json:"AMCP"`      // 适用加工材料大类P
	AMCM      null.Int    //`json:"AMCM"`      // 适用加工材料大类M
	AMCK      null.Int    //`json:"AMCK"`      // 适用加工材料大类K
	AMCN      null.Int    //`json:"AMCN"`      // 适用加工材料大类N
	AMCS      null.Int    //`json:"AMCS"`      // 适用加工材料大类S
	AMCH      null.Int    //`json:"AMCH"`      // 适用加工材料大类H
	AMCO      null.Int    //`json:"AMCO"`      // 适用加工材料大类O
}

func drillGunDataInit(data *DrillGun) *model.DrillGun {
	newData := &model.DrillGun{
		BMC:       data.BMC,
		SUBSTRATE: data.SUBSTRATE,
		CCMS:      data.CCMS,
		CCTMS:     data.CCTMS,
		CZCMS:     data.CZCMS,
		HAND:      data.HAND,
		DC:        data.DC,
		LF:        data.LF,
		LU:        data.LU,
		CSP:       data.CSP,
		DPC:       data.DPC,
		AMCP:      data.AMCP,
		AMCM:      data.AMCM,
		AMCK:      data.AMCK,
		AMCN:      data.AMCN,
		AMCS:      data.AMCS,
		AMCH:      data.AMCH,
		AMCO:      data.AMCO,
	}
	newData.Category = data.Category
	newData.SubCategory = data.SubCategory
	newData.OrderNumber = data.OrderNumber
	newData.Manufacturer = data.Manufacturer
	newData.Submitter = data.Submitter
	return newData
}

// DrillGunInsert 新增
func DrillGunInsert(data *DrillGun) (int64, error) {
	newData := drillGunDataInit(data)
	id, err := newData.DrillGunInsertToDB()
	if err != nil {
		return 0, err
	}
	return id, nil
}
