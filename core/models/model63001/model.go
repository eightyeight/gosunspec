// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model63001

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block63001 - SunSpec Test Model 1 -

const (
	ModelID = 63001
)

const (
	Acc16        = "acc16"
	Acc16_u      = "acc16_u"
	Acc32        = "acc32"
	Acc32_u      = "acc32_u"
	Acc64        = "acc64"
	Acc64_u      = "acc64_u"
	Bitfield16   = "bitfield16"
	Bitfield16_u = "bitfield16_u"
	Bitfield32   = "bitfield32"
	Bitfield32_u = "bitfield32_u"
	Enum16       = "enum16"
	Enum16_u     = "enum16_u"
	Enum32       = "enum32"
	Enum32_u     = "enum32_u"
	Float32      = "float32"
	Float32_u    = "float32_u"
	Int16_1      = "int16_1"
	Int16_11     = "int16_11"
	Int16_12     = "int16_12"
	Int16_2      = "int16_2"
	Int16_3      = "int16_3"
	Int16_4      = "int16_4"
	Int16_5      = "int16_5"
	Int16_u      = "int16_u"
	Int32        = "int32"
	Int32_1      = "int32_1"
	Int32_2      = "int32_2"
	Int32_3      = "int32_3"
	Int32_4      = "int32_4"
	Int32_5      = "int32_5"
	Int32_u      = "int32_u"
	Int64        = "int64"
	Int64_u      = "int64_u"
	Ipaddr       = "ipaddr"
	Ipaddr_u     = "ipaddr_u"
	Ipv6addr     = "ipv6addr"
	Ipv6addr_u   = "ipv6addr_u"
	Pad_1        = "pad_1"
	Pad_2        = "pad_2"
	String       = "string"
	String_u     = "string_u"
	Sunssf_1     = "sunssf_1"
	Sunssf_2     = "sunssf_2"
	Sunssf_3     = "sunssf_3"
	Sunssf_4     = "sunssf_4"
	Sunssf_5     = "sunssf_5"
	Sunssf_6     = "sunssf_6"
	Sunssf_7     = "sunssf_7"
	Sunssf_8     = "sunssf_8"
	Sunssf_9     = "sunssf_9"
	Uint16_1     = "uint16_1"
	Uint16_11    = "uint16_11"
	Uint16_12    = "uint16_12"
	Uint16_13    = "uint16_13"
	Uint16_2     = "uint16_2"
	Uint16_3     = "uint16_3"
	Uint16_4     = "uint16_4"
	Uint16_5     = "uint16_5"
	Uint16_u     = "uint16_u"
	Uint32       = "uint32"
	Uint32_1     = "uint32_1"
	Uint32_2     = "uint32_2"
	Uint32_3     = "uint32_3"
	Uint32_4     = "uint32_4"
	Uint32_5     = "uint32_5"
	Uint32_u     = "uint32_u"
)

type Block63001Repeat struct {
	sunssf_8  core.ScaleFactor `sunspec:"offset=0"`
	int16_11  int16            `sunspec:"offset=1,sf=sunssf_8,access=rw"`
	int16_12  int16            `sunspec:"offset=2,sf=sunssf_9"`
	int16_u   int16            `sunspec:"offset=3"`
	uint16_11 uint16           `sunspec:"offset=4,sf=sunssf_8,access=rw"`
	uint16_12 uint16           `sunspec:"offset=5,sf=sunssf_9"`
	uint16_13 uint16           `sunspec:"offset=6"`
	uint16_u  uint16           `sunspec:"offset=7"`
	int32     int32            `sunspec:"offset=8,sf=sunssf_1,access=rw"`
	int32_u   int32            `sunspec:"offset=10"`
	uint32    uint32           `sunspec:"offset=12,sf=sunssf_9,access=rw"`
	uint32_u  uint32           `sunspec:"offset=14"`
	sunssf_9  core.ScaleFactor `sunspec:"offset=16"`
	pad_2     core.Pad         `sunspec:"offset=17"`
}

