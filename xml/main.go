// xml implements data structures and functions for parsing, working with and
// exporting SunSpec data in XML exchange format. The XML schema in use is
// described in the SunSpec Data Exchange Specification version 1.2,
// particularly on page 11. Note that the parser in this package is case-sensitive
// even though some examples in the specification use `sunSpecData` instead of
// `SunSpecData` as the top-level tag name (this package assumes the latter).
// http://sunspec.org/wp-content/uploads/2015/06/SunSpec-Model-Data-Exchange-12021.pdf
package xml

import (
	"encoding/xml"
	"io"
	"time"
)

type Data struct {
	XMLName xml.Name `xml:"SunSpecData"`
	Version string   `xml:"v,attr"`
	Devices []Device `xml:"d"`
}

type Device struct {
	XMLName       xml.Name  `xml:"d"`
	CorrelationId uint32    `xml:"cid,attr,omitempty"`
	Id            string    `xml:"id,attr,omitempty"`
	Namespace     string    `xml:"ns,attr"`
	LoggerId      string    `xml:"lid,attr,omitempty"`
	Manufacturer  string    `xml:"man,attr"`
	Model         string    `xml:"mod,attr"`
	Serial        string    `xml:"sn,attr"`
	Timestamp     time.Time `xml:"t,attr"`
	Models        []Model   `xml:"m"`
}

// Note that we can use omitempty on Index because indices in SunSpec XML start
// at 1. Therefore an index of 0 (unused) will not be serialised.
type Model struct {
	XMLName   xml.Name `xml:"m"`
	Id        uint32   `xml:"id,attr"`
	Namespace string   `xml:"ns,attr,omitempty"`
	Index     uint32   `xml:"x,attr,omitempty"`
	Points    []Point  `xml:"p"`
}

// Note that we can use omitempty on ScaleFactor because a scale factor of 0
// means no scaling. Therefore an sf of 0 is meaningless anyway.
type Point struct {
	XMLName     xml.Name  `xml:"p"`
	Description string    `xml:"d,attr,omitempty"`
	Id          string    `xml:"id,attr"`
	ScaleFactor int16     `xml:"sf,attr,omitempty"`
	Timestamp   time.Time `xml:"t,attr,omitempty"`
	Unit        string    `xml:"u,attr,omitempty"`
	Index       uint32    `xml:"x,attr,omitempty"`
	Value       string    `xml:",chardata"`
}

func parseXML(reader io.Reader) (data []Data, err error) {
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(&data)
	return
}