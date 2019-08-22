// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model805

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block805 - Lithium-Ion Module Model -

const (
	ModelID          = 805
	ModelLabel       = "Lithium-Ion Module Model"
	ModelDescription = ""
)

const (
	CellSt         = "CellSt"
	CellTmp        = "CellTmp"
	CellTmpAvg     = "CellTmpAvg"
	CellTmpMax     = "CellTmpMax"
	CellTmpMaxCell = "CellTmpMaxCell"
	CellTmpMin     = "CellTmpMin"
	CellTmpMinCell = "CellTmpMinCell"
	CellV          = "CellV"
	CellVAvg       = "CellVAvg"
	CellVMax       = "CellVMax"
	CellVMaxCell   = "CellVMaxCell"
	CellVMin       = "CellVMin"
	CellVMinCell   = "CellVMinCell"
	CellV_SF       = "CellV_SF"
	DoD            = "DoD"
	DoD_SF         = "DoD_SF"
	ModIdx         = "ModIdx"
	NCell          = "NCell"
	NCellBal       = "NCellBal"
	NCyc           = "NCyc"
	SN             = "SN"
	SoC            = "SoC"
	SoC_SF         = "SoC_SF"
	SoH            = "SoH"
	SoH_SF         = "SoH_SF"
	StrIdx         = "StrIdx"
	Tmp_SF         = "Tmp_SF"
	V              = "V"
	V_SF           = "V_SF"
)

type Block805Repeat struct {
	CellV   uint16             `sunspec:"offset=0,sf=CellV_SF"`
	CellTmp int16              `sunspec:"offset=1,sf=Tmp_SF"`
	CellSt  sunspec.Bitfield32 `sunspec:"offset=2"`
}

type Block805 struct {
	StrIdx         uint16              `sunspec:"offset=0"`
	ModIdx         uint16              `sunspec:"offset=1"`
	NCell          uint16              `sunspec:"offset=2"`
	SoC            uint16              `sunspec:"offset=3,sf=SoC_SF"`
	DoD            uint16              `sunspec:"offset=4,sf=DoD_SF"`
	SoH            uint16              `sunspec:"offset=5,sf=SoH_SF"`
	NCyc           uint32              `sunspec:"offset=6"`
	V              uint16              `sunspec:"offset=8,sf=V_SF"`
	CellVMax       uint16              `sunspec:"offset=9,sf=CellV_SF"`
	CellVMaxCell   uint16              `sunspec:"offset=10"`
	CellVMin       uint16              `sunspec:"offset=11,sf=CellV_SF"`
	CellVMinCell   uint16              `sunspec:"offset=12"`
	CellVAvg       uint16              `sunspec:"offset=13,sf=CellV_SF"`
	CellTmpMax     int16               `sunspec:"offset=14,sf=Tmp_SF"`
	CellTmpMaxCell uint16              `sunspec:"offset=15"`
	CellTmpMin     int16               `sunspec:"offset=16,sf=Tmp_SF"`
	CellTmpMinCell uint16              `sunspec:"offset=17"`
	CellTmpAvg     int16               `sunspec:"offset=18,sf=Tmp_SF"`
	NCellBal       uint16              `sunspec:"offset=19"`
	SN             string              `sunspec:"offset=20,len=16"`
	SoC_SF         sunspec.ScaleFactor `sunspec:"offset=36"`
	SoH_SF         sunspec.ScaleFactor `sunspec:"offset=37"`
	DoD_SF         sunspec.ScaleFactor `sunspec:"offset=38"`
	V_SF           sunspec.ScaleFactor `sunspec:"offset=39"`
	CellV_SF       sunspec.ScaleFactor `sunspec:"offset=40"`
	Tmp_SF         sunspec.ScaleFactor `sunspec:"offset=41"`

	Repeats []Block805Repeat
}

