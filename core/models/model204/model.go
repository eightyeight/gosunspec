// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model204

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block204 - delta-connect three phase (abc) meter -

const (
	ModelID = 204
)

const (
	A               = "A"
	A_SF            = "A_SF"
	AphA            = "AphA"
	AphB            = "AphB"
	AphC            = "AphC"
	Evt             = "Evt"
	Hz              = "Hz"
	Hz_SF           = "Hz_SF"
	PF              = "PF"
	PF_SF           = "PF_SF"
	PFphA           = "PFphA"
	PFphB           = "PFphB"
	PFphC           = "PFphC"
	PPV             = "PPV"
	PhV             = "PhV"
	PhVphA          = "PhVphA"
	PhVphAB         = "PhVphAB"
	PhVphB          = "PhVphB"
	PhVphBC         = "PhVphBC"
	PhVphC          = "PhVphC"
	PhVphCA         = "PhVphCA"
	TotVAhExp       = "TotVAhExp"
	TotVAhExpPhA    = "TotVAhExpPhA"
	TotVAhExpPhB    = "TotVAhExpPhB"
	TotVAhExpPhC    = "TotVAhExpPhC"
	TotVAhImp       = "TotVAhImp"
	TotVAhImpPhA    = "TotVAhImpPhA"
	TotVAhImpPhB    = "TotVAhImpPhB"
	TotVAhImpPhC    = "TotVAhImpPhC"
	TotVAh_SF       = "TotVAh_SF"
	TotVArhExpQ3    = "TotVArhExpQ3"
	TotVArhExpQ3PhA = "TotVArhExpQ3PhA"
	TotVArhExpQ3PhB = "TotVArhExpQ3PhB"
	TotVArhExpQ3PhC = "TotVArhExpQ3PhC"
	TotVArhExpQ4    = "TotVArhExpQ4"
	TotVArhExpQ4PhA = "TotVArhExpQ4PhA"
	TotVArhExpQ4PhB = "TotVArhExpQ4PhB"
	TotVArhExpQ4PhC = "TotVArhExpQ4PhC"
	TotVArhImpQ1    = "TotVArhImpQ1"
	TotVArhImpQ1PhA = "TotVArhImpQ1PhA"
	TotVArhImpQ1PhB = "TotVArhImpQ1PhB"
	TotVArhImpQ1PhC = "TotVArhImpQ1PhC"
	TotVArhImpQ2    = "TotVArhImpQ2"
	TotVArhImpQ2PhA = "TotVArhImpQ2PhA"
	TotVArhImpQ2PhB = "TotVArhImpQ2PhB"
	TotVArhImpQ2PhC = "TotVArhImpQ2PhC"
	TotVArh_SF      = "TotVArh_SF"
	TotWhExp        = "TotWhExp"
	TotWhExpPhA     = "TotWhExpPhA"
	TotWhExpPhB     = "TotWhExpPhB"
	TotWhExpPhC     = "TotWhExpPhC"
	TotWhImp        = "TotWhImp"
	TotWhImpPhA     = "TotWhImpPhA"
	TotWhImpPhB     = "TotWhImpPhB"
	TotWhImpPhC     = "TotWhImpPhC"
	TotWh_SF        = "TotWh_SF"
	VA              = "VA"
	VAR             = "VAR"
	VAR_SF          = "VAR_SF"
	VARphA          = "VARphA"
	VARphB          = "VARphB"
	VARphC          = "VARphC"
	VA_SF           = "VA_SF"
	VAphA           = "VAphA"
	VAphB           = "VAphB"
	VAphC           = "VAphC"
	V_SF            = "V_SF"
	W               = "W"
	W_SF            = "W_SF"
	WphA            = "WphA"
	WphB            = "WphB"
	WphC            = "WphC"
)

