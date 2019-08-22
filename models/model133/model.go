// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model133

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block133 - Basic Scheduling - Basic Scheduling

const (
	ModelID          = 133
	ModelLabel       = "Basic Scheduling"
	ModelDescription = "Basic Scheduling "
)

const (
	ActIndx = "ActIndx"
	ActPts  = "ActPts"
	ActSchd = "ActSchd"
	IntvTyp = "IntvTyp"
	ModEna  = "ModEna"
	NPts    = "NPts"
	NSchd   = "NSchd"
	Nam     = "Nam"
	Pad     = "Pad"
	RepPer  = "RepPer"
	RmpTms  = "RmpTms"
	StrTms  = "StrTms"
	WinTms  = "WinTms"
	X1      = "X1"
	X10     = "X10"
	X2      = "X2"
	X3      = "X3"
	X4      = "X4"
	X5      = "X5"
	X6      = "X6"
	X7      = "X7"
	X8      = "X8"
	X9      = "X9"
	XTyp    = "XTyp"
	X_SF    = "X_SF"
	Y1      = "Y1"
	Y10     = "Y10"
	Y2      = "Y2"
	Y3      = "Y3"
	Y4      = "Y4"
	Y5      = "Y5"
	Y6      = "Y6"
	Y7      = "Y7"
	Y8      = "Y8"
	Y9      = "Y9"
	YTyp    = "YTyp"
	Y_SF    = "Y_SF"
)

type Block133Repeat struct {
	ActPts  uint16              `sunspec:"offset=0,len=1,access=rw"`
	StrTms  uint32              `sunspec:"offset=1,len=2,access=rw"`
	RepPer  uint16              `sunspec:"offset=3,len=1,access=rw"`
	IntvTyp sunspec.Enum16      `sunspec:"offset=4,len=1,access=rw"`
	XTyp    sunspec.Enum16      `sunspec:"offset=5,len=1,access=rw"`
	X_SF    sunspec.ScaleFactor `sunspec:"offset=6,len=1,access=rw"`
	YTyp    sunspec.Enum16      `sunspec:"offset=7,len=1,access=rw"`
	Y_SF    sunspec.ScaleFactor `sunspec:"offset=8,len=1,access=rw"`
	X1      int32               `sunspec:"offset=9,len=2,sf=X_SF,access=rw"`
	Y1      int32               `sunspec:"offset=11,len=2,sf=Y_SF,access=rw"`
	X2      int32               `sunspec:"offset=13,len=2,sf=X_SF,access=rw"`
	Y2      int32               `sunspec:"offset=15,len=2,sf=Y_SF,access=rw"`
	X3      int32               `sunspec:"offset=17,len=2,sf=X_SF,access=rw"`
	Y3      int32               `sunspec:"offset=19,len=2,sf=Y_SF,access=rw"`
	X4      int32               `sunspec:"offset=21,len=2,sf=X_SF,access=rw"`
	Y4      int32               `sunspec:"offset=23,len=2,sf=Y_SF,access=rw"`
	X5      int32               `sunspec:"offset=25,len=2,sf=X_SF,access=rw"`
	Y5      int32               `sunspec:"offset=27,len=2,sf=Y_SF,access=rw"`
	X6      int32               `sunspec:"offset=29,len=2,sf=X_SF,access=rw"`
	Y6      int32               `sunspec:"offset=31,len=2,sf=Y_SF,access=rw"`
	X7      int32               `sunspec:"offset=33,len=2,sf=X_SF,access=rw"`
	Y7      int32               `sunspec:"offset=35,len=2,sf=Y_SF,access=rw"`
	X8      int32               `sunspec:"offset=37,len=2,sf=X_SF,access=rw"`
	Y8      int32               `sunspec:"offset=39,len=2,sf=Y_SF,access=rw"`
	X9      int32               `sunspec:"offset=41,len=2,sf=X_SF,access=rw"`
	Y9      int32               `sunspec:"offset=43,len=2,sf=Y_SF,access=rw"`
	X10     int32               `sunspec:"offset=45,len=2,sf=X_SF,access=rw"`
	Y10     int32               `sunspec:"offset=47,len=2,sf=Y_SF,access=rw"`
	Nam     string              `sunspec:"offset=49,len=8,access=rw"`
	WinTms  uint16              `sunspec:"offset=57,len=1,access=rw"`
	RmpTms  uint16              `sunspec:"offset=58,len=1,access=rw"`
	ActIndx uint16              `sunspec:"offset=59,len=1,access=r"`
}

type Block133 struct {
	ActSchd sunspec.Bitfield32 `sunspec:"offset=0,len=2,access=rw"`
	ModEna  sunspec.Bitfield16 `sunspec:"offset=2,len=1,access=rw"`
	NSchd   uint16             `sunspec:"offset=3,len=1,access=r"`
	NPts    uint16             `sunspec:"offset=4,len=1,access=r"`
	Pad     sunspec.Pad        `sunspec:"offset=5,len=1,access=r"`

	Repeats []Block133Repeat
}

