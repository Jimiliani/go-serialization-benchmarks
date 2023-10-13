package cbor_strings

import (
	"testing"

	"github.com/fxamacker/cbor/v2"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadStringsData("../../loader/data/strings.csv", 100_000)
var newDataLarge = make([]loader.StringData, 0, 100_000)
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

var dataMedium = loader.LoadStringsData("../../loader/data/strings.csv", 100)
var newDataMedium = make([]loader.StringData, 0, 100)
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

var dataSmall = loader.LoadStringsData("../../loader/data/strings.csv", 1)
var newDataSmall = make([]loader.StringData, 0, 1)
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