type Block204 struct {
	A               int16            `sunspec:"offset=0,sf=A_SF"`
	AphA            int16            `sunspec:"offset=1,sf=A_SF"`
	AphB            int16            `sunspec:"offset=2,sf=A_SF"`
	AphC            int16            `sunspec:"offset=3,sf=A_SF"`
	A_SF            core.ScaleFactor `sunspec:"offset=4"`
	PhV             int16            `sunspec:"offset=5,sf=V_SF"`
	PhVphA          int16            `sunspec:"offset=6,sf=V_SF"`
	PhVphB          int16            `sunspec:"offset=7,sf=V_SF"`
	PhVphC          int16            `sunspec:"offset=8,sf=V_SF"`
	PPV             int16            `sunspec:"offset=9,sf=V_SF"`
	PhVphAB         int16            `sunspec:"offset=10,sf=V_SF"`
	PhVphBC         int16            `sunspec:"offset=11,sf=V_SF"`
	PhVphCA         int16            `sunspec:"offset=12,sf=V_SF"`
	V_SF            core.ScaleFactor `sunspec:"offset=13"`
	Hz              int16            `sunspec:"offset=14,sf=Hz_SF"`
	Hz_SF           core.ScaleFactor `sunspec:"offset=15"`
	W               int16            `sunspec:"offset=16,sf=W_SF"`
	WphA            int16            `sunspec:"offset=17,sf=W_SF"`
	WphB            int16            `sunspec:"offset=18,sf=W_SF"`
	WphC            int16            `sunspec:"offset=19,sf=W_SF"`
	W_SF            core.ScaleFactor `sunspec:"offset=20"`
	VA              int16            `sunspec:"offset=21,sf=VA_SF"`
	VAphA           int16            `sunspec:"offset=22,sf=VA_SF"`
	VAphB           int16            `sunspec:"offset=23,sf=VA_SF"`
	VAphC           int16            `sunspec:"offset=24,sf=VA_SF"`
	VA_SF           core.ScaleFactor `sunspec:"offset=25"`
	VAR             int16            `sunspec:"offset=26,sf=VAR_SF"`
	VARphA          int16            `sunspec:"offset=27,sf=VAR_SF"`
	VARphB          int16            `sunspec:"offset=28,sf=VAR_SF"`
	VARphC          int16            `sunspec:"offset=29,sf=VAR_SF"`
	VAR_SF          core.ScaleFactor `sunspec:"offset=30"`
	PF              int16            `sunspec:"offset=31,sf=PF_SF"`
	PFphA           int16            `sunspec:"offset=32,sf=PF_SF"`
	PFphB           int16            `sunspec:"offset=33,sf=PF_SF"`
	PFphC           int16            `sunspec:"offset=34,sf=PF_SF"`
	PF_SF           core.ScaleFactor `sunspec:"offset=35"`
	TotWhExp        core.Acc32       `sunspec:"offset=36,sf=TotWh_SF"`
	TotWhExpPhA     core.Acc32       `sunspec:"offset=38,sf=TotWh_SF"`
	TotWhExpPhB     core.Acc32       `sunspec:"offset=40,sf=TotWh_SF"`
	TotWhExpPhC     core.Acc32       `sunspec:"offset=42,sf=TotWh_SF"`
	TotWhImp        core.Acc32       `sunspec:"offset=44,sf=TotWh_SF"`
	TotWhImpPhA     core.Acc32       `sunspec:"offset=46,sf=TotWh_SF"`
	TotWhImpPhB     core.Acc32       `sunspec:"offset=48,sf=TotWh_SF"`
	TotWhImpPhC     core.Acc32       `sunspec:"offset=50,sf=TotWh_SF"`
	TotWh_SF        core.ScaleFactor `sunspec:"offset=52"`
	TotVAhExp       core.Acc32       `sunspec:"offset=53,sf=TotVAh_SF"`
	TotVAhExpPhA    core.Acc32       `sunspec:"offset=55,sf=TotVAh_SF"`
	TotVAhExpPhB    core.Acc32       `sunspec:"offset=57,sf=TotVAh_SF"`
	TotVAhExpPhC    core.Acc32       `sunspec:"offset=59,sf=TotVAh_SF"`
	TotVAhImp       core.Acc32       `sunspec:"offset=61,sf=TotVAh_SF"`
	TotVAhImpPhA    core.Acc32       `sunspec:"offset=63,sf=TotVAh_SF"`
	TotVAhImpPhB    core.Acc32       `sunspec:"offset=65,sf=TotVAh_SF"`
	TotVAhImpPhC    core.Acc32       `sunspec:"offset=67,sf=TotVAh_SF"`
	TotVAh_SF       core.ScaleFactor `sunspec:"offset=69"`
	TotVArhImpQ1    core.Acc32       `sunspec:"offset=70,sf=TotVArh_SF"`
	TotVArhImpQ1PhA core.Acc32       `sunspec:"offset=72,sf=TotVArh_SF"`
	TotVArhImpQ1PhB core.Acc32       `sunspec:"offset=74,sf=TotVArh_SF"`
	TotVArhImpQ1PhC core.Acc32       `sunspec:"offset=76,sf=TotVArh_SF"`
	TotVArhImpQ2    core.Acc32       `sunspec:"offset=78,sf=TotVArh_SF"`
	TotVArhImpQ2PhA core.Acc32       `sunspec:"offset=80,sf=TotVArh_SF"`
	TotVArhImpQ2PhB core.Acc32       `sunspec:"offset=82,sf=TotVArh_SF"`
	TotVArhImpQ2PhC core.Acc32       `sunspec:"offset=84,sf=TotVArh_SF"`
	TotVArhExpQ3    core.Acc32       `sunspec:"offset=86,sf=TotVArh_SF"`
	TotVArhExpQ3PhA core.Acc32       `sunspec:"offset=88,sf=TotVArh_SF"`
	TotVArhExpQ3PhB core.Acc32       `sunspec:"offset=90,sf=TotVArh_SF"`
	TotVArhExpQ3PhC core.Acc32       `sunspec:"offset=92,sf=TotVArh_SF"`
	TotVArhExpQ4    core.Acc32       `sunspec:"offset=94,sf=TotVArh_SF"`
	TotVArhExpQ4PhA core.Acc32       `sunspec:"offset=96,sf=TotVArh_SF"`
	TotVArhExpQ4PhB core.Acc32       `sunspec:"offset=98,sf=TotVArh_SF"`
	TotVArhExpQ4PhC core.Acc32       `sunspec:"offset=100,sf=TotVArh_SF"`
	TotVArh_SF      core.ScaleFactor `sunspec:"offset=102"`
	Evt             core.Bitfield32  `sunspec:"offset=103"`
}