func (self *Block133) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "schedule",
		Length: 66,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 6,
				Type:   "fixed",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActSchd, Offset: 0, Type: typelabel.Bitfield32, Access: "rw", Length: 2, Mandatory: true, Label: "ActSchd", Description: "Bitfield of active schedules"},
					smdx.PointElement{Id: ModEna, Offset: 2, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "ModEna", Description: "Is basic scheduling active."},
					smdx.PointElement{Id: NSchd, Offset: 3, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NSchd", Description: "Number of schedules supported (recommend min. 4, max 32)"},
					smdx.PointElement{Id: NPts, Offset: 4, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NPts", Description: "Number of schedule entries supported (maximum of 10)."},
					smdx.PointElement{Id: Pad, Offset: 5, Type: typelabel.Pad, Access: "r", Length: 1, Label: "Pad", Description: "Pad register."},
				},
			},
			smdx.BlockElement{
				Length: 60,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActPts, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActPts", Description: "Number of active entries in schedule."},
					smdx.PointElement{Id: StrTms, Offset: 1, Type: typelabel.Uint32, Units: "Secs", Access: "rw", Length: 2, Mandatory: true, Label: "StrTms", Description: "Schedule start in seconds since 2000 JAN 01 00:00:00 UTC."},
					smdx.PointElement{Id: RepPer, Offset: 3, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "RepPer", Description: "The repetition count for time-based schedules (0=repeat forever)"},
					smdx.PointElement{Id: IntvTyp, Offset: 4, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true, Label: "SchdTyp", Description: "The repetition frequency for time-based schedules: no repeat=0"},
					smdx.PointElement{Id: XTyp, Offset: 5, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true, Label: "XTyp", Description: "The meaning of the X-values in the array. "},
					smdx.PointElement{Id: X_SF, Offset: 6, Type: typelabel.ScaleFactor, Access: "rw", Length: 1, Mandatory: true, Label: "X_SF", Description: "Scale factor for schedule range values."},
					smdx.PointElement{Id: YTyp, Offset: 7, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true, Label: "YTyp", Description: "The meaning of the Y-values in the array."},
					smdx.PointElement{Id: Y_SF, Offset: 8, Type: typelabel.ScaleFactor, Access: "rw", Length: 1, Mandatory: true, Label: "Y_SF", Description: "Scale factor for schedule target values."},
					smdx.PointElement{Id: X1, Offset: 9, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Mandatory: true, Label: "X1", Description: "Entry 1 range."},
					smdx.PointElement{Id: Y1, Offset: 11, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Mandatory: true, Label: "Y1", Description: "Entry 1 target."},
					smdx.PointElement{Id: X2, Offset: 13, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X2", Description: "Entry 2 range."},
					smdx.PointElement{Id: Y2, Offset: 15, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y2", Description: "Entry 2 target."},
					smdx.PointElement{Id: X3, Offset: 17, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X3", Description: "Entry 3 range."},
					smdx.PointElement{Id: Y3, Offset: 19, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y3", Description: "Entry 3 target."},
					smdx.PointElement{Id: X4, Offset: 21, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X4", Description: "Entry 4 range."},
					smdx.PointElement{Id: Y4, Offset: 23, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y4", Description: "Entry 4 target."},
					smdx.PointElement{Id: X5, Offset: 25, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X5", Description: "Entry 15range."},
					smdx.PointElement{Id: Y5, Offset: 27, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y5", Description: "Entry 5 target."},
					smdx.PointElement{Id: X6, Offset: 29, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X6", Description: "Entry 6 range."},
					smdx.PointElement{Id: Y6, Offset: 31, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y6", Description: "Entry 6 target."},
					smdx.PointElement{Id: X7, Offset: 33, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X7", Description: "Entry 7 range."},
					smdx.PointElement{Id: Y7, Offset: 35, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y7", Description: "Entry 7 target."},
					smdx.PointElement{Id: X8, Offset: 37, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X8", Description: "Entry 8 range."},
					smdx.PointElement{Id: Y8, Offset: 39, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y8", Description: "Entry 8 target."},
					smdx.PointElement{Id: X9, Offset: 41, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X9", Description: "Entry 9 range."},
					smdx.PointElement{Id: Y9, Offset: 43, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y9", Description: "Entry 9 target."},
					smdx.PointElement{Id: X10, Offset: 45, Type: typelabel.Int32, ScaleFactor: "X_SF", Access: "rw", Length: 2, Label: "X10", Description: "Entry 10 range."},
					smdx.PointElement{Id: Y10, Offset: 47, Type: typelabel.Int32, ScaleFactor: "Y_SF", Access: "rw", Length: 2, Label: "Y10", Description: "Entry 10 target."},
					smdx.PointElement{Id: Nam, Offset: 49, Type: typelabel.String, Access: "rw", Length: 8, Label: "Nam", Description: "Optional description for schedule."},
					smdx.PointElement{Id: WinTms, Offset: 57, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "WinTms", Description: "Time window for schedule entry change."},
					smdx.PointElement{Id: RmpTms, Offset: 58, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RmpTms", Description: "Ramp time for moving from current target to new target."},
					smdx.PointElement{Id: ActIndx, Offset: 59, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "ActIndx", Description: "Index of active entry in the active schedule."},
				},
			},
		}})
}