func (self *Block805) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "lithium-ion-module",
		Length: 46,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 42,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: StrIdx, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "String Index", Description: "Index of the string containing the module."},
					smdx.PointElement{Id: ModIdx, Offset: 1, Type: typelabel.Uint16, Mandatory: true, Label: "Module Index", Description: "Index of the module within the string."},
					smdx.PointElement{Id: NCell, Offset: 2, Type: typelabel.Uint16, Mandatory: true, Label: "Module Cell Count", Description: "Count of all cells in the module."},
					smdx.PointElement{Id: SoC, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%", Label: "Module SoC", Description: "Module state of charge, expressed as a percentage."},
					smdx.PointElement{Id: DoD, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "DoD_SF", Units: "%", Label: "Depth of Discharge", Description: "Depth of discharge for the module."},
					smdx.PointElement{Id: SoH, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "SoH_SF", Units: "%", Label: "Module SoH", Description: "Module state of health, expressed as a percentage."},
					smdx.PointElement{Id: NCyc, Offset: 6, Type: typelabel.Uint32, Label: "Cycle Count", Description: "Count of cycles executed."},
					smdx.PointElement{Id: V, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true, Label: "Module Voltage", Description: "Voltage of the module."},
					smdx.PointElement{Id: CellVMax, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Max Cell Voltage", Description: "Maximum voltage for all cells in the module."},
					smdx.PointElement{Id: CellVMaxCell, Offset: 10, Type: typelabel.Uint16, Label: "Max Cell Voltage Cell", Description: "Cell with the maximum voltage."},
					smdx.PointElement{Id: CellVMin, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Min Cell Voltage", Description: "Minimum voltage for all cells in the module."},
					smdx.PointElement{Id: CellVMinCell, Offset: 12, Type: typelabel.Uint16, Label: "Min Cell Voltage Cell", Description: "Cell with the minimum voltage."},
					smdx.PointElement{Id: CellVAvg, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Average Cell Voltage", Description: "Average voltage for all cells in the module."},
					smdx.PointElement{Id: CellTmpMax, Offset: 14, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C", Mandatory: true, Label: "Max Cell Temperature", Description: "Maximum temperature for all cells in the module."},
					smdx.PointElement{Id: CellTmpMaxCell, Offset: 15, Type: typelabel.Uint16, Label: "Max Cell Temperature Cell", Description: "Cell with the maximum cell temperature."},
					smdx.PointElement{Id: CellTmpMin, Offset: 16, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C", Mandatory: true, Label: "Min Cell Temperature", Description: "Minimum temperature for all cells in the module."},
					smdx.PointElement{Id: CellTmpMinCell, Offset: 17, Type: typelabel.Uint16, Label: "Min Cell Temperature Cell", Description: "Cell with the minimum cell temperature."},
					smdx.PointElement{Id: CellTmpAvg, Offset: 18, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C", Mandatory: true, Label: "Average Cell Temperature", Description: "Average temperature for all cells in the module."},
					smdx.PointElement{Id: NCellBal, Offset: 19, Type: typelabel.Uint16, Label: "Balanced Cell Count", Description: "Number of cells currently being balanced in the module."},
					smdx.PointElement{Id: SN, Offset: 20, Type: typelabel.String, Length: 16, Label: "Serial Number", Description: "Serial number for the module."},
					smdx.PointElement{Id: SoC_SF, Offset: 36, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: SoH_SF, Offset: 37, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: DoD_SF, Offset: 38, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: V_SF, Offset: 39, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: CellV_SF, Offset: 40, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: Tmp_SF, Offset: 41, Type: typelabel.ScaleFactor, Mandatory: true},
				},
			},
			smdx.BlockElement{Name: "lithium-ion-module-cell",
				Length: 4,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: CellV, Offset: 0, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V", Mandatory: true, Label: "Cell Voltage", Description: "Cell terminal voltage."},
					smdx.PointElement{Id: CellTmp, Offset: 1, Type: typelabel.Int16, ScaleFactor: "Tmp_SF", Units: "C", Mandatory: true, Label: "Cell Temperature", Description: "Cell temperature."},
					smdx.PointElement{Id: CellSt, Offset: 2, Type: typelabel.Bitfield32, Label: "Cell Status", Description: "Status of the cell."},
				},
			},
		}})
}