type Block63001 struct {
	sunssf_1     core.ScaleFactor `sunspec:"offset=0"`
	sunssf_2     core.ScaleFactor `sunspec:"offset=1"`
	sunssf_3     core.ScaleFactor `sunspec:"offset=2"`
	sunssf_4     core.ScaleFactor `sunspec:"offset=3"`
	int16_1      int16            `sunspec:"offset=4,sf=sunssf_1"`
	int16_2      int16            `sunspec:"offset=5,sf=sunssf_2"`
	int16_3      int16            `sunspec:"offset=6,sf=sunssf_3"`
	int16_4      int16            `sunspec:"offset=7,sf=sunssf_4,access=rw"`
	int16_5      int16            `sunspec:"offset=8"`
	int16_u      int16            `sunspec:"offset=9"`
	uint16_1     uint16           `sunspec:"offset=10,sf=sunssf_1"`
	uint16_2     uint16           `sunspec:"offset=11,sf=sunssf_2"`
	uint16_3     uint16           `sunspec:"offset=12,sf=sunssf_3"`
	uint16_4     uint16           `sunspec:"offset=13,sf=sunssf_4,access=rw"`
	uint16_5     uint16           `sunspec:"offset=14"`
	uint16_u     uint16           `sunspec:"offset=15"`
	acc16        core.Acc16       `sunspec:"offset=16"`
	acc16_u      core.Acc16       `sunspec:"offset=17"`
	enum16       core.Enum16      `sunspec:"offset=18"`
	enum16_u     core.Enum16      `sunspec:"offset=19"`
	bitfield16   core.Bitfield16  `sunspec:"offset=20"`
	bitfield16_u core.Bitfield16  `sunspec:"offset=21"`
	int32_1      int32            `sunspec:"offset=22,sf=sunssf_5"`
	int32_2      int32            `sunspec:"offset=24,sf=sunssf_6"`
	int32_3      int32            `sunspec:"offset=26,sf=sunssf_7,access=rw"`
	int32_4      int32            `sunspec:"offset=28"`
	int32_5      int32            `sunspec:"offset=30"`
	int32_u      int32            `sunspec:"offset=32"`
	uint32_1     uint32           `sunspec:"offset=34,sf=sunssf_5"`
	uint32_2     uint32           `sunspec:"offset=36,sf=sunssf_6"`
	uint32_3     uint32           `sunspec:"offset=38,sf=sunssf_7,access=rw"`
	uint32_4     uint32           `sunspec:"offset=40,sf=1"`
	uint32_5     uint32           `sunspec:"offset=42"`
	uint32_u     uint32           `sunspec:"offset=44"`
	acc32        core.Acc32       `sunspec:"offset=46"`
	acc32_u      core.Acc32       `sunspec:"offset=48"`
	enum32       core.Enum32      `sunspec:"offset=50"`
	enum32_u     core.Enum32      `sunspec:"offset=52"`
	bitfield32   core.Bitfield32  `sunspec:"offset=54"`
	bitfield32_u core.Bitfield32  `sunspec:"offset=56"`
	ipaddr       core.Ipaddr      `sunspec:"offset=58,access=rw"`
	ipaddr_u     core.Ipaddr      `sunspec:"offset=60"`
	int64        int64            `sunspec:"offset=62,access=rw"`
	int64_u      int64            `sunspec:"offset=66"`
	acc64        core.Acc64       `sunspec:"offset=70"`
	acc64_u      core.Acc64       `sunspec:"offset=74"`
	ipv6addr     core.Ipv6addr    `sunspec:"offset=78"`
	ipv6addr_u   core.Ipv6addr    `sunspec:"offset=86"`
	float32      float32          `sunspec:"offset=94,access=rw"`
	float32_u    float32          `sunspec:"offset=96"`
	string       core.String      `sunspec:"offset=98,len=16,access=rw"`
	string_u     core.String      `sunspec:"offset=114,len=16"`
	sunssf_5     core.ScaleFactor `sunspec:"offset=130"`
	sunssf_6     core.ScaleFactor `sunspec:"offset=131"`
	sunssf_7     core.ScaleFactor `sunspec:"offset=132"`
	pad_1        core.Pad         `sunspec:"offset=133"`

	Repeats []Block63001Repeat
}

