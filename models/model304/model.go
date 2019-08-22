// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model304

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block304 - Inclinometer Model - Include to support orientation measurements

const (
	ModelID          = 304
	ModelLabel       = "Inclinometer Model"
	ModelDescription = "Include to support orientation measurements"
)

const (
	Inclx = "Inclx"
	Incly = "Incly"
	Inclz = "Inclz"
)

type Block304Repeat struct {
	Inclx int32 `sunspec:"offset=0,sf=-2"`
	Incly int32 `sunspec:"offset=2,sf=-2"`
	Inclz int32 `sunspec:"offset=4,sf=-2"`
}

type Block304 struct {
	Repeats []Block304Repeat
}

func (self *Block304) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "inclinometer",
		Length: 6,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{Name: "incl",
				Length: 6,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: Inclx, Offset: 0, Type: typelabel.Int32, ScaleFactor: "-2", Units: "Degrees", Mandatory: true, Label: "X", Description: "X-Axis inclination"},
					smdx.PointElement{Id: Incly, Offset: 2, Type: typelabel.Int32, ScaleFactor: "-2", Units: "Degrees", Label: "Y", Description: "Y-Axis inclination"},
					smdx.PointElement{Id: Inclz, Offset: 4, Type: typelabel.Int32, ScaleFactor: "-2", Units: "Degrees", Label: "Z", Description: "Z-Axis inclination"},
				},
			},
		}})
}
