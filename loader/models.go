package loader

import (
	"encoding/xml"

	"github.com/linkedin/goavro/v2"
)

type MixedDataSlice struct {
	XMLName    xml.Name    `xml:"intDataSlice"`
	MixedDatas []MixedData `xml:"mixedData"`
}
type MixedData struct {
	Int1     *int64  `avro:"int1" xml:"int1"`
	Int2     *int64  `avro:"int2" xml:"int2"`
	Int3     *int64  `avro:"int3" xml:"int3"`
	Int4     *int64  `avro:"int4" xml:"int4"`
	Int5     *int64  `avro:"int5" xml:"int5"`
	Int6     *int64  `avro:"int6" xml:"int6"`
	Int7     *int64  `avro:"int7" xml:"int7"`
	Int8     *int64  `avro:"int8" xml:"int8"`
	Int9     *int64  `avro:"int9" xml:"int9"`
	Int10    *int64  `avro:"int10" xml:"int10"`
	String1  *string `avro:"string1" xml:"string1"`
	String2  *string `avro:"string2" xml:"string2"`
	String3  *string `avro:"string3" xml:"string3"`
	String4  *string `avro:"string4" xml:"string4"`
	String5  *string `avro:"string5" xml:"string5"`
	String6  *string `avro:"string6" xml:"string6"`
	String7  *string `avro:"string7" xml:"string7"`
	String8  *string `avro:"string8" xml:"string8"`
	String9  *string `avro:"string9" xml:"string9"`
	String10 *string `avro:"string10" xml:"string10"`
}

func (m MixedData) ToMap() map[string]interface{} {
	d := make(map[string]interface{}, 20)
	if m.Int1 != nil {
		d["int1"] = goavro.Union("long", *m.Int1)
	} else {
		d["int1"] = goavro.Union("null", nil)
	}
	if m.Int2 != nil {
		d["int2"] = goavro.Union("long", *m.Int2)
	} else {
		d["int2"] = goavro.Union("null", nil)
	}
	if m.Int3 != nil {
		d["int3"] = goavro.Union("long", *m.Int3)
	} else {
		d["int3"] = goavro.Union("null", nil)
	}
	if m.Int4 != nil {
		d["int4"] = goavro.Union("long", *m.Int4)
	} else {
		d["int4"] = goavro.Union("null", nil)
	}
	if m.Int5 != nil {
		d["int5"] = goavro.Union("long", *m.Int5)
	} else {
		d["int5"] = goavro.Union("null", nil)
	}
	if m.Int6 != nil {
		d["int6"] = goavro.Union("long", *m.Int6)
	} else {
		d["int6"] = goavro.Union("null", nil)
	}
	if m.Int7 != nil {
		d["int7"] = goavro.Union("long", *m.Int7)
	} else {
		d["int7"] = goavro.Union("null", nil)
	}
	if m.Int8 != nil {
		d["int8"] = goavro.Union("long", *m.Int8)
	} else {
		d["int8"] = goavro.Union("null", nil)
	}
	if m.Int9 != nil {
		d["int9"] = goavro.Union("long", *m.Int9)
	} else {
		d["int9"] = goavro.Union("null", nil)
	}
	if m.Int10 != nil {
		d["int10"] = goavro.Union("long", *m.Int10)
	} else {
		d["int10"] = goavro.Union("null", nil)
	}

	if m.String1 != nil {
		d["string1"] = goavro.Union("string", *m.String1)
	} else {
		d["string1"] = goavro.Union("null", nil)
	}
	if m.String2 != nil {
		d["string2"] = goavro.Union("string", *m.String2)
	} else {
		d["string2"] = goavro.Union("null", nil)
	}
	if m.String3 != nil {
		d["string3"] = goavro.Union("string", *m.String3)
	} else {
		d["string3"] = goavro.Union("null", nil)
	}
	if m.String4 != nil {
		d["string4"] = goavro.Union("string", *m.String4)
	} else {
		d["string4"] = goavro.Union("null", nil)
	}
	if m.String5 != nil {
		d["string5"] = goavro.Union("string", *m.String5)
	} else {
		d["string5"] = goavro.Union("null", nil)
	}
	if m.String6 != nil {
		d["string6"] = goavro.Union("string", *m.String6)
	} else {
		d["string6"] = goavro.Union("null", nil)
	}
	if m.String7 != nil {
		d["string7"] = goavro.Union("string", *m.String7)
	} else {
		d["string7"] = goavro.Union("null", nil)
	}
	if m.String8 != nil {
		d["string8"] = goavro.Union("string", *m.String8)
	} else {
		d["string8"] = goavro.Union("null", nil)
	}
	if m.String9 != nil {
		d["string9"] = goavro.Union("string", *m.String9)
	} else {
		d["string9"] = goavro.Union("null", nil)
	}
	if m.String10 != nil {
		d["string10"] = goavro.Union("string", *m.String10)
	} else {
		d["string10"] = goavro.Union("null", nil)
	}

	return d
}

