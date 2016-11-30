// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model804

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block804 - Lithium-Ion String Model -

const (
	ModelID = 804
)

const (
	A                 = "A"
	A_SF              = "A_SF"
	CellVAvg          = "CellVAvg"
	CellVMax          = "CellVMax"
	CellVMaxMod       = "CellVMaxMod"
	CellVMin          = "CellVMin"
	CellVMinMod       = "CellVMinMod"
	CellV_SF          = "CellV_SF"
	ConFail           = "ConFail"
	DoD               = "DoD"
	DoD_SF            = "DoD_SF"
	Evt1              = "Evt1"
	Evt2              = "Evt2"
	EvtVnd1           = "EvtVnd1"
	EvtVnd2           = "EvtVnd2"
	Idx               = "Idx"
	ModCellTmpAvg     = "ModCellTmpAvg"
	ModCellTmpMax     = "ModCellTmpMax"
	ModCellTmpMaxCell = "ModCellTmpMaxCell"
	ModCellTmpMin     = "ModCellTmpMin"
	ModCellTmpMinCell = "ModCellTmpMinCell"
	ModCellVAvg       = "ModCellVAvg"
	ModCellVMax       = "ModCellVMax"
	ModCellVMaxCell   = "ModCellVMaxCell"
	ModCellVMin       = "ModCellVMin"
	ModCellVMinCell   = "ModCellVMinCell"
	ModNCell          = "ModNCell"
	ModSoC            = "ModSoC"
	ModSoH            = "ModSoH"
	ModTmpAvg         = "ModTmpAvg"
	ModTmpMax         = "ModTmpMax"
	ModTmpMaxMod      = "ModTmpMaxMod"
	ModTmpMin         = "ModTmpMin"
	ModTmpMinMod      = "ModTmpMinMod"
	ModTmp_SF         = "ModTmp_SF"
	NCellBal          = "NCellBal"
	NCyc              = "NCyc"
	NMod              = "NMod"
	Pad2              = "Pad2"
	Pad3              = "Pad3"
	Pad4              = "Pad4"
	SetCon            = "SetCon"
	SetEna            = "SetEna"
	SoC               = "SoC"
	SoC_SF            = "SoC_SF"
	SoH               = "SoH"
	SoH_SF            = "SoH_SF"
	St                = "St"
	V                 = "V"
	V_SF              = "V_SF"
)

type Block804Repeat struct {
	ModNCell          uint16      `sunspec:"offset=0"`
	ModSoC            uint16      `sunspec:"offset=1,sf=SoC_SF"`
	ModSoH            uint16      `sunspec:"offset=2,sf=SoH_SF"`
	ModCellVMax       uint16      `sunspec:"offset=3,sf=CellV_SF"`
	ModCellVMaxCell   uint16      `sunspec:"offset=4"`
	ModCellVMin       uint16      `sunspec:"offset=5,sf=CellV_SF"`
	ModCellVMinCell   uint16      `sunspec:"offset=6,sf=CellV_SF"`
	ModCellVAvg       uint16      `sunspec:"offset=7,sf=CellV_SF"`
	ModCellTmpMax     int16       `sunspec:"offset=8,sf=ModTmp_SF"`
	ModCellTmpMaxCell uint16      `sunspec:"offset=9"`
	ModCellTmpMin     int16       `sunspec:"offset=10,sf=ModTmp_SF"`
	ModCellTmpMinCell uint16      `sunspec:"offset=11"`
	ModCellTmpAvg     int16       `sunspec:"offset=12,sf=ModTmp_SF"`
	Pad2              sunspec.Pad `sunspec:"offset=13"`
	Pad3              sunspec.Pad `sunspec:"offset=14"`
	Pad4              sunspec.Pad `sunspec:"offset=15"`
}

type Block804 struct {
	Idx          uint16              `sunspec:"offset=0"`
	NMod         uint16              `sunspec:"offset=1"`
	St           sunspec.Bitfield32  `sunspec:"offset=2"`
	ConFail      sunspec.Enum16      `sunspec:"offset=4"`
	NCellBal     uint16              `sunspec:"offset=5"`
	SoC          uint16              `sunspec:"offset=6,sf=SoC_SF"`
	DoD          uint16              `sunspec:"offset=7,sf=StrDoC_SF"`
	NCyc         uint32              `sunspec:"offset=8"`
	SoH          uint16              `sunspec:"offset=9,sf=SoH_SF"`
	A            int16               `sunspec:"offset=10,sf=A_SF"`
	V            uint16              `sunspec:"offset=11,sf=V_SF"`
	CellVMax     uint16              `sunspec:"offset=12,sf=CellV_SF"`
	CellVMaxMod  uint16              `sunspec:"offset=13"`
	CellVMin     uint16              `sunspec:"offset=14,sf=CellV_SF"`
	CellVMinMod  uint16              `sunspec:"offset=15"`
	CellVAvg     uint16              `sunspec:"offset=16,sf=CellV_SF"`
	ModTmpMax    int16               `sunspec:"offset=17,sf=ModTmp_SF"`
	ModTmpMaxMod uint16              `sunspec:"offset=18"`
	ModTmpMin    int16               `sunspec:"offset=19,sf=ModTmp_SF"`
	ModTmpMinMod uint16              `sunspec:"offset=20"`
	ModTmpAvg    int16               `sunspec:"offset=21,sf=ModTmp_SF"`
	Evt1         sunspec.Bitfield32  `sunspec:"offset=22"`
	Evt2         sunspec.Bitfield32  `sunspec:"offset=24"`
	EvtVnd1      sunspec.Bitfield32  `sunspec:"offset=26"`
	EvtVnd2      sunspec.Bitfield32  `sunspec:"offset=28"`
	SetEna       sunspec.Enum16      `sunspec:"offset=29,access=rw"`
	SetCon       sunspec.Enum16      `sunspec:"offset=30,access=rw"`
	SoC_SF       sunspec.ScaleFactor `sunspec:"offset=31"`
	SoH_SF       sunspec.ScaleFactor `sunspec:"offset=32"`
	DoD_SF       sunspec.ScaleFactor `sunspec:"offset=33"`
	A_SF         sunspec.ScaleFactor `sunspec:"offset=34"`
	V_SF         sunspec.ScaleFactor `sunspec:"offset=35"`
	CellV_SF     sunspec.ScaleFactor `sunspec:"offset=36"`
	ModTmp_SF    sunspec.ScaleFactor `sunspec:"offset=37"`

	Repeats []Block804Repeat
}

