// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model501

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block501 - Solar Module - A solar module model supporing DC-DC converter

const (
	ModelID = 501
)

const (
	Ctl      = "Ctl"
	CtlVal   = "CtlVal"
	CtlVend  = "CtlVend"
	Evt      = "Evt"
	EvtVend  = "EvtVend"
	InA      = "InA"
	InV      = "InV"
	InW      = "InW"
	InWh     = "InWh"
	OutA     = "OutA"
	OutV     = "OutV"
	OutW     = "OutW"
	OutWh    = "OutWh"
	Stat     = "Stat"
	StatVend = "StatVend"
	Tmp      = "Tmp"
	Tms      = "Tms"
)

type Block501 struct {
	Stat     core.Enum16     `sunspec:"offset=0"`
	StatVend core.Enum16     `sunspec:"offset=1"`
	Evt      core.Bitfield32 `sunspec:"offset=2"`
	EvtVend  core.Bitfield32 `sunspec:"offset=4"`
	Ctl      core.Enum16     `sunspec:"offset=6,access=rw"`
	CtlVend  core.Enum32     `sunspec:"offset=7,access=rw"`
	CtlVal   int32           `sunspec:"offset=9,access=rw"`
	Tms      uint32          `sunspec:"offset=11"`
	OutA     float32         `sunspec:"offset=13"`
	OutV     float32         `sunspec:"offset=15"`
	OutWh    float32         `sunspec:"offset=17"`
	OutW     float32         `sunspec:"offset=19"`
	Tmp      float32         `sunspec:"offset=21"`
	InA      float32         `sunspec:"offset=23"`
	InV      float32         `sunspec:"offset=25"`
	InWh     float32         `sunspec:"offset=27"`
	InW      float32         `sunspec:"offset=29"`
}

func (self *Block501) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "solar_module",
		Length: 31,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 31,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Stat, Offset: 0, Type: typelabel.Enum16, Mandatory: true},
					smdx.PointElement{Id: StatVend, Offset: 1, Type: typelabel.Enum16},
					smdx.PointElement{Id: Evt, Offset: 2, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: EvtVend, Offset: 4, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: Ctl, Offset: 6, Type: typelabel.Enum16, Access: "rw"},
					smdx.PointElement{Id: CtlVend, Offset: 7, Type: typelabel.Enum32, Access: "rw"},
					smdx.PointElement{Id: CtlVal, Offset: 9, Type: typelabel.Int32, Access: "rw"},
					smdx.PointElement{Id: Tms, Offset: 11, Type: typelabel.Uint32, Units: "Secs"},
					smdx.PointElement{Id: OutA, Offset: 13, Type: typelabel.Float32, Units: "A"},
					smdx.PointElement{Id: OutV, Offset: 15, Type: typelabel.Float32, Units: "V"},
					smdx.PointElement{Id: OutWh, Offset: 17, Type: typelabel.Float32, Units: "Wh"},
					smdx.PointElement{Id: OutW, Offset: 19, Type: typelabel.Float32, Units: "W"},
					smdx.PointElement{Id: Tmp, Offset: 21, Type: typelabel.Float32, Units: "C"},
					smdx.PointElement{Id: InA, Offset: 23, Type: typelabel.Float32, Units: "A"},
					smdx.PointElement{Id: InV, Offset: 25, Type: typelabel.Float32, Units: "V"},
					smdx.PointElement{Id: InWh, Offset: 27, Type: typelabel.Float32, Units: "Wh"},
					smdx.PointElement{Id: InW, Offset: 29, Type: typelabel.Float32, Units: "W"},
				},
			},
		}})
}
