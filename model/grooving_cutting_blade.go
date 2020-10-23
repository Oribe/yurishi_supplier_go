package model

import (
	"fmt"

	"github.com/guregu/null"
)

// GroovingCuttingBlade 切槽切断刀片
type GroovingCuttingBlade struct {
	CutterBase
	BladeModel                    null.String `db:"bladeModel"`                    //	刀片型号
	MaterialLevel                 null.String `db:"materialLevel"`                 //	材质等级代码
	BladeInterfaceCode            null.String `db:"bladeInterfaceCode"`            //	刀片接口代码
	ManufacturerLogo              null.String `db:"manufacturerLogo"`              //	制造商标识
	Hand                          null.String `db:"hand"`                          //	切削方向
	BladeBladePadSizeCode         null.String `db:"bladeBladePadSizeCode"`         //	刀片/刀片垫尺寸代码
	CuttingEdgeCount              null.String `db:"cuttingEdgeCount"`              //	刀片可转位次数
	BladeCuttingWidth             null.String `db:"bladeCuttingWidth"`             //	刀片切削宽度
	BladeMaxCuttingDepth          null.String `db:"bladeMaxCuttingDepth"`          //	刀片最大切削深度
	ChipbreakerManufacturerLogo   null.String `db:"chipbreakerManufacturerLogo"`   //	断屑槽制造商标识
	P                             null.Int    `db:"P"`                             //	P 非合金钢、合金钢、铁素体和马氏体不锈钢
	M                             null.Int    `db:"M"`                             //	奥氏体不锈钢、双相不锈钢
	K                             null.Int    `db:"K"`                             //	铸块
	N                             null.Int    `db:"N"`                             //	有色金属
	S                             null.Int    `db:"S"`                             //	钛合金和高温合金
	H                             null.Int    `db:"H"`                             //	硬材料
	O                             null.Int    `db:"O"`                             //	非金属和符合材料
	OuterGroove                   null.Int    `db:"outerGroove"`                   //	外圆切槽
	CutOff                        null.Int    `db:"cutOff"`                        //	切断
	FaceGrooving                  null.Int    `db:"faceGrooving"`                  //	端面切槽
	ExternalCuttingSquareShoulder null.Int    `db:"externalCuttingSquareShoulder"` //	外圆切削-方肩
	ExternalProfileCutting        null.Int    `db:"externalProfileCutting"`        //	外圆仿形切削
	OuterBack                     null.Int    `db:"outerBack"`                     //	外圆背车
	InternalGroove                null.Int    `db:"internalGroove"`                //	內圆切槽
	InternalProfileCutting        null.Int    `db:"internalProfileCutting"`        //	內圆仿形切削
	InternalCuttingSquareShoulder null.Int    `db:"internalCuttingSquareShoulder"` //	內圆切削-方肩
	InternalBoring                null.Int    `db:"internalBoring"`                //	內圆背镗
}

// InsertToDB 添加
func (g *GroovingCuttingBlade) InsertToDB() (int64, error) {
	sql := `INSERT INTO
						grooving_cutting_blades
						(category,subCategory,manufacturer,orderNumber,bladeModel,materialLevel,bladeInterfaceCode,manufacturerLogo,hand,bladeBladePadSizeCode,cuttingEdgeCount,bladeCuttingWidth,bladeMaxCuttingDepth,chipbreakerManufacturerLogo,P,M,K,N,S,H,O,outerGroove,cutOff,faceGrooving,externalCuttingSquareShoulder,externalProfileCutting,outerBack,internalGroove,internalProfileCutting,internalCuttingSquareShoulder,internalBoring,submitter)
					VALUES
						(:category,:subCategory,:manufacturer,:orderNumber,:bladeModel,:materialLevel,:bladeInterfaceCode,:manufacturerLogo,:hand,:bladeBladePadSizeCode,:cuttingEdgeCount,:bladeCuttingWidth,:bladeMaxCuttingDepth,:chipbreakerManufacturerLogo,:P,:M,:K,:N,:S,:H,:O,:outerGroove,:cutOff,:faceGrooving,:externalCuttingSquareShoulder,:externalProfileCutting,:outerBack,:internalGroove,:internalProfileCutting,:internalCuttingSquareShoulder,:internalBoring,:submitter)`
	result, err := db.NamedExec(sql, g)
	if err != nil {
		error := fmt.Errorf("grooving_cutting_blade insert failed: %v", err.Error())
		return 0, error
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		error := fmt.Errorf("grooving_cutting_blade get last id failed: %v", err.Error())
		return lastID, error
	}
	return lastID, nil
}

// GroovingCuttingBladeQueryWithOrderNumber 查询
func GroovingCuttingBladeQueryWithOrderNumber(list *[]GroovingCuttingBlade, orderNumber, userID string) error {
	sql := `SELECT
						category,subCategory,manufacturer,orderNumber,bladeModel,materialLevel,bladeInterfaceCode,manufacturerLogo,hand,bladeBladePadSizeCode,cuttingEdgeCount,bladeCuttingWidth,bladeMaxCuttingDepth,chipbreakerManufacturerLogo,P,M,K,N,S,H,O,outerGroove,cutOff,faceGrooving,externalCuttingSquareShoulder,externalProfileCutting,outerBack,internalGroove,internalProfileCutting,internalCuttingSquareShoulder,internalBoring,submitter
					FROM
						grooving_cutting_blades
					WHERE
						orderNumber like CONCAT('%',?,'%')
					AND
						userID = ?`
	stmt, err := db.Preparex(sql)
	if err != nil {
		error := fmt.Errorf("grooving_cutting_blade prepare sql failed: %v", err.Error())
		return error
	}
	err = stmt.Select(list, orderNumber, userID)
	if err != nil {
		error := fmt.Errorf("grooving_cutting_blade query failed: %v", err.Error())
		return error
	}
	return nil
}
