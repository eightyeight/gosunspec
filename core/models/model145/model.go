// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model145

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block145 - Extended Settings - Inverter controls extended settings

const (
	ModelID = 145
)

const (
	AGra         = "AGra"
	ConnRmpDnRte = "ConnRmpDnRte"
	ConnRmpUpRte = "ConnRmpUpRte"
	EmgRmpDnRte  = "EmgRmpDnRte"
	EmgRmpUpRte  = "EmgRmpUpRte"
	NomRmpDnRte  = "NomRmpDnRte"
	NomRmpUpRte  = "NomRmpUpRte"
	Rmp_SF       = "Rmp_SF"
)

type Block145 struct {
	NomRmpUpRte  uint16           `sunspec:"offset=0,len=1,sf=Rmp_SF,access=rw"`
	NomRmpDnRte  uint16           `sunspec:"offset=1,len=1,sf=Rmp_SF,access=rw"`
	EmgRmpUpRte  uint16           `sunspec:"offset=2,len=1,sf=Rmp_SF,access=rw"`
	EmgRmpDnRte  uint16           `sunspec:"offset=3,len=1,sf=Rmp_SF,access=rw"`
	ConnRmpUpRte uint16           `sunspec:"offset=4,len=1,sf=Rmp_SF,access=rw"`
	ConnRmpDnRte uint16           `sunspec:"offset=5,len=1,sf=Rmp_SF,access=rw"`
	AGra         uint16           `sunspec:"offset=6,len=1,sf=Rmp_SF,access=rw"`
	Rmp_SF       core.ScaleFactor `sunspec:"offset=7,len=1,access=r"`
}

func (self *Block145) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "ext_settings",
		Length: 8,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 8,
				Type:   "fixed",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: NomRmpUpRte, Offset: 0, Type: typelabel.Uint16, ScaleFactor: "Rmp_SF", Units: "Pct", Access: "rw", Length: 1},
					smdx.PointElement{Id: NomRmpDnRte, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "Rmp_SF", Units: "Pct", Access: "rw", Length: 1},
					smdx.PointElement{Id: EmgRmpUpRte, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "Rmp_SF", Units: "Pct", Access: "rw", Length: 1},
					smdx.PointElement{Id: EmgRmpDnRte, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "Rmp_SF", Units: "Pct", Access: "rw", Length: 1},
					smdx.PointElement{Id: ConnRmpUpRte, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "Rmp_SF", Units: "Pct", Access: "rw", Length: 1},
					smdx.PointElement{Id: ConnRmpDnRte, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "Rmp_SF", Units: "Pct", Access: "rw", Length: 1},
					smdx.PointElement{Id: AGra, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "Rmp_SF", Units: "Pct", Access: "rw", Length: 1},
					smdx.PointElement{Id: Rmp_SF, Offset: 7, Type: typelabel.Sunssf, Access: "r", Length: 1},
				},
			},
		}})
}