func (self *Block204) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "ac_meter",
		Length: 105,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 105,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: A, Offset: 0, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: AphA, Offset: 1, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: AphB, Offset: 2, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: AphC, Offset: 3, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: A_SF, Offset: 4, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: PhV, Offset: 5, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: PhVphA, Offset: 6, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: PhVphB, Offset: 7, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: PhVphC, Offset: 8, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: PPV, Offset: 9, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: PhVphAB, Offset: 10, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: PhVphBC, Offset: 11, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: PhVphCA, Offset: 12, Type: typelabel.Int16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: V_SF, Offset: 13, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: Hz, Offset: 14, Type: typelabel.Int16, ScaleFactor: "Hz_SF", Units: "Hz", Mandatory: true},
					smdx.PointElement{Id: Hz_SF, Offset: 15, Type: typelabel.Sunssf},
					smdx.PointElement{Id: W, Offset: 16, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W", Mandatory: true},
					smdx.PointElement{Id: WphA, Offset: 17, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W"},
					smdx.PointElement{Id: WphB, Offset: 18, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W"},
					smdx.PointElement{Id: WphC, Offset: 19, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W"},
					smdx.PointElement{Id: W_SF, Offset: 20, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: VA, Offset: 21, Type: typelabel.Int16, ScaleFactor: "VA_SF", Units: "VA"},
					smdx.PointElement{Id: VAphA, Offset: 22, Type: typelabel.Int16, ScaleFactor: "VA_SF", Units: "VA"},
					smdx.PointElement{Id: VAphB, Offset: 23, Type: typelabel.Int16, ScaleFactor: "VA_SF", Units: "VA"},
					smdx.PointElement{Id: VAphC, Offset: 24, Type: typelabel.Int16, ScaleFactor: "VA_SF", Units: "VA"},
					smdx.PointElement{Id: VA_SF, Offset: 25, Type: typelabel.Sunssf},
					smdx.PointElement{Id: VAR, Offset: 26, Type: typelabel.Int16, ScaleFactor: "VAR_SF", Units: "var"},
					smdx.PointElement{Id: VARphA, Offset: 27, Type: typelabel.Int16, ScaleFactor: "VAR_SF", Units: "var"},
					smdx.PointElement{Id: VARphB, Offset: 28, Type: typelabel.Int16, ScaleFactor: "VAR_SF", Units: "var"},
					smdx.PointElement{Id: VARphC, Offset: 29, Type: typelabel.Int16, ScaleFactor: "VAR_SF", Units: "var"},
					smdx.PointElement{Id: VAR_SF, Offset: 30, Type: typelabel.Sunssf},
					smdx.PointElement{Id: PF, Offset: 31, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "Pct"},
					smdx.PointElement{Id: PFphA, Offset: 32, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "Pct"},
					smdx.PointElement{Id: PFphB, Offset: 33, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "Pct"},
					smdx.PointElement{Id: PFphC, Offset: 34, Type: typelabel.Int16, ScaleFactor: "PF_SF", Units: "Pct"},
					smdx.PointElement{Id: PF_SF, Offset: 35, Type: typelabel.Sunssf},
					smdx.PointElement{Id: TotWhExp, Offset: 36, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh", Mandatory: true},
					smdx.PointElement{Id: TotWhExpPhA, Offset: 38, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh"},
					smdx.PointElement{Id: TotWhExpPhB, Offset: 40, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh"},
					smdx.PointElement{Id: TotWhExpPhC, Offset: 42, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh"},
					smdx.PointElement{Id: TotWhImp, Offset: 44, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh", Mandatory: true},
					smdx.PointElement{Id: TotWhImpPhA, Offset: 46, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh"},
					smdx.PointElement{Id: TotWhImpPhB, Offset: 48, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh"},
					smdx.PointElement{Id: TotWhImpPhC, Offset: 50, Type: typelabel.Acc32, ScaleFactor: "TotWh_SF", Units: "Wh"},
					smdx.PointElement{Id: TotWh_SF, Offset: 52, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: TotVAhExp, Offset: 53, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAhExpPhA, Offset: 55, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAhExpPhB, Offset: 57, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAhExpPhC, Offset: 59, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAhImp, Offset: 61, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAhImpPhA, Offset: 63, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAhImpPhB, Offset: 65, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAhImpPhC, Offset: 67, Type: typelabel.Acc32, ScaleFactor: "TotVAh_SF", Units: "VAh"},
					smdx.PointElement{Id: TotVAh_SF, Offset: 69, Type: typelabel.Sunssf},
					smdx.PointElement{Id: TotVArhImpQ1, Offset: 70, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhImpQ1PhA, Offset: 72, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhImpQ1PhB, Offset: 74, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhImpQ1PhC, Offset: 76, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhImpQ2, Offset: 78, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhImpQ2PhA, Offset: 80, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhImpQ2PhB, Offset: 82, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhImpQ2PhC, Offset: 84, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ3, Offset: 86, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ3PhA, Offset: 88, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ3PhB, Offset: 90, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ3PhC, Offset: 92, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ4, Offset: 94, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ4PhA, Offset: 96, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ4PhB, Offset: 98, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArhExpQ4PhC, Offset: 100, Type: typelabel.Acc32, ScaleFactor: "TotVArh_SF", Units: "varh"},
					smdx.PointElement{Id: TotVArh_SF, Offset: 102, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Evt, Offset: 103, Type: typelabel.Bitfield32, Mandatory: true},
				},
			},
		}})
}
