package go_json_mixed

import (
	"fmt"
	"testing"

	"github.com/goccy/go-json"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadMixedData("../../loader/data/mixed.csv", 100_000)
var newDataLarge = make([]loader.MixedData, 0, 100_000)
var serializedLarge, _ = json.Marshal(dataLarge)

func BenchmarkGoJSONSerializeLarge(b *testing.B) {
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

func BenchmarkGoJSONDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedLarge, &newDataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadMixedData("../../loader/data/mixed.csv", 100)
var newDataMedium = make([]loader.MixedData, 0, 100)
var serializedMedium, _ = json.Marshal(dataMedium)

func BenchmarkGoJSONSerializeMedium(b *testing.B) {
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

func BenchmarkGoJSONDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedMedium, &newDataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadMixedData("../../loader/data/mixed.csv", 1)
var newDataSmall = make([]loader.MixedData, 0, 1)
var serializedSmall, _ = json.Marshal(dataSmall)

func BenchmarkGoJSONSerializeSmall(b *testing.B) {
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

func BenchmarkGoJSONDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedSmall, &newDataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
