// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model806

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block806 - Flow Battery Model -

const (
	ModelID = 806
)

const (
	BatStTBD = "BatStTBD"
	BatTBD   = "BatTBD"
)

type Block806Repeat struct {
	BatStTBD uint16 `sunspec:"offset=0"`
}

type Block806 struct {
	BatTBD uint16 `sunspec:"offset=0"`

	Repeats []Block806Repeat
}

func (self *Block806) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "flow_battery",
		Length: 2,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 1,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: BatTBD, Offset: 0, Type: typelabel.Uint16, Mandatory: true},
				},
			},
			smdx.BlockElement{Name: "battery_string",
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: BatStTBD, Offset: 0, Type: typelabel.Uint16, Mandatory: true},
				},
			},
		}})
}