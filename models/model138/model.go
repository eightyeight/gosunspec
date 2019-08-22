// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model138

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block138 - HVRTC - HVRT must remain connected

const (
	ModelID          = 138
	ModelLabel       = "HVRTC"
	ModelDescription = "HVRT must remain connected"
)

const (
	ActCrv   = "ActCrv"
	ActPt    = "ActPt"
	CrvNam   = "CrvNam"
	ModEna   = "ModEna"
	NCrv     = "NCrv"
	NPt      = "NPt"
	Pad      = "Pad"
	ReadOnly = "ReadOnly"
	RmpTms   = "RmpTms"
	RvrtTms  = "RvrtTms"
	Tms1     = "Tms1"
	Tms10    = "Tms10"
	Tms11    = "Tms11"
	Tms12    = "Tms12"
	Tms13    = "Tms13"
	Tms14    = "Tms14"
	Tms15    = "Tms15"
	Tms16    = "Tms16"
	Tms17    = "Tms17"
	Tms18    = "Tms18"
	Tms19    = "Tms19"
	Tms2     = "Tms2"
	Tms20    = "Tms20"
	Tms3     = "Tms3"
	Tms4     = "Tms4"
	Tms5     = "Tms5"
	Tms6     = "Tms6"
	Tms7     = "Tms7"
	Tms8     = "Tms8"
	Tms9     = "Tms9"
	Tms_SF   = "Tms_SF"
	V1       = "V1"
	V10      = "V10"
	V11      = "V11"
	V12      = "V12"
	V13      = "V13"
	V14      = "V14"
	V15      = "V15"
	V16      = "V16"
	V17      = "V17"
	V18      = "V18"
	V19      = "V19"
	V2       = "V2"
	V20      = "V20"
	V3       = "V3"
	V4       = "V4"
	V5       = "V5"
	V6       = "V6"
	V7       = "V7"
	V8       = "V8"
	V9       = "V9"
	V_SF     = "V_SF"
	WinTms   = "WinTms"
)

type Block138Repeat struct {
	ActPt    uint16         `sunspec:"offset=0,len=1,access=rw"`
	Tms1     uint16         `sunspec:"offset=1,len=1,sf=Tms_SF,access=rw"`
	V1       uint16         `sunspec:"offset=2,len=1,sf=V_SF,access=rw"`
	Tms2     uint16         `sunspec:"offset=3,len=1,sf=Tms_SF,access=rw"`
	V2       uint16         `sunspec:"offset=4,len=1,sf=V_SF,access=rw"`
	Tms3     uint16         `sunspec:"offset=5,len=1,sf=Tms_SF,access=rw"`
	V3       uint16         `sunspec:"offset=6,len=1,sf=V_SF,access=rw"`
	Tms4     uint16         `sunspec:"offset=7,len=1,sf=Tms_SF,access=rw"`
	V4       uint16         `sunspec:"offset=8,len=1,sf=V_SF,access=rw"`
	Tms5     uint16         `sunspec:"offset=9,len=1,sf=Tms_SF,access=rw"`
	V5       uint16         `sunspec:"offset=10,len=1,sf=V_SF,access=rw"`
	Tms6     uint16         `sunspec:"offset=11,len=1,sf=Tms_SF,access=rw"`
	V6       uint16         `sunspec:"offset=12,len=1,sf=V_SF,access=rw"`
	Tms7     uint16         `sunspec:"offset=13,len=1,sf=Tms_SF,access=rw"`
	V7       uint16         `sunspec:"offset=14,len=1,sf=V_SF,access=rw"`
	Tms8     uint16         `sunspec:"offset=15,len=1,sf=Tms_SF,access=rw"`
	V8       uint16         `sunspec:"offset=16,len=1,sf=V_SF,access=rw"`
	Tms9     uint16         `sunspec:"offset=17,len=1,sf=Tms_SF,access=rw"`
	V9       uint16         `sunspec:"offset=18,len=1,sf=V_SF,access=rw"`
	Tms10    uint16         `sunspec:"offset=19,len=1,sf=Tms_SF,access=rw"`
	V10      uint16         `sunspec:"offset=20,len=1,sf=V_SF,access=rw"`
	Tms11    uint16         `sunspec:"offset=21,len=1,sf=Tms_SF,access=rw"`
	V11      uint16         `sunspec:"offset=22,len=1,sf=V_SF,access=rw"`
	Tms12    uint16         `sunspec:"offset=23,len=1,sf=Tms_SF,access=rw"`
	V12      uint16         `sunspec:"offset=24,len=1,sf=V_SF,access=rw"`
	Tms13    uint16         `sunspec:"offset=25,len=1,sf=Tms_SF,access=rw"`
	V13      uint16         `sunspec:"offset=26,len=1,sf=V_SF,access=rw"`
	Tms14    uint16         `sunspec:"offset=27,len=1,sf=Tms_SF,access=rw"`
	V14      uint16         `sunspec:"offset=28,len=1,sf=V_SF,access=rw"`
	Tms15    uint16         `sunspec:"offset=29,len=1,sf=Tms_SF,access=rw"`
	V15      uint16         `sunspec:"offset=30,len=1,sf=V_SF,access=rw"`
	Tms16    uint16         `sunspec:"offset=31,len=1,sf=Tms_SF,access=rw"`
	V16      uint16         `sunspec:"offset=32,len=1,sf=V_SF,access=rw"`
	Tms17    uint16         `sunspec:"offset=33,len=1,sf=Tms_SF,access=rw"`
	V17      uint16         `sunspec:"offset=34,len=1,sf=V_SF,access=rw"`
	Tms18    uint16         `sunspec:"offset=35,len=1,sf=Tms_SF,access=rw"`
	V18      uint16         `sunspec:"offset=36,len=1,sf=V_SF,access=rw"`
	Tms19    uint16         `sunspec:"offset=37,len=1,sf=Tms_SF,access=rw"`
	V19      uint16         `sunspec:"offset=38,len=1,sf=V_SF,access=rw"`
	Tms20    uint16         `sunspec:"offset=39,len=1,sf=Tms_SF,access=rw"`
	V20      uint16         `sunspec:"offset=40,len=1,sf=V_SF,access=rw"`
	CrvNam   string         `sunspec:"offset=41,len=8,access=rw"`
	ReadOnly sunspec.Enum16 `sunspec:"offset=49,len=1,access=r"`
}

type Block138 struct {
	ActCrv  uint16              `sunspec:"offset=0,len=1,access=rw"`
	ModEna  sunspec.Bitfield16  `sunspec:"offset=1,len=1,access=rw"`
	WinTms  uint16              `sunspec:"offset=2,len=1,access=rw"`
	RvrtTms uint16              `sunspec:"offset=3,len=1,access=rw"`
	RmpTms  uint16              `sunspec:"offset=4,len=1,access=rw"`
	NCrv    uint16              `sunspec:"offset=5,len=1,access=r"`
	NPt     uint16              `sunspec:"offset=6,len=1,access=r"`
	Tms_SF  sunspec.ScaleFactor `sunspec:"offset=7,len=1,access=r"`
	V_SF    sunspec.ScaleFactor `sunspec:"offset=8,len=1,access=r"`
	Pad     sunspec.Pad         `sunspec:"offset=9,len=1,access=r"`

	Repeats []Block138Repeat
}

func (self *Block138) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "hvrtc",
		Length: 60,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 10,
				Type:   "fixed",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActCrv, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActCrv", Description: "Index of active curve. 0=no active curve."},
					smdx.PointElement{Id: ModEna, Offset: 1, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true, Label: "ModEna", Description: "HVRT control mode. Enable active curve.  Bitfield value."},
					smdx.PointElement{Id: WinTms, Offset: 2, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "WinTms", Description: "Time window for HVRT change."},
					smdx.PointElement{Id: RvrtTms, Offset: 3, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RvrtTms", Description: "Timeout period for HVRT curve selection."},
					smdx.PointElement{Id: RmpTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1, Label: "RmpTms", Description: "Ramp time for moving from current mode to new mode."},
					smdx.PointElement{Id: NCrv, Offset: 5, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NCrv", Description: "Number of curves supported (recommend 4)."},
					smdx.PointElement{Id: NPt, Offset: 6, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true, Label: "NPt", Description: "Number of curve points supported (maximum of 20)."},
					smdx.PointElement{Id: Tms_SF, Offset: 7, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "Tms_SF", Description: "Scale factor for duration."},
					smdx.PointElement{Id: V_SF, Offset: 8, Type: typelabel.ScaleFactor, Access: "r", Length: 1, Mandatory: true, Label: "V_SF", Description: "Scale factor for percent VRef."},
					smdx.PointElement{Id: Pad, Offset: 9, Type: typelabel.Pad, Access: "r", Length: 1},
				},
			},
			smdx.BlockElement{Name: "curve",
				Length: 50,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActPt, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true, Label: "ActPt", Description: "Number of active points in array."},
					smdx.PointElement{Id: Tms1, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Mandatory: true, Label: "Tms1", Description: "Point 1 must remain connected duration."},
					smdx.PointElement{Id: V1, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Mandatory: true, Label: "V1", Description: "Point 1 must remain connected voltage."},
					smdx.PointElement{Id: Tms2, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms2", Description: "Point 2 must remain connected duration."},
					smdx.PointElement{Id: V2, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V2", Description: "Point 2 must remain connected voltage."},
					smdx.PointElement{Id: Tms3, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms3", Description: "Point 3 must remain connected duration."},
					smdx.PointElement{Id: V3, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V3", Description: "Point 3 must remain connected voltage."},
					smdx.PointElement{Id: Tms4, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms4", Description: "Point 4 must remain connected duration."},
					smdx.PointElement{Id: V4, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V4", Description: "Point 4 must remain connected voltage."},
					smdx.PointElement{Id: Tms5, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms5", Description: "Point 5 must remain connected duration."},
					smdx.PointElement{Id: V5, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V5", Description: "Point 5 must remain connected voltage."},
					smdx.PointElement{Id: Tms6, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms6", Description: "Point 6 must remain connected duration."},
					smdx.PointElement{Id: V6, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V6", Description: "Point 6 must remain connected voltage."},
					smdx.PointElement{Id: Tms7, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms7", Description: "Point 7 must remain connected duration."},
					smdx.PointElement{Id: V7, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V7", Description: "Point 7 must remain connected voltage."},
					smdx.PointElement{Id: Tms8, Offset: 15, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms8", Description: "Point 8 must remain connected duration."},
					smdx.PointElement{Id: V8, Offset: 16, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V8", Description: "Point 8 must remain connected voltage."},
					smdx.PointElement{Id: Tms9, Offset: 17, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms9", Description: "Point 9 must remain connected duration."},
					smdx.PointElement{Id: V9, Offset: 18, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V9", Description: "Point 9 must remain connected voltage."},
					smdx.PointElement{Id: Tms10, Offset: 19, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms10", Description: "Point 10 must remain connected duration."},
					smdx.PointElement{Id: V10, Offset: 20, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V10", Description: "Point 10 must remain connected voltage."},
					smdx.PointElement{Id: Tms11, Offset: 21, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms11", Description: "Point 11 must remain connected duration."},
					smdx.PointElement{Id: V11, Offset: 22, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V11", Description: "Point 11 must remain connected voltage."},
					smdx.PointElement{Id: Tms12, Offset: 23, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms12", Description: "Point 12 must remain connected duration."},
					smdx.PointElement{Id: V12, Offset: 24, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V12", Description: "Point 12 must remain connected voltage."},
					smdx.PointElement{Id: Tms13, Offset: 25, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms13", Description: "Point 13 must remain connected duration."},
					smdx.PointElement{Id: V13, Offset: 26, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V13", Description: "Point 13 must remain connected voltage."},
					smdx.PointElement{Id: Tms14, Offset: 27, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms14", Description: "Point 14 must remain connected duration."},
					smdx.PointElement{Id: V14, Offset: 28, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V14", Description: "Point 14 must remain connected voltage."},
					smdx.PointElement{Id: Tms15, Offset: 29, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms15", Description: "Point 15 must remain connected duration."},
					smdx.PointElement{Id: V15, Offset: 30, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V15", Description: "Point 15 must remain connected voltage."},
					smdx.PointElement{Id: Tms16, Offset: 31, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms16", Description: "Point 16 must remain connected duration."},
					smdx.PointElement{Id: V16, Offset: 32, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V16", Description: "Point 16 must remain connected voltage."},
					smdx.PointElement{Id: Tms17, Offset: 33, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms17", Description: "Point 17 must remain connected duration."},
					smdx.PointElement{Id: V17, Offset: 34, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V17", Description: "Point 17 must remain connected voltage."},
					smdx.PointElement{Id: Tms18, Offset: 35, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms18", Description: "Point 18 must remain connected duration."},
					smdx.PointElement{Id: V18, Offset: 36, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V18", Description: "Point 18 must remain connected voltage."},
					smdx.PointElement{Id: Tms19, Offset: 37, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms19", Description: "Point 19 must remain connected duration."},
					smdx.PointElement{Id: V19, Offset: 38, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V19", Description: "Point 19 must remain connected voltage."},
					smdx.PointElement{Id: Tms20, Offset: 39, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Label: "Tms20", Description: "Point 20 must remain connected duration."},
					smdx.PointElement{Id: V20, Offset: 40, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Label: "V20", Description: "Point 20 must remain connected voltage."},
					smdx.PointElement{Id: CrvNam, Offset: 41, Type: typelabel.String, Access: "rw", Length: 8, Label: "CrvNam", Description: "Optional description for curve."},
					smdx.PointElement{Id: ReadOnly, Offset: 49, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true, Label: "ReadOnly", Description: "Enumerated value indicates if curve is read-only or can be modified."},
				},
			},
		}})
}
