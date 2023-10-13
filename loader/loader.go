package loader

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/linkedin/goavro/v2"

	protoSchema "goFormatsBenchmarking/protobuf/schema"
)

func LoadMixedData(filename string, rowsToRead int64) []MixedData {
	// Open the CSV file
	if filename[len(filename)-4:] != ".csv" {
		filename += ".csv"
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error opening file: %w", err))
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV data
	var line []string
	line, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	// Parse the CSV data into a Go structure
	data := make([]MixedData, 0, rowsToRead)
	defer func() {
		fmt.Println(fmt.Sprintf("data cap=%d, data len=%d", cap(data), len(data)))
	}()
	for line != nil && rowsToRead > 0 {
		rowsToRead -= 1
		anotherData := MixedData{}
		if line[0] != "" {
			val, _ := strconv.ParseInt(line[0], 10, 64)
			anotherData.Int1 = &val
		}
		if line[1] != "" {
			val, _ := strconv.ParseInt(line[1], 10, 64)
			anotherData.Int2 = &val
		}
		if line[2] != "" {
			val, _ := strconv.ParseInt(line[2], 10, 64)
			anotherData.Int3 = &val
		}
		if line[3] != "" {
			val, _ := strconv.ParseInt(line[3], 10, 64)
			anotherData.Int4 = &val
		}
		if line[4] != "" {
			val, _ := strconv.ParseInt(line[4], 10, 64)
			anotherData.Int5 = &val
		}
		if line[5] != "" {
			val, _ := strconv.ParseInt(line[5], 10, 64)
			anotherData.Int6 = &val
		}
		if line[6] != "" {
			val, _ := strconv.ParseInt(line[6], 10, 64)
			anotherData.Int7 = &val
		}
		if line[7] != "" {
			val, _ := strconv.ParseInt(line[7], 10, 64)
			anotherData.Int8 = &val
		}
		if line[8] != "" {
			val, _ := strconv.ParseInt(line[8], 10, 64)
			anotherData.Int9 = &val
		}
		if line[9] != "" {
			val, _ := strconv.ParseInt(line[9], 10, 64)
			anotherData.Int10 = &val
		}

		if line[10] != "" {
			anotherData.String1 = &line[10]
		}
		if line[11] != "" {
			anotherData.String2 = &line[11]
		}
		if line[12] != "" {
			anotherData.String3 = &line[12]
		}
		if line[13] != "" {
			anotherData.String4 = &line[13]
		}
		if line[14] != "" {
			anotherData.String5 = &line[14]
		}
		if line[15] != "" {
			anotherData.String6 = &line[15]
		}
		if line[16] != "" {
			anotherData.String7 = &line[16]
		}
		if line[17] != "" {
			anotherData.String8 = &line[17]
		}
		if line[18] != "" {
			anotherData.String9 = &line[18]
		}
		if line[19] != "" {
			anotherData.String10 = &line[19]
		}

		data = append(data, anotherData)

		line, err = reader.Read()
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return data
		}
	}

	return data
}

