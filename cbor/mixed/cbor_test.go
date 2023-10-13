package cbor_mixed

import (
	"fmt"
	"testing"

	"github.com/fxamacker/cbor/v2"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadMixedData("../../loader/data/mixed.csv", 100_000)
var newDataLarge = make([]loader.MixedData, 0, 100_000)
var serializedLarge, _ = cbor.Marshal(dataLarge)

func BenchmarkCBORSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := cbor.Marshal(dataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := cbor.Marshal(dataLarge)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkCBORDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := cbor.Unmarshal(serializedLarge, &newDataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadMixedData("../../loader/data/mixed.csv", 100)
var newDataMedium = make([]loader.MixedData, 0, 100)
var serializedMedium, _ = cbor.Marshal(dataMedium)

func BenchmarkCBORSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := cbor.Marshal(dataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := cbor.Marshal(dataMedium)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkCBORDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := cbor.Unmarshal(serializedMedium, &newDataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadMixedData("../../loader/data/mixed.csv", 1)
var newDataSmall = make([]loader.MixedData, 0, 1)
var serializedSmall, _ = cbor.Marshal(dataSmall)

func BenchmarkCBORSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := cbor.Marshal(dataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := cbor.Marshal(dataSmall)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkCBORDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := cbor.Unmarshal(serializedSmall, &newDataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
