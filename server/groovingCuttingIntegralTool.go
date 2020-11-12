package server

import (
	"manufacture_supplier_go/model"

	"github.com/guregu/null"
)

// GroovingCuttingIntegralTool 整体式切槽切断
type GroovingCuttingIntegralTool struct {
	model.CutterBase
	ToolModel                       null.String
	MaterialLevel                   null.String
	SideInterfaceCode               null.String
	ManufacturerLogo                null.String
	ConnectionDiameter              null.String
	Hand                            null.String
	FunctionLength                  null.String
	CuttingEdgeCount                null.String
	AvailableLength                 null.String
	BladeCuttingWidth               null.String
	AllowableMinimumCuttingDiameter null.String
	BladeMaxCuttingDepth            null.String
	CoolantSupplyCharacteristics    null.Int
	P                               null.Int
	M                               null.Int
	K                               null.Int
	N                               null.Int
	S                               null.Int
	H                               null.Int
	O                               null.Int
	InternalGroove                  null.Int
	InternalProfileCutting          null.Int
	SquareShoulder                  null.Int
	InternalBoring                  null.Int
	CutOff                          null.Int
	FaceGrooving                    null.Int
}

func groovingCuttingIntegralToolRowInitial(body GroovingCuttingIntegralTool) *model.GroovingCuttingIntegralTool {
	row := &model.GroovingCuttingIntegralTool{
		ToolModel:                       body.ToolModel,
		MaterialLevel:                   body.MaterialLevel,
		SideInterfaceCode:               body.SideInterfaceCode,
		ManufacturerLogo:                body.ManufacturerLogo,
		ConnectionDiameter:              body.ConnectionDiameter,
		Hand:                            body.Hand,
		FunctionLength:                  body.FunctionLength,
		CuttingEdgeCount:                body.CuttingEdgeCount,
		AvailableLength:                 body.AvailableLength,
		BladeCuttingWidth:               body.BladeCuttingWidth,
		AllowableMinimumCuttingDiameter: body.AllowableMinimumCuttingDiameter,
		BladeMaxCuttingDepth:            body.BladeMaxCuttingDepth,
		CoolantSupplyCharacteristics:    body.CoolantSupplyCharacteristics,
		P:                               body.P,
		M:                               body.M,
		K:                               body.K,
		N:                               body.N,
		S:                               body.S,
		H:                               body.H,
		O:                               body.O,
		InternalGroove:                  body.InternalGroove,
		InternalProfileCutting:          body.InternalProfileCutting,
		SquareShoulder:                  body.SquareShoulder,
		InternalBoring:                  body.InternalBoring,
		CutOff:                          body.CutOff,
		FaceGrooving:                    body.FaceGrooving,
	}
	row.ID = body.ID
	row.Category = body.Category
	row.SubCategory = body.SubCategory
	row.OrderNumber = body.OrderNumber
	row.Manufacturer = body.Manufacturer
	row.Submitter = body.Submitter
	return row
}

// GroovingCuttingIntegralToolUpsert 更新或者新增
func GroovingCuttingIntegralToolUpsert(body GroovingCuttingIntegralTool) (int, error) {
	row := groovingCuttingIntegralToolRowInitial(body)
	return row.Upsert()
}

// GroovingCuttingIntegralToolQueryByOrderNumber 查询
func GroovingCuttingIntegralToolQueryByOrderNumber(orderNumber, userID string) ([]model.GroovingCuttingIntegralTool, error) {
	rows := []model.GroovingCuttingIntegralTool{}
	err := model.GroovingCuttingIntegralToolQueryByOrderNumber(&rows, orderNumber, userID)
	return rows, err
}
