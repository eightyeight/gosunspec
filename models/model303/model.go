// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model303

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block303 - Back of Module Temperature Model - Include to support variable number of  back of module temperature measurements

const (
	ModelID          = 303
	ModelLabel       = "Back of Module Temperature Model"
	ModelDescription = "Include to support variable number of  back of module temperature measurements"
)

const (
	TmpBOM = "TmpBOM"
)

type Block303Repeat struct {
	TmpBOM int16 `sunspec:"offset=0,sf=-1"`
}

type Block303 struct {
	Repeats []Block303Repeat
}

func (self *Block303) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "bom_temp",
		Length: 1,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{Name: "temp",
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: TmpBOM, Offset: 0, Type: typelabel.Int16, ScaleFactor: "-1", Units: "C", Mandatory: true, Label: "Temp", Description: "Back of module temperature measurement"},
				},
			},
		}})
}
