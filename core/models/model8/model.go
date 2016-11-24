// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model8

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block8 - Get Device Security Certificate - Security model for PKI

const (
	ModelID = 8
)

const (
	Cert = "Cert"
	Fmt  = "Fmt"
	N    = "N"
)

type Block8Repeat struct {
	Cert uint16 `sunspec:"offset=0,access=r"`
}

type Block8 struct {
	Fmt core.Enum16 `sunspec:"offset=0,access=r"`
	N   uint16      `sunspec:"offset=1,access=r"`

	Repeats []Block8Repeat
}

func (self *Block8) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 3,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 2,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Fmt, Offset: 0, Type: typelabel.Enum16, Access: "r", Mandatory: true},
					smdx.PointElement{Id: N, Offset: 1, Type: typelabel.Uint16, Access: "r", Mandatory: true},
				},
			},
			smdx.BlockElement{
				Length: 1,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: Cert, Offset: 0, Type: typelabel.Uint16, Access: "r", Mandatory: true},
				},
			},
		}})
}