type IntDataSlice struct {
	XMLName  xml.Name  `xml:"intDataSlice"`
	IntDatas []IntData `xml:"intData"`
}
type IntData struct {
	Int1  *int64 `avro:"int1" xml:"int1"`
	Int2  *int64 `avro:"int2" xml:"int2"`
	Int3  *int64 `avro:"int3" xml:"int3"`
	Int4  *int64 `avro:"int4" xml:"int4"`
	Int5  *int64 `avro:"int5" xml:"int5"`
	Int6  *int64 `avro:"int6" xml:"int6"`
	Int7  *int64 `avro:"int7" xml:"int7"`
	Int8  *int64 `avro:"int8" xml:"int8"`
	Int9  *int64 `avro:"int9" xml:"int9"`
	Int10 *int64 `avro:"int10" xml:"int10"`
	Int11 *int64 `avro:"int11" xml:"int11"`
	Int12 *int64 `avro:"int12" xml:"int12"`
	Int13 *int64 `avro:"int13" xml:"int13"`
	Int14 *int64 `avro:"int14" xml:"int14"`
	Int15 *int64 `avro:"int15" xml:"int15"`
	Int16 *int64 `avro:"int16" xml:"int16"`
	Int17 *int64 `avro:"int17" xml:"int17"`
	Int18 *int64 `avro:"int18" xml:"int18"`
	Int19 *int64 `avro:"int19" xml:"int19"`
	Int20 *int64 `avro:"int20" xml:"int20"`
}

func (m IntData) ToMap() map[string]interface{} {
	d := make(map[string]interface{}, 20)
	if m.Int1 != nil {
		d["int1"] = goavro.Union("long", *m.Int1)
	} else {
		d["int1"] = goavro.Union("null", nil)
	}
	if m.Int2 != nil {
		d["int2"] = goavro.Union("long", *m.Int2)
	} else {
		d["int2"] = goavro.Union("null", nil)
	}
	if m.Int3 != nil {
		d["int3"] = goavro.Union("long", *m.Int3)
	} else {
		d["int3"] = goavro.Union("null", nil)
	}
	if m.Int4 != nil {
		d["int4"] = goavro.Union("long", *m.Int4)
	} else {
		d["int4"] = goavro.Union("null", nil)
	}
	if m.Int5 != nil {
		d["int5"] = goavro.Union("long", *m.Int5)
	} else {
		d["int5"] = goavro.Union("null", nil)
	}
	if m.Int6 != nil {
		d["int6"] = goavro.Union("long", *m.Int6)
	} else {
		d["int6"] = goavro.Union("null", nil)
	}
	if m.Int7 != nil {
		d["int7"] = goavro.Union("long", *m.Int7)
	} else {
		d["int7"] = goavro.Union("null", nil)
	}
	if m.Int8 != nil {
		d["int8"] = goavro.Union("long", *m.Int8)
	} else {
		d["int8"] = goavro.Union("null", nil)
	}
	if m.Int9 != nil {
		d["int9"] = goavro.Union("long", *m.Int9)
	} else {
		d["int9"] = goavro.Union("null", nil)
	}
	if m.Int10 != nil {
		d["int10"] = goavro.Union("long", *m.Int10)
	} else {
		d["int10"] = goavro.Union("null", nil)
	}
	if m.Int11 != nil {
		d["int11"] = goavro.Union("long", *m.Int11)
	} else {
		d["int11"] = goavro.Union("null", nil)
	}
	if m.Int12 != nil {
		d["int12"] = goavro.Union("long", *m.Int12)
	} else {
		d["int12"] = goavro.Union("null", nil)
	}
	if m.Int13 != nil {
		d["int13"] = goavro.Union("long", *m.Int13)
	} else {
		d["int13"] = goavro.Union("null", nil)
	}
	if m.Int14 != nil {
		d["int14"] = goavro.Union("long", *m.Int14)
	} else {
		d["int14"] = goavro.Union("null", nil)
	}
	if m.Int15 != nil {
		d["int15"] = goavro.Union("long", *m.Int15)
	} else {
		d["int15"] = goavro.Union("null", nil)
	}
	if m.Int16 != nil {
		d["int16"] = goavro.Union("long", *m.Int16)
	} else {
		d["int16"] = goavro.Union("null", nil)
	}
	if m.Int17 != nil {
		d["int17"] = goavro.Union("long", *m.Int17)
	} else {
		d["int17"] = goavro.Union("null", nil)
	}
	if m.Int18 != nil {
		d["int18"] = goavro.Union("long", *m.Int18)
	} else {
		d["int18"] = goavro.Union("null", nil)
	}
	if m.Int19 != nil {
		d["int19"] = goavro.Union("long", *m.Int19)
	} else {
		d["int19"] = goavro.Union("null", nil)
	}
	if m.Int20 != nil {
		d["int20"] = goavro.Union("long", *m.Int20)
	} else {
		d["int20"] = goavro.Union("null", nil)
	}

	return d
}

