// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model403

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block403 - String Combiner (Current) - A basic string combiner model

const (
	ModelID          = 403
	ModelLabel       = "String Combiner (Current)"
	ModelDescription = "A basic string combiner model"
)

const (
	DCA        = "DCA"
	DCAMax     = "DCAMax"
	DCA_SF     = "DCA_SF"
	DCAhr      = "DCAhr"
	DCAhr_SF   = "DCAhr_SF"
	DCV        = "DCV"
	DCV_SF     = "DCV_SF"
	Evt        = "Evt"
	EvtVnd     = "EvtVnd"
	InDCA      = "InDCA"
	InDCA_SF   = "InDCA_SF"
	InDCAhr    = "InDCAhr"
	InDCAhr_SF = "InDCAhr_SF"
	InEvt      = "InEvt"
	InEvtVnd   = "InEvtVnd"
	InID       = "InID"
	N          = "N"
	Tmp        = "Tmp"
)

type Block403Repeat struct {
	InID     uint16             `sunspec:"offset=0"`
	InEvt    sunspec.Bitfield32 `sunspec:"offset=1"`
	InEvtVnd sunspec.Bitfield32 `sunspec:"offset=3"`
	InDCA    int16              `sunspec:"offset=5,sf=InDCA_SF"`
	InDCAhr  sunspec.Acc32      `sunspec:"offset=6,sf=InDCAhr_SF"`
}

type Block403 struct {
	DCA_SF     sunspec.ScaleFactor `sunspec:"offset=0"`
	DCAhr_SF   sunspec.ScaleFactor `sunspec:"offset=1"`
	DCV_SF     sunspec.ScaleFactor `sunspec:"offset=2"`
	DCAMax     uint16              `sunspec:"offset=3,sf=DCA_SF"`
	N          sunspec.Count       `sunspec:"offset=4"`
	Evt        sunspec.Bitfield32  `sunspec:"offset=5"`
	EvtVnd     sunspec.Bitfield32  `sunspec:"offset=7"`
	DCA        int16               `sunspec:"offset=9,sf=DCA_SF"`
	DCAhr      sunspec.Acc32       `sunspec:"offset=10,sf=DCAhr_SF"`
	DCV        int16               `sunspec:"offset=12,sf=DCV_SF"`
	Tmp        int16               `sunspec:"offset=13"`
	InDCA_SF   sunspec.ScaleFactor `sunspec:"offset=14"`
	InDCAhr_SF sunspec.ScaleFactor `sunspec:"offset=15"`

	Repeats []Block403Repeat
}

func (self *Block403) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "string_combiner",
		Length: 24,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 16,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: DCA_SF, Offset: 0, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: DCAhr_SF, Offset: 1, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: DCV_SF, Offset: 2, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: DCAMax, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "DCA_SF", Units: "A", Mandatory: true, Label: "Rating", Description: "Maximum DC Current Rating"},
					smdx.PointElement{Id: N, Offset: 4, Type: typelabel.Count, Mandatory: true, Label: "N", Description: "Number of Inputs"},
					smdx.PointElement{Id: Evt, Offset: 5, Type: typelabel.Bitfield32, Mandatory: true, Label: "Event", Description: "Bitmask value.  Events"},
					smdx.PointElement{Id: EvtVnd, Offset: 7, Type: typelabel.Bitfield32, Label: "Vendor Event", Description: "Bitmask value.  Vendor defined events"},
					smdx.PointElement{Id: DCA, Offset: 9, Type: typelabel.Int16, ScaleFactor: "DCA_SF", Units: "A", Mandatory: true, Label: "Amps", Description: "Total measured current"},
					smdx.PointElement{Id: DCAhr, Offset: 10, Type: typelabel.Acc32, ScaleFactor: "DCAhr_SF", Units: "Ah", Label: "Amp-hours", Description: "Total metered Amp-hours"},
					smdx.PointElement{Id: DCV, Offset: 12, Type: typelabel.Int16, ScaleFactor: "DCV_SF", Units: "V", Label: "Voltage", Description: "Output Voltage"},
					smdx.PointElement{Id: Tmp, Offset: 13, Type: typelabel.Int16, Units: "C", Label: "Temp", Description: "Internal operating temperature"},
					smdx.PointElement{Id: InDCA_SF, Offset: 14, Type: typelabel.ScaleFactor},
					smdx.PointElement{Id: InDCAhr_SF, Offset: 15, Type: typelabel.ScaleFactor},
				},
			},
			smdx.BlockElement{Name: "string",
				Length: 8,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: InID, Offset: 0, Type: typelabel.Uint16, Mandatory: true, Label: "ID", Description: "Uniquely identifies this input set"},
					smdx.PointElement{Id: InEvt, Offset: 1, Type: typelabel.Bitfield32, Mandatory: true, Label: "Input Event", Description: "String Input Event Flags"},
					smdx.PointElement{Id: InEvtVnd, Offset: 3, Type: typelabel.Bitfield32, Label: "Input Event Vendor", Description: "String Input Vendor Event Flags"},
					smdx.PointElement{Id: InDCA, Offset: 5, Type: typelabel.Int16, ScaleFactor: "InDCA_SF", Units: "A", Mandatory: true, Label: "Amps", Description: "String Input Current"},
					smdx.PointElement{Id: InDCAhr, Offset: 6, Type: typelabel.Acc32, ScaleFactor: "InDCAhr_SF", Units: "Ah", Label: "Amp-hours", Description: "String Input Amp-Hours"},
				},
			},
		}})
}