func (self *Block63001) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 152,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 134,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Sunssf_1, Offset: 0, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Sunssf_2, Offset: 1, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Sunssf_3, Offset: 2, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Sunssf_4, Offset: 3, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Int16_1, Offset: 4, Type: typelabel.Int16, ScaleFactor: "sunssf_1"},
					smdx.PointElement{Id: Int16_2, Offset: 5, Type: typelabel.Int16, ScaleFactor: "sunssf_2"},
					smdx.PointElement{Id: Int16_3, Offset: 6, Type: typelabel.Int16, ScaleFactor: "sunssf_3"},
					smdx.PointElement{Id: Int16_4, Offset: 7, Type: typelabel.Int16, ScaleFactor: "sunssf_4", Access: "rw"},
					smdx.PointElement{Id: Int16_5, Offset: 8, Type: typelabel.Int16},
					smdx.PointElement{Id: Int16_u, Offset: 9, Type: typelabel.Int16},
					smdx.PointElement{Id: Uint16_1, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "sunssf_1"},
					smdx.PointElement{Id: Uint16_2, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "sunssf_2"},
					smdx.PointElement{Id: Uint16_3, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "sunssf_3"},
					smdx.PointElement{Id: Uint16_4, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "sunssf_4", Access: "rw"},
					smdx.PointElement{Id: Uint16_5, Offset: 14, Type: typelabel.Uint16},
					smdx.PointElement{Id: Uint16_u, Offset: 15, Type: typelabel.Uint16},
					smdx.PointElement{Id: Acc16, Offset: 16, Type: typelabel.Acc16},
					smdx.PointElement{Id: Acc16_u, Offset: 17, Type: typelabel.Acc16},
					smdx.PointElement{Id: Enum16, Offset: 18, Type: typelabel.Enum16},
					smdx.PointElement{Id: Enum16_u, Offset: 19, Type: typelabel.Enum16},
					smdx.PointElement{Id: Bitfield16, Offset: 20, Type: typelabel.Bitfield16},
					smdx.PointElement{Id: Bitfield16_u, Offset: 21, Type: typelabel.Bitfield16},
					smdx.PointElement{Id: Int32_1, Offset: 22, Type: typelabel.Int32, ScaleFactor: "sunssf_5"},
					smdx.PointElement{Id: Int32_2, Offset: 24, Type: typelabel.Int32, ScaleFactor: "sunssf_6"},
					smdx.PointElement{Id: Int32_3, Offset: 26, Type: typelabel.Int32, ScaleFactor: "sunssf_7", Access: "rw"},
					smdx.PointElement{Id: Int32_4, Offset: 28, Type: typelabel.Int32},
					smdx.PointElement{Id: Int32_5, Offset: 30, Type: typelabel.Int32},
					smdx.PointElement{Id: Int32_u, Offset: 32, Type: typelabel.Int32},
					smdx.PointElement{Id: Uint32_1, Offset: 34, Type: typelabel.Uint32, ScaleFactor: "sunssf_5"},
					smdx.PointElement{Id: Uint32_2, Offset: 36, Type: typelabel.Uint32, ScaleFactor: "sunssf_6"},
					smdx.PointElement{Id: Uint32_3, Offset: 38, Type: typelabel.Uint32, ScaleFactor: "sunssf_7", Access: "rw"},
					smdx.PointElement{Id: Uint32_4, Offset: 40, Type: typelabel.Uint32, ScaleFactor: "1"},
					smdx.PointElement{Id: Uint32_5, Offset: 42, Type: typelabel.Uint32},
					smdx.PointElement{Id: Uint32_u, Offset: 44, Type: typelabel.Uint32},
					smdx.PointElement{Id: Acc32, Offset: 46, Type: typelabel.Acc32},
					smdx.PointElement{Id: Acc32_u, Offset: 48, Type: typelabel.Acc32},
					smdx.PointElement{Id: Enum32, Offset: 50, Type: typelabel.Enum32},
					smdx.PointElement{Id: Enum32_u, Offset: 52, Type: typelabel.Enum32},
					smdx.PointElement{Id: Bitfield32, Offset: 54, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: Bitfield32_u, Offset: 56, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: Ipaddr, Offset: 58, Type: typelabel.Ipaddr, Access: "rw"},
					smdx.PointElement{Id: Ipaddr_u, Offset: 60, Type: typelabel.Ipaddr},
					smdx.PointElement{Id: Int64, Offset: 62, Type: typelabel.Int64, Access: "rw"},
					smdx.PointElement{Id: Int64_u, Offset: 66, Type: typelabel.Int64},
					smdx.PointElement{Id: Acc64, Offset: 70, Type: typelabel.Acc64},
					smdx.PointElement{Id: Acc64_u, Offset: 74, Type: typelabel.Acc64},
					smdx.PointElement{Id: Ipv6addr, Offset: 78, Type: typelabel.Ipv6addr},
					smdx.PointElement{Id: Ipv6addr_u, Offset: 86, Type: typelabel.Ipv6addr},
					smdx.PointElement{Id: Float32, Offset: 94, Type: typelabel.Float32, Access: "rw"},
					smdx.PointElement{Id: Float32_u, Offset: 96, Type: typelabel.Float32},
					smdx.PointElement{Id: String, Offset: 98, Type: typelabel.String, Access: "rw", Length: 16},
					smdx.PointElement{Id: String_u, Offset: 114, Type: typelabel.String, Length: 16},
					smdx.PointElement{Id: Sunssf_5, Offset: 130, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Sunssf_6, Offset: 131, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Sunssf_7, Offset: 132, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Pad_1, Offset: 133, Type: typelabel.Pad},
				},
			},
			smdx.BlockElement{
				Length: 18,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: Sunssf_8, Offset: 0, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Int16_11, Offset: 1, Type: typelabel.Int16, ScaleFactor: "sunssf_8", Access: "rw"},
					smdx.PointElement{Id: Int16_12, Offset: 2, Type: typelabel.Int16, ScaleFactor: "sunssf_9"},
					smdx.PointElement{Id: Int16_u, Offset: 3, Type: typelabel.Int16},
					smdx.PointElement{Id: Uint16_11, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "sunssf_8", Access: "rw"},
					smdx.PointElement{Id: Uint16_12, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "sunssf_9"},
					smdx.PointElement{Id: Uint16_13, Offset: 6, Type: typelabel.Uint16},
					smdx.PointElement{Id: Uint16_u, Offset: 7, Type: typelabel.Uint16},
					smdx.PointElement{Id: Int32, Offset: 8, Type: typelabel.Int32, ScaleFactor: "sunssf_1", Access: "rw"},
					smdx.PointElement{Id: Int32_u, Offset: 10, Type: typelabel.Int32},
					smdx.PointElement{Id: Uint32, Offset: 12, Type: typelabel.Uint32, ScaleFactor: "sunssf_9", Access: "rw"},
					smdx.PointElement{Id: Uint32_u, Offset: 14, Type: typelabel.Uint32},
					smdx.PointElement{Id: Sunssf_9, Offset: 16, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Pad_2, Offset: 17, Type: typelabel.Pad},
				},
			},
		}})
}
