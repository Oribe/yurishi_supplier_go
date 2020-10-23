package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// GroovingCuttingBlade 切槽切断刀片
type GroovingCuttingBlade struct {
	model.CutterBase
	BladeModel                    null.String `json:"bladeModel"`                    //	刀片型号
	MaterialLevel                 null.String `json:"materialLevel"`                 //	材质等级代码
	BladeInterfaceCode            null.String `json:"bladeInterfaceCode"`            //	刀片接口代码
	ManufacturerLogo              null.String `json:"manufacturerLogo"`              //	制造商标识
	Hand                          null.String `json:"hand"`                          //	切削方向
	BladeBladePadSizeCode         null.String `json:"bladeBladePadSizeCode"`         //	刀片/刀片垫尺寸代码
	CuttingEdgeCount              null.String `json:"cuttingEdgeCount"`              //	刀片可转位次数
	BladeCuttingWidth             null.String `json:"bladeCuttingWidth"`             //	刀片切削宽度
	BladeMaxCuttingDepth          null.String `json:"bladeMaxCuttingDepth"`          //	刀片最大切削深度
	ChipbreakerManufacturerLogo   null.String `json:"chipbreakerManufacturerLogo"`   //	断屑槽制造商标识
	P                             null.Int    `json:"P"`                             //	P 非合金钢、合金钢、铁素体和马氏体不锈钢
	M                             null.Int    `json:"M"`                             //	奥氏体不锈钢、双相不锈钢
	K                             null.Int    `json:"K"`                             //	铸块
	N                             null.Int    `json:"N"`                             //	有色金属
	S                             null.Int    `json:"S"`                             //	钛合金和高温合金
	H                             null.Int    `json:"H"`                             //	硬材料
	O                             null.Int    `json:"O"`                             //	非金属和符合材料
	OuterGroove                   null.Int    `json:"outerGroove"`                   //	外圆切槽
	CutOff                        null.Int    `json:"cutOff"`                        //	切断
	FaceGrooving                  null.Int    `json:"faceGrooving"`                  //	端面切槽
	ExternalCuttingSquareShoulder null.Int    `json:"externalCuttingSquareShoulder"` //	外圆切削-方肩
	ExternalProfileCutting        null.Int    `json:"externalProfileCutting"`        //	外圆仿形切削
	OuterBack                     null.Int    `json:"outerBack"`                     //	外圆背车
	InternalGroove                null.Int    `json:"internalGroove"`                //	內圆切槽
	InternalProfileCutting        null.Int    `json:"internalProfileCutting"`        //	內圆仿形切削
	InternalCuttingSquareShoulder null.Int    `json:"internalCuttingSquareShoulder"` //	內圆切削-方肩
	InternalBoring                null.Int    `json:"internalBoring"`                //	內圆背镗
}

func groovingCuttingBladeRowInitial(body GroovingCuttingBlade) model.GroovingCuttingBlade {
	row := model.GroovingCuttingBlade{
		BladeModel:                    body.BladeModel,
		MaterialLevel:                 body.MaterialLevel,
		ManufacturerLogo:              body.ManufacturerLogo,
		Hand:                          body.Hand,
		BladeBladePadSizeCode:         body.BladeBladePadSizeCode,
		CuttingEdgeCount:              body.CuttingEdgeCount,
		BladeCuttingWidth:             body.BladeCuttingWidth,
		BladeMaxCuttingDepth:          body.BladeMaxCuttingDepth,
		ChipbreakerManufacturerLogo:   body.ChipbreakerManufacturerLogo,
		P:                             body.P,
		M:                             body.M,
		K:                             body.K,
		N:                             body.N,
		S:                             body.S,
		H:                             body.H,
		O:                             body.O,
		OuterGroove:                   body.OuterGroove,
		CutOff:                        body.CutOff,
		FaceGrooving:                  body.FaceGrooving,
		ExternalCuttingSquareShoulder: body.ExternalCuttingSquareShoulder,
		ExternalProfileCutting:        body.ExternalProfileCutting,
		OuterBack:                     body.OuterBack,
		InternalGroove:                body.InternalGroove,
		InternalProfileCutting:        body.InternalProfileCutting,
		InternalCuttingSquareShoulder: body.InternalCuttingSquareShoulder,
		InternalBoring:                body.InternalBoring,
	}
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.Manufacturer = body.Manufacturer
	row.OrderNumber = body.OrderNumber
	row.Submitter = body.Submitter
	return row
}

// GroovingCuttingBladeInsert 添加到数据库
func GroovingCuttingBladeInsert(body GroovingCuttingBlade) (int64, error) {
	row := groovingCuttingBladeRowInitial(body)
	return row.InsertToDB()
}

// GroovingCuttingBladeQueryWithOrderNumber 查询
func GroovingCuttingBladeQueryWithOrderNumber(orderNumber, userID string) ([]model.GroovingCuttingBlade, error) {
	rows := []model.GroovingCuttingBlade{}
	err := model.GroovingCuttingBladeQueryWithOrderNumber(&rows, orderNumber, userID)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
