package encoding_json_strings

import (
	"encoding/json"
	"fmt"
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadStringsData("../../loader/data/strings.csv", 100_000)
var newDataLarge = make([]loader.StringData, 0, 100_000)
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
	b.StopTimer()
	d, _ := json.Marshal(dataLarge)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
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

var dataMedium = loader.LoadStringsData("../../loader/data/strings.csv", 100)
var newDataMedium = make([]loader.StringData, 0, 100)
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
	b.StopTimer()
	d, _ := json.Marshal(dataMedium)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
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

var dataSmall = loader.LoadStringsData("../../loader/data/strings.csv", 1)
var newDataSmall = make([]loader.StringData, 0, 1)
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
	b.StopTimer()
	d, _ := json.Marshal(dataSmall)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
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
