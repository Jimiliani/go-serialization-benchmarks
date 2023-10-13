package avro_textual_int

import (
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadAvroIntData("../../loader/data/int.csv", 100_000)
var serializedLarge, _ = loader.AvroIntCodec.TextualFromNative(nil, dataLarge)

func BenchmarkAvroTextualSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroIntCodec.TextualFromNative(nil, dataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroTextualDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroIntCodec.NativeFromTextual(serializedLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadAvroIntData("../../loader/data/int.csv", 100)
var serializedMedium, _ = loader.AvroIntCodec.TextualFromNative(nil, dataMedium)

func BenchmarkAvroTextualSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroIntCodec.TextualFromNative(nil, dataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroTextualDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroIntCodec.NativeFromTextual(serializedMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadAvroIntData("../../loader/data/int.csv", 1)
var serializedSmall, _ = loader.AvroIntCodec.TextualFromNative(nil, dataSmall)

func BenchmarkAvroTextualSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroIntCodec.TextualFromNative(nil, dataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroTextualDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroIntCodec.NativeFromTextual(serializedSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