func LoadIntData(filename string, rowsToRead int64) []IntData {
	// Open the CSV file
	if filename[len(filename)-4:] != ".csv" {
		filename += ".csv"
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error opening file: %w", err))
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV data
	var line []string
	line, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	// Parse the CSV data into a Go structure
	data := make([]IntData, 0, rowsToRead)
	defer func() {
		fmt.Println(fmt.Sprintf("data cap=%d, data len=%d", cap(data), len(data)))
	}()
	for line != nil && rowsToRead > 0 {
		rowsToRead -= 1
		anotherData := IntData{}
		if line[0] != "" {
			val, _ := strconv.ParseInt(line[0], 10, 64)
			anotherData.Int1 = &val
		}
		if line[1] != "" {
			val, _ := strconv.ParseInt(line[1], 10, 64)
			anotherData.Int2 = &val
		}
		if line[2] != "" {
			val, _ := strconv.ParseInt(line[2], 10, 64)
			anotherData.Int3 = &val
		}
		if line[3] != "" {
			val, _ := strconv.ParseInt(line[3], 10, 64)
			anotherData.Int4 = &val
		}
		if line[4] != "" {
			val, _ := strconv.ParseInt(line[4], 10, 64)
			anotherData.Int5 = &val
		}
		if line[5] != "" {
			val, _ := strconv.ParseInt(line[5], 10, 64)
			anotherData.Int6 = &val
		}
		if line[6] != "" {
			val, _ := strconv.ParseInt(line[6], 10, 64)
			anotherData.Int7 = &val
		}
		if line[7] != "" {
			val, _ := strconv.ParseInt(line[7], 10, 64)
			anotherData.Int8 = &val
		}
		if line[8] != "" {
			val, _ := strconv.ParseInt(line[8], 10, 64)
			anotherData.Int9 = &val
		}
		if line[9] != "" {
			val, _ := strconv.ParseInt(line[9], 10, 64)
			anotherData.Int10 = &val
		}
		if line[10] != "" {
			val, _ := strconv.ParseInt(line[10], 10, 64)
			anotherData.Int11 = &val
		}
		if line[11] != "" {
			val, _ := strconv.ParseInt(line[11], 10, 64)
			anotherData.Int12 = &val
		}
		if line[12] != "" {
			val, _ := strconv.ParseInt(line[12], 10, 64)
			anotherData.Int13 = &val
		}
		if line[13] != "" {
			val, _ := strconv.ParseInt(line[13], 10, 64)
			anotherData.Int14 = &val
		}
		if line[14] != "" {
			val, _ := strconv.ParseInt(line[14], 10, 64)
			anotherData.Int15 = &val
		}
		if line[15] != "" {
			val, _ := strconv.ParseInt(line[15], 10, 64)
			anotherData.Int16 = &val
		}
		if line[16] != "" {
			val, _ := strconv.ParseInt(line[16], 10, 64)
			anotherData.Int17 = &val
		}
		if line[17] != "" {
			val, _ := strconv.ParseInt(line[17], 10, 64)
			anotherData.Int18 = &val
		}
		if line[18] != "" {
			val, _ := strconv.ParseInt(line[18], 10, 64)
			anotherData.Int19 = &val
		}
		if line[19] != "" {
			val, _ := strconv.ParseInt(line[19], 10, 64)
			anotherData.Int20 = &val
		}

		data = append(data, anotherData)

		line, err = reader.Read()
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return data
		}
	}

	return data
}

func LoadStringsData(filename string, rowsToRead int64) []StringData {
	// Open the CSV file
	if filename[len(filename)-4:] != ".csv" {
		filename += ".csv"
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("error opening file: %w", err))
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the CSV data
	var line []string
	line, err = reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	// Parse the CSV data into a Go structure
	data := make([]StringData, 0, rowsToRead)
	defer func() {
		fmt.Println(fmt.Sprintf("data cap=%d, data len=%d", cap(data), len(data)))
	}()
	for line != nil && rowsToRead > 0 {
		rowsToRead -= 1
		anotherData := StringData{}

		if line[0] != "" {
			anotherData.String1 = &line[0]
		}
		if line[1] != "" {
			anotherData.String2 = &line[1]
		}
		if line[2] != "" {
			anotherData.String3 = &line[2]
		}
		if line[3] != "" {
			anotherData.String4 = &line[3]
		}
		if line[4] != "" {
			anotherData.String5 = &line[4]
		}
		if line[5] != "" {
			anotherData.String6 = &line[5]
		}
		if line[6] != "" {
			anotherData.String7 = &line[6]
		}
		if line[7] != "" {
			anotherData.String8 = &line[7]
		}
		if line[8] != "" {
			anotherData.String9 = &line[8]
		}
		if line[9] != "" {
			anotherData.String10 = &line[9]
		}
		if line[10] != "" {
			anotherData.String11 = &line[10]
		}
		if line[11] != "" {
			anotherData.String12 = &line[11]
		}
		if line[12] != "" {
			anotherData.String13 = &line[12]
		}
		if line[13] != "" {
			anotherData.String14 = &line[13]
		}
		if line[14] != "" {
			anotherData.String15 = &line[14]
		}
		if line[15] != "" {
			anotherData.String16 = &line[15]
		}
		if line[16] != "" {
			anotherData.String17 = &line[16]
		}
		if line[17] != "" {
			anotherData.String18 = &line[17]
		}
		if line[18] != "" {
			anotherData.String19 = &line[18]
		}
		if line[19] != "" {
			anotherData.String20 = &line[19]
		}

		data = append(data, anotherData)

		line, err = reader.Read()
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return data
		}
	}

	return data
}

