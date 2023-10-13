package segmentio_encoding_json_strings

import (
	"fmt"
	"testing"

	"github.com/segmentio/encoding/json"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadStringsData("../../loader/data/strings.csv", 100_000)
var newDataLarge = make([]loader.StringData, 0, 100_000)
var serializedLarge, _ = json.Marshal(dataLarge)

func BenchmarkSegmentioEncodingJSONSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := json.Marshal(dataLarge)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
}

func BenchmarkSegmentioEncodingJSONDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedLarge, &newDataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadStringsData("../../loader/data/strings.csv", 100)
var newDataMedium = make([]loader.StringData, 0, 100)
var serializedMedium, _ = json.Marshal(dataMedium)

func BenchmarkSegmentioEncodingJSONSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := json.Marshal(dataMedium)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
}

func BenchmarkSegmentioEncodingJSONDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedMedium, &newDataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadStringsData("../../loader/data/strings.csv", 1)
var newDataSmall = make([]loader.StringData, 0, 1)
var serializedSmall, _ = json.Marshal(dataSmall)

func BenchmarkSegmentioEncodingJSONSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := json.Marshal(dataSmall)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
}

func BenchmarkSegmentioEncodingJSONDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedSmall, &newDataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
