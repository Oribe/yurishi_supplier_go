package constant

const (
	Turning         = 0x0100 // 车削
	Milling         = 0x0200 // 铣削
	Drill           = 0x0300 // 钻削
	Boring          = 0x0400 // 镗削
	Reaming         = 0x0500 // 铰削
	Screw           = 0x0600 // 螺纹
	Tool            = 0x0700 // 工具
	GroovingCutting = 0x0800 // 切槽切断
)

const (
	TurningOutter   = Turning | 0x01 // 外圆车削刀体
	TurningInner    = Turning | 0x02 // 内园车削刀具
	TurningInsert   = Turning | 0x04 // 车削刀片
	TurningIntegral = Turning | 0x05 // 整体式车削刀具

	MillingIndexable                    = Milling | 0x01 // 可转位铣刀
	MillingRepHead                      = Milling | 0x02 // 可更换铣刀头
	MillingIntegral                     = Milling | 0x03 // 整体式铣刀
	MillingIndexableInsert              = Milling | 0x04 // 可转位铣刀片
	MillingWeldingEdgeCutter            = Milling | 0x05 // 焊接刃铣刀片
	MillingMandrelModularIntegralCutter = Milling | 0x06 // 芯轴和模块化整体铣刀

	DrillIntegral      = Drill | 0x01 // 整体式钻头
	DrillReplace       = Drill | 0x02 // 可更换钻削刀头
	DrillRotateBlade   = Drill | 0x03 // 可转位钻头片
	DrillReplaceBody   = Drill | 0x04 // 可更换刀头钻削刀体
	DrillGun           = Drill | 0x05 // 枪钻
	DrillIndexableBody = Drill | 0x06 // 可转位钻头刀体

	BoringPhase    = Boring | 0x01 // 粗镗刀具
	BoringFine     = Boring | 0x02 // 精镗刀具
	BoringBigDSlot = Boring | 0x03 // 大直径桥板镗刀
	BoringBlade    = Boring | 0x04 // 镗削刀片
	BoringPathTool = Boring | 0x05 // 小径镗削刀具
	BoringHead     = Boring | 0x06 // 镗头
	KnifeHolder    = Boring | 0x07 // 刀夹

	ReamingIntegral        = Reaming | 0x01 // 整体式铰刀
	ReamingReplaceHead     = Reaming | 0x02 // 可更换铰刀刀头
	ReamingAdjSingleBlade  = Reaming | 0x03 // 单刀片可调绞刀
	ReamingReplaceHeadBody = Reaming | 0x04 // 可更换刀头铰削刀体
	ReamingBlade           = Reaming | 0x05 // 铰削刀片
	ReamingBladeCutterBody = Reaming | 0x06 // 刀片式铰削刀具刀体

	ScrewTurning                      = Screw | 0x01  // 螺纹车削刀具
	ScrewMilling                      = Screw | 0x02  // 螺纹铣削刀具
	ScrewTap                          = Screw | 0x03  // 丝锥
	ScrewIntegralMilling              = Screw | 0x04  // 整体式螺纹铣刀
	ScrewReplaceHeadMillingBody       = Screw | 0x05  // 可更换刀头螺纹铣刀体
	ScrewReplaceMillingHead           = Screw | 0x06  // 可更换螺纹铣刀头
	ScrewTurningInsert                = Screw | 0x07  // 螺纹车削刀片
	ScrewIntegralTurningTool          = Screw | 0x08  // 整体式螺纹车削刀具
	ScrewThreadMillingInsert          = Screw | 0x09  // 螺纹铣刀片
	ScrewBladeThreadMillingBody       = Screw | 0x10  // 刀片式螺纹铣刀刀体
	ScrewExternalThreadTurningBody    = Screw | 0x11  // 外螺纹车削刀体
	ScrewInternalThreadTurningBody    = Screw | 0x12  // 内螺纹车削刀体
	ScrewReplaceHeadThreadTurningBody = Screw | 0x13  // 可更换刀头螺纹车削刀体
	ScrewReplaceThreadTurningHead     = Screw | 0x14  // 可更换螺纹车削刀头
	ScrewDie                          = Screw | 0x015 // 板牙

	ToolRotateHandle               = Tool | 0x01 // 旋转刀柄系统
	toolTurningHolderAdapterSystem = Tool | 0x02 // 固定转接头
	ToolSleeve                     = Tool | 0x03 // 套筒与夹头
	ToolRivet                      = Tool | 0x04 // 拉钉

	GroovingCuttingIntegralTool = GroovingCutting | 0x01 // 整体式切槽切断
	GroovingCuttingBlade        = GroovingCutting | 0x02 // 切槽切断刀片
	GroovingExternalBody        = GroovingCutting | 0x03 // 外圆切槽切断刀体
	GroovingInternalBody        = GroovingCutting | 0x04 // 內圆切槽切断刀体
)