func (self *Block804) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "lithium_ion_string",
		Length: 56,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 38,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Idx, Offset: 0, Type: typelabel.Uint16, Mandatory: true},
					smdx.PointElement{Id: NMod, Offset: 1, Type: typelabel.Uint16, Mandatory: true},
					smdx.PointElement{Id: St, Offset: 2, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: ConFail, Offset: 4, Type: typelabel.Enum16},
					smdx.PointElement{Id: NCellBal, Offset: 5, Type: typelabel.Uint16},
					smdx.PointElement{Id: SoC, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%", Mandatory: true},
					smdx.PointElement{Id: DoD, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "StrDoC_SF", Units: "%"},
					smdx.PointElement{Id: NCyc, Offset: 8, Type: typelabel.Uint32},
					smdx.PointElement{Id: SoH, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "SoH_SF", Units: "%"},
					smdx.PointElement{Id: A, Offset: 10, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: V, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: CellVMax, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: CellVMaxMod, Offset: 13, Type: typelabel.Uint16},
					smdx.PointElement{Id: CellVMin, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: CellVMinMod, Offset: 15, Type: typelabel.Uint16},
					smdx.PointElement{Id: CellVAvg, Offset: 16, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: ModTmpMax, Offset: 17, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: ModTmpMaxMod, Offset: 18, Type: typelabel.Uint16, Mandatory: true},
					smdx.PointElement{Id: ModTmpMin, Offset: 19, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: ModTmpMinMod, Offset: 20, Type: typelabel.Uint16, Mandatory: true},
					smdx.PointElement{Id: ModTmpAvg, Offset: 21, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: Evt1, Offset: 22, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: Evt2, Offset: 24, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: EvtVnd1, Offset: 26, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: EvtVnd2, Offset: 28, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: SetEna, Offset: 29, Type: typelabel.Enum16, Access: "rw"},
					smdx.PointElement{Id: SetCon, Offset: 30, Type: typelabel.Enum16, Access: "rw"},
					smdx.PointElement{Id: SoC_SF, Offset: 31, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: SoH_SF, Offset: 32, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: DoD_SF, Offset: 33, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: A_SF, Offset: 34, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: V_SF, Offset: 35, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: CellV_SF, Offset: 36, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: ModTmp_SF, Offset: 37, Type: typelabel.ScaleFactor, Mandatory: true},
				},
			},
			smdx.BlockElement{Name: "lithium_ion_string_module",
				Length: 16,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ModNCell, Offset: 0, Type: typelabel.Uint16, Mandatory: true},
					smdx.PointElement{Id: ModSoC, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%"},
					smdx.PointElement{Id: ModSoH, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "SoH_SF", Units: "%"},
					smdx.PointElement{Id: ModCellVMax, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: ModCellVMaxCell, Offset: 4, Type: typelabel.Uint16},
					smdx.PointElement{Id: ModCellVMin, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: ModCellVMinCell, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: ModCellVAvg, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: ModCellTmpMax, Offset: 8, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: ModCellTmpMaxCell, Offset: 9, Type: typelabel.Uint16},
					smdx.PointElement{Id: ModCellTmpMin, Offset: 10, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: ModCellTmpMinCell, Offset: 11, Type: typelabel.Uint16},
					smdx.PointElement{Id: ModCellTmpAvg, Offset: 12, Type: typelabel.Int16, ScaleFactor: "ModTmp_SF", Units: "C", Mandatory: true},
					smdx.PointElement{Id: Pad2, Offset: 13, Type: typelabel.Pad, Mandatory: true},
					smdx.PointElement{Id: Pad3, Offset: 14, Type: typelabel.Pad, Mandatory: true},
					smdx.PointElement{Id: Pad4, Offset: 15, Type: typelabel.Pad, Mandatory: true},
				},
			},
		}})
}
