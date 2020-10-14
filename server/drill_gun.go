package server

import (
	"manufacture_supplier_go/model"
)

// DrillGunQueryWithOrder 查询
func DrillGunQueryWithOrder(orderNumber string, userID int) ([]model.DrillGun, error) {
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
func DrillGunInsert(data *DrillGun, userID int) (int64, error) {
	newData := drillGunDataInit(data)
	id, err := newData.DrillGunInsertToDB()
	if err != nil {
		return 0, err
	}
	return id, nil
}