func LoadMixedProtoData(filename string, rowsToRead int64) *protoSchema.MixedDataSlice {
	dataSlice := LoadMixedData(filename, rowsToRead)

	protoData := make([]*protoSchema.MixedData, len(dataSlice))

	for i, data := range dataSlice {
		protoData[i] = &protoSchema.MixedData{}
		if data.Int1 != nil {
			protoData[i].Int1 = wrapperspb.Int64(*data.Int1)
		}
		if data.Int2 != nil {
			protoData[i].Int2 = wrapperspb.Int64(*data.Int2)
		}
		if data.Int3 != nil {
			protoData[i].Int3 = wrapperspb.Int64(*data.Int3)
		}
		if data.Int4 != nil {
			protoData[i].Int4 = wrapperspb.Int64(*data.Int4)
		}
		if data.Int5 != nil {
			protoData[i].Int5 = wrapperspb.Int64(*data.Int5)
		}
		if data.Int6 != nil {
			protoData[i].Int6 = wrapperspb.Int64(*data.Int6)
		}
		if data.Int7 != nil {
			protoData[i].Int7 = wrapperspb.Int64(*data.Int7)
		}
		if data.Int8 != nil {
			protoData[i].Int8 = wrapperspb.Int64(*data.Int8)
		}
		if data.Int9 != nil {
			protoData[i].Int9 = wrapperspb.Int64(*data.Int9)
		}
		if data.Int10 != nil {
			protoData[i].Int10 = wrapperspb.Int64(*data.Int10)
		}

		if data.String1 != nil {
			protoData[i].String1 = wrapperspb.String(*data.String1)
		}
		if data.String2 != nil {
			protoData[i].String2 = wrapperspb.String(*data.String2)
		}
		if data.String3 != nil {
			protoData[i].String3 = wrapperspb.String(*data.String3)
		}
		if data.String4 != nil {
			protoData[i].String4 = wrapperspb.String(*data.String4)
		}
		if data.String5 != nil {
			protoData[i].String5 = wrapperspb.String(*data.String5)
		}
		if data.String6 != nil {
			protoData[i].String6 = wrapperspb.String(*data.String6)
		}
		if data.String7 != nil {
			protoData[i].String7 = wrapperspb.String(*data.String7)
		}
		if data.String8 != nil {
			protoData[i].String8 = wrapperspb.String(*data.String8)
		}
		if data.String9 != nil {
			protoData[i].String9 = wrapperspb.String(*data.String9)
		}
		if data.String10 != nil {
			protoData[i].String10 = wrapperspb.String(*data.String10)
		}
	}

	return &protoSchema.MixedDataSlice{Data: protoData}
}

