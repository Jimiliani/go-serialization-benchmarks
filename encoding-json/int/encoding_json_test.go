package encoding_json_int

import (
	"encoding/json"
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadIntData("../../loader/data/int.csv", 100_000)
var newDataLarge = make([]loader.IntData, 0, 100_000)
var serializedLarge, _ = json.Marshal(dataLarge)

func BenchmarkEncodingJSONSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(dataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodingJSONDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedLarge, &newDataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadIntData("../../loader/data/int.csv", 100)
var newDataMedium = make([]loader.IntData, 0, 100)
var serializedMedium, _ = json.Marshal(dataMedium)

func BenchmarkEncodingJSONSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(dataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodingJSONDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedMedium, &newDataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadIntData("../../loader/data/int.csv", 1)
var newDataSmall = make([]loader.IntData, 0, 1)
var serializedSmall, _ = json.Marshal(dataSmall)

func BenchmarkEncodingJSONSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(dataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEncodingJSONDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(serializedSmall, &newDataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
