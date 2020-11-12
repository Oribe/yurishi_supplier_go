package model

import (
	"fmt"
	"manufacture_supplier_go/util"

	"github.com/guregu/null"
)

// GroovingCuttingIntegralTool 整体式切槽切断
type GroovingCuttingIntegralTool struct {
	CutterBase
	ToolModel                       null.String `db:"toolModel" json:"toolModel"`                                             // 刀具型号
	MaterialLevel                   null.String `db:"materialLevel" json:"materialLevel"`                                     // 材质等级代码
	SideInterfaceCode               null.String `db:"sideInterfaceCode" json:"sideInterfaceCode"`                             // 机床侧接口代码
	ManufacturerLogo                null.String `db:"manufacturerLogo" json:"manufacturerLogo"`                               // 制造商标识
	ConnectionDiameter              null.String `db:"connectionDiameter" json:"connectionDiameter"`                           // 连接直径
	Hand                            null.String `db:"hand" json:"hand"`                                                       // 切削方向
	FunctionLength                  null.String `db:"functionLength" json:"functionLength"`                                   // 刀片可转位次数
	CuttingEdgeCount                null.String `db:"cuttingEdgeCount" json:"cuttingEdgeCount"`                               // 功能长度
	AvailableLength                 null.String `db:"availableLength" json:"availableLength"`                                 // 可用长度
	BladeCuttingWidth               null.String `db:"bladeCuttingWidth" json:"bladeCuttingWidth"`                             // 刀片切削宽度
	AllowableMinimumCuttingDiameter null.String `db:"allowableMinimumCuttingDiameter" json:"allowableMinimumCuttingDiameter"` // 允许最小切削直径
	BladeMaxCuttingDepth            null.String `db:"bladeMaxCuttingDepth" json:"bladeMaxCuttingDepth"`                       // 刀片最大切削深度
	CoolantSupplyCharacteristics    null.Int    `db:"coolantSupplyCharacteristics" json:"coolantSupplyCharacteristics"`       // 冷却液供给特性
	P                               null.Int    `db:"P" json:"P"`                                                             // p 非合金钢、合金钢、铁素体和马氏体不锈钢
	M                               null.Int    `db:"M" json:"M"`                                                             // 奥氏体不锈钢、双相不锈钢
	K                               null.Int    `db:"K" json:"K"`                                                             // 铸块
	N                               null.Int    `db:"N" json:"N"`                                                             // 有色金属
	S                               null.Int    `db:"S" json:"S"`                                                             // 钛合金和高温合金
	H                               null.Int    `db:"H" json:"H"`                                                             // 硬材料
	O                               null.Int    `db:"O" json:"O"`                                                             // 非金属和符合材料
	InternalGroove                  null.Int    `db:"internalGroove" json:"internalGroove"`                                   // 內圆切槽
	InternalProfileCutting          null.Int    `db:"internalProfileCutting" json:"internalProfileCutting"`                   // 內圆仿形切削
	SquareShoulder                  null.Int    `db:"squareShoulder" json:"squareShoulder"`                                   // 内圆车削-方肩
	InternalBoring                  null.Int    `db:"internalBoring" json:"internalBoring"`                                   // 內圆背镗
	CutOff                          null.Int    `db:"cutOff" json:"cutOff"`                                                   // 切断
	FaceGrooving                    null.Int    `db:"faceGrooving" json:"faceGrooving"`                                       // 端面切槽
}