func LoadIntProtoData(filename string, rowsToRead int64) *protoSchema.IntDataSlice {
	dataSlice := LoadIntData(filename, rowsToRead)

	protoData := make([]*protoSchema.IntData, len(dataSlice))

	for i, data := range dataSlice {
		protoData[i] = &protoSchema.IntData{}
		if data.Int1 != nil {
			protoData[i].Int1 = wrapperspb.Int64(*data.Int1)
		}
		if data.Int2 != nil {
			protoData[i].Int2 = wrapperspb.Int64(*data.Int2)
		}
		if data.Int3 != nil {
			protoData[i].Int3 = wrapperspb.Int64(*data.Int3)
		}
		if data.Int4 != nil {
			protoData[i].Int4 = wrapperspb.Int64(*data.Int4)
		}
		if data.Int5 != nil {
			protoData[i].Int5 = wrapperspb.Int64(*data.Int5)
		}
		if data.Int6 != nil {
			protoData[i].Int6 = wrapperspb.Int64(*data.Int6)
		}
		if data.Int7 != nil {
			protoData[i].Int7 = wrapperspb.Int64(*data.Int7)
		}
		if data.Int8 != nil {
			protoData[i].Int8 = wrapperspb.Int64(*data.Int8)
		}
		if data.Int9 != nil {
			protoData[i].Int9 = wrapperspb.Int64(*data.Int9)
		}
		if data.Int10 != nil {
			protoData[i].Int10 = wrapperspb.Int64(*data.Int10)
		}
		if data.Int11 != nil {
			protoData[i].Int11 = wrapperspb.Int64(*data.Int11)
		}
		if data.Int12 != nil {
			protoData[i].Int12 = wrapperspb.Int64(*data.Int12)
		}
		if data.Int13 != nil {
			protoData[i].Int13 = wrapperspb.Int64(*data.Int13)
		}
		if data.Int14 != nil {
			protoData[i].Int14 = wrapperspb.Int64(*data.Int14)
		}
		if data.Int15 != nil {
			protoData[i].Int15 = wrapperspb.Int64(*data.Int15)
		}
		if data.Int16 != nil {
			protoData[i].Int16 = wrapperspb.Int64(*data.Int16)
		}
		if data.Int17 != nil {
			protoData[i].Int17 = wrapperspb.Int64(*data.Int17)
		}
		if data.Int18 != nil {
			protoData[i].Int18 = wrapperspb.Int64(*data.Int18)
		}
		if data.Int19 != nil {
			protoData[i].Int19 = wrapperspb.Int64(*data.Int19)
		}
		if data.Int20 != nil {
			protoData[i].Int20 = wrapperspb.Int64(*data.Int20)
		}

	}

	return &protoSchema.IntDataSlice{Data: protoData}
}

func LoadStringsProtoData(filename string, rowsToRead int64) *protoSchema.StringDataSlice {
	dataSlice := LoadStringsData(filename, rowsToRead)

	protoData := make([]*protoSchema.StringData, len(dataSlice))

	for i, data := range dataSlice {
		protoData[i] = &protoSchema.StringData{}
		if data.String1 != nil {
			protoData[i].String1 = wrapperspb.String(*data.String1)
		}
		if data.String2 != nil {
			protoData[i].String2 = wrapperspb.String(*data.String2)
		}
		if data.String3 != nil {
			protoData[i].String3 = wrapperspb.String(*data.String3)
		}
		if data.String4 != nil {
			protoData[i].String4 = wrapperspb.String(*data.String4)
		}
		if data.String5 != nil {
			protoData[i].String5 = wrapperspb.String(*data.String5)
		}
		if data.String6 != nil {
			protoData[i].String6 = wrapperspb.String(*data.String6)
		}
		if data.String7 != nil {
			protoData[i].String7 = wrapperspb.String(*data.String7)
		}
		if data.String8 != nil {
			protoData[i].String8 = wrapperspb.String(*data.String8)
		}
		if data.String9 != nil {
			protoData[i].String9 = wrapperspb.String(*data.String9)
		}
		if data.String10 != nil {
			protoData[i].String10 = wrapperspb.String(*data.String10)
		}
		if data.String11 != nil {
			protoData[i].String11 = wrapperspb.String(*data.String11)
		}
		if data.String12 != nil {
			protoData[i].String12 = wrapperspb.String(*data.String12)
		}
		if data.String13 != nil {
			protoData[i].String13 = wrapperspb.String(*data.String13)
		}
		if data.String14 != nil {
			protoData[i].String14 = wrapperspb.String(*data.String14)
		}
		if data.String15 != nil {
			protoData[i].String15 = wrapperspb.String(*data.String15)
		}
		if data.String16 != nil {
			protoData[i].String16 = wrapperspb.String(*data.String16)
		}
		if data.String17 != nil {
			protoData[i].String17 = wrapperspb.String(*data.String17)
		}
		if data.String18 != nil {
			protoData[i].String18 = wrapperspb.String(*data.String18)
		}
		if data.String19 != nil {
			protoData[i].String19 = wrapperspb.String(*data.String19)
		}
		if data.String20 != nil {
			protoData[i].String20 = wrapperspb.String(*data.String20)
		}
	}

	return &protoSchema.StringDataSlice{Data: protoData}
}