type StringDataSlice struct {
	XMLName     xml.Name     `xml:"intDataSlice"`
	StringDatas []StringData `xml:"stringData"`
}
type StringData struct {
	String1  *string `avro:"string1" xml:"string1"`
	String2  *string `avro:"string2" xml:"string2"`
	String3  *string `avro:"string3" xml:"string3"`
	String4  *string `avro:"string4" xml:"string4"`
	String5  *string `avro:"string5" xml:"string5"`
	String6  *string `avro:"string6" xml:"string6"`
	String7  *string `avro:"string7" xml:"string7"`
	String8  *string `avro:"string8" xml:"string8"`
	String9  *string `avro:"string9" xml:"string9"`
	String10 *string `avro:"string10" xml:"string10"`
	String11 *string `avro:"string11" xml:"string11"`
	String12 *string `avro:"string12" xml:"string12"`
	String13 *string `avro:"string13" xml:"string13"`
	String14 *string `avro:"string14" xml:"string14"`
	String15 *string `avro:"string15" xml:"string15"`
	String16 *string `avro:"string16" xml:"string16"`
	String17 *string `avro:"string17" xml:"string17"`
	String18 *string `avro:"string18" xml:"string18"`
	String19 *string `avro:"string19" xml:"string19"`
	String20 *string `avro:"string20" xml:"string20"`
}

func (m StringData) ToMap() map[string]interface{} {
	d := make(map[string]interface{}, 20)
	if m.String1 != nil {
		d["string1"] = goavro.Union("string", *m.String1)
	} else {
		d["string1"] = goavro.Union("null", nil)
	}
	if m.String2 != nil {
		d["string2"] = goavro.Union("string", *m.String2)
	} else {
		d["string2"] = goavro.Union("null", nil)
	}
	if m.String3 != nil {
		d["string3"] = goavro.Union("string", *m.String3)
	} else {
		d["string3"] = goavro.Union("null", nil)
	}
	if m.String4 != nil {
		d["string4"] = goavro.Union("string", *m.String4)
	} else {
		d["string4"] = goavro.Union("null", nil)
	}
	if m.String5 != nil {
		d["string5"] = goavro.Union("string", *m.String5)
	} else {
		d["string5"] = goavro.Union("null", nil)
	}
	if m.String6 != nil {
		d["string6"] = goavro.Union("string", *m.String6)
	} else {
		d["string6"] = goavro.Union("null", nil)
	}
	if m.String7 != nil {
		d["string7"] = goavro.Union("string", *m.String7)
	} else {
		d["string7"] = goavro.Union("null", nil)
	}
	if m.String8 != nil {
		d["string8"] = goavro.Union("string", *m.String8)
	} else {
		d["string8"] = goavro.Union("null", nil)
	}
	if m.String9 != nil {
		d["string9"] = goavro.Union("string", *m.String9)
	} else {
		d["string9"] = goavro.Union("null", nil)
	}
	if m.String10 != nil {
		d["string10"] = goavro.Union("string", *m.String10)
	} else {
		d["string10"] = goavro.Union("null", nil)
	}
	if m.String11 != nil {
		d["string11"] = goavro.Union("string", *m.String11)
	} else {
		d["string11"] = goavro.Union("null", nil)
	}
	if m.String12 != nil {
		d["string12"] = goavro.Union("string", *m.String12)
	} else {
		d["string12"] = goavro.Union("null", nil)
	}
	if m.String13 != nil {
		d["string13"] = goavro.Union("string", *m.String13)
	} else {
		d["string13"] = goavro.Union("null", nil)
	}
	if m.String14 != nil {
		d["string14"] = goavro.Union("string", *m.String14)
	} else {
		d["string14"] = goavro.Union("null", nil)
	}
	if m.String15 != nil {
		d["string15"] = goavro.Union("string", *m.String15)
	} else {
		d["string15"] = goavro.Union("null", nil)
	}
	if m.String16 != nil {
		d["string16"] = goavro.Union("string", *m.String16)
	} else {
		d["string16"] = goavro.Union("null", nil)
	}
	if m.String17 != nil {
		d["string17"] = goavro.Union("string", *m.String17)
	} else {
		d["string17"] = goavro.Union("null", nil)
	}
	if m.String18 != nil {
		d["string18"] = goavro.Union("string", *m.String18)
	} else {
		d["string18"] = goavro.Union("null", nil)
	}
	if m.String19 != nil {
		d["string19"] = goavro.Union("string", *m.String19)
	} else {
		d["string19"] = goavro.Union("null", nil)
	}
	if m.String20 != nil {
		d["string20"] = goavro.Union("string", *m.String20)
	} else {
		d["string20"] = goavro.Union("null", nil)
	}

	return d
}