func (g *GroovingCuttingIntegralTool) insertToDB() error {
	sql := `INSERT INTO
						grooving_cutting_integral_tool_suppliers
						(category,subCategory,manufacturer,orderNumber,toolModel,materialLevel,sideInterfaceCode,manufacturerLogo,connectionDiameter,hand,functionLength,cuttingEdgeCount,availableLength,bladeCuttingWidth,allowableMinimumCuttingDiameter,bladeMaxCuttingDepth,coolantSupplyCharacteristics,P,M,K,N,S,H,O,internalGroove,internalProfileCutting,squareShoulder,internalBoring,cutOff,faceGrooving,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:toolModel,:materialLevel,:sideInterfaceCode,:manufacturerLogo,:connectionDiameter,:hand,:functionLength,:cuttingEdgeCount,:availableLength,:bladeCuttingWidth,:allowableMinimumCuttingDiameter,:bladeMaxCuttingDepth,:coolantSupplyCharacteristics,:P,:M,:K,:N,:S,:H,:O,:internalGroove,:internalProfileCutting,:squareShoulder,:internalBoring,:cutOff,:faceGrooving,:submitter)		
	`
	_, err := db.NamedExec(sql, g)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingIntegralTool insert failed: %v", err.Error())
		return error
	}
	return nil
}

func (g *GroovingCuttingIntegralTool) updateToDB() error {
	sql := `UPDATE 
						grooving_cutting_integral_tool_suppliers
					SET
						category=:category,subCategory=:subCategory,manufacturer=:manufacturer,orderNumber=:orderNumber,toolModel=:toolModel,materialLevel=:materialLevel,sideInterfaceCode=:sideInterfaceCode,manufacturerLogo=:manufacturerLogo,connectionDiameter=:connectionDiameter,
						hand=:hand,functionLength=:functionLength,cuttingEdgeCount=:cuttingEdgeCount,availableLength=:availableLength,bladeCuttingWidth=:bladeCuttingWidth,allowableMinimumCuttingDiameter=:allowableMinimumCuttingDiameter,bladeMaxCuttingDepth=:bladeMaxCuttingDepth,coolantSupplyCharacteristics=:coolantSupplyCharacteristics,
						P=:P,M=:M,K=:K,N=:N,S=:S,H=:H,O=:O,internalGroove=:internalGroove,internalProfileCutting=:internalProfileCutting,squareShoulder=:squareShoulder,internalBoring=:internalBoring,cutOff=:cutOff,faceGrooving=:faceGrooving,submitter=:submitter
					WHERE
						id = :id`
	_, err := db.NamedExec(sql, g)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingIntegralTool update failed: %v", err.Error())
		return error
	}
	return nil
}

func (g *GroovingCuttingIntegralTool) dataIsExit() (bool, error) {
	if g.ID < 0 {
		return false, nil
	}
	r := GroovingCuttingIntegralTool{}
	sql := `SELECT * FROM grooving_cutting_integral_tool_suppliers WHERE id = ? LIMIT 1`
	err := db.Get(&r, sql, g.ID)
	if err != nil {
		return false, err
	}
	if r.ID == g.ID { // 存在
		return true, nil
	}
	return false, nil
}

// Upsert 新增或更新
func (g *GroovingCuttingIntegralTool) Upsert() (int, error) {
	b, err := g.dataIsExit()
	if err != nil {
		return -1, err
	}
	if b { // 存在
		return util.UPDATE, g.updateToDB()
	}
	return util.INSERT, g.insertToDB()
}

// GroovingCuttingIntegralToolQueryByOrderNumber 查询
func GroovingCuttingIntegralToolQueryByOrderNumber(rows *[]GroovingCuttingIntegralTool, orderNumber, userID string) error {
	sql := `SELECT
						id,category,subCategory,manufacturer,orderNumber,toolModel,materialLevel,sideInterfaceCode,manufacturerLogo,connectionDiameter,hand,functionLength,cuttingEdgeCount,availableLength,bladeCuttingWidth,allowableMinimumCuttingDiameter,bladeMaxCuttingDepth,coolantSupplyCharacteristics,P,M,K,N,S,H,O,internalGroove,internalProfileCutting,squareShoulder,internalBoring,cutOff,faceGrooving,submitter
					FROM
						grooving_cutting_integral_tool_suppliers
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						submitter = ?
					`
	err := db.Select(rows, sql, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("GroovingCuttingIntegralTool query failed: %v", err.Error())
		return error
	}
	return nil
}