func LoadAvroMixedData(filename string, rowsToRead int64) []map[string]interface{} {
	dataSlice := LoadMixedData(filename, rowsToRead)

	protoData := make([]map[string]interface{}, len(dataSlice))

	for i, data := range dataSlice {
		protoData[i] = data.ToMap()
	}

	return protoData
}
func LoadAvroIntData(filename string, rowsToRead int64) []map[string]interface{} {
	dataSlice := LoadIntData(filename, rowsToRead)

	protoData := make([]map[string]interface{}, len(dataSlice))

	for i, data := range dataSlice {
		protoData[i] = data.ToMap()
	}

	return protoData
}
func LoadAvroStringsData(filename string, rowsToRead int64) []map[string]interface{} {
	dataSlice := LoadStringsData(filename, rowsToRead)

	protoData := make([]map[string]interface{}, len(dataSlice))

	for i, data := range dataSlice {
		protoData[i] = data.ToMap()
	}

	return protoData
}

var AvroMixedCodec = getAvroMixedCodec()
var AvroIntCodec = getAvroIntCodec()
var AvroStringCodec = getAvroStringsCodec()

func getAvroMixedCodec() *goavro.Codec {
	codec, err := goavro.NewCodec(`
	{
		"type": "array",
		"items": {
          "type": "record",
          "name": "data",
          "fields" : [
            {
				"name": "int1", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int2", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int3", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int4", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int5", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int6", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int7", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int8", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int9", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int10", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "string1", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string2", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string3", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string4", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string5", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string6", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string7", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string8", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string9", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string10", 
				"type": ["null", "string"], 
				"default": null
			}
          ]
        }
	}`,
	)
	if err != nil {
		fmt.Println(err)
	}

	return codec
}

func getAvroIntCodec() *goavro.Codec {
	codec, err := goavro.NewCodec(`
	{
		"type": "array",
		"items": {
          "type": "record",
          "name": "data",
          "fields" : [
            {
				"name": "int1", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int2", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int3", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int4", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int5", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int6", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int7", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int8", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int9", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int10", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int11", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int12", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int13", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int14", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int15", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int16", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int17", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int18", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int19", 
				"type": ["null", "long"], 
				"default": null
			},
            {
				"name": "int20", 
				"type": ["null", "long"], 
				"default": null
			}
          ]
        }
	}`,
	)
	if err != nil {
		fmt.Println(err)
	}

	return codec
}

func getAvroStringsCodec() *goavro.Codec {
	codec, err := goavro.NewCodec(`
	{
		"type": "array",
		"items": {
          "type": "record",
          "name": "data",
          "fields" : [
            {
				"name": "string1", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string2", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string3", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string4", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string5", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string6", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string7", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string8", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string9", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string10", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string11", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string12", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string13", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string14", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string15", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string16", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string17", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string18", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string19", 
				"type": ["null", "string"], 
				"default": null
			},
            {
				"name": "string20", 
				"type": ["null", "string"], 
				"default": null
			}
          ]
        }
	}`,
	)
	if err != nil {
		fmt.Println(err)
	}

	return codec
}
