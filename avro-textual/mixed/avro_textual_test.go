package avro_textual_mixed

import (
	"fmt"
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadAvroMixedData("../../loader/data/mixed.csv", 100_000)
var serializedLarge, _ = loader.AvroMixedCodec.TextualFromNative(nil, dataLarge)

func BenchmarkAvroTextualSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroMixedCodec.TextualFromNative(nil, dataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := loader.AvroMixedCodec.TextualFromNative(nil, dataSmall)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkAvroTextualDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroMixedCodec.NativeFromTextual(serializedLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadAvroMixedData("../../loader/data/mixed.csv", 100)
var serializedMedium, _ = loader.AvroMixedCodec.TextualFromNative(nil, dataMedium)

func BenchmarkAvroTextualSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroMixedCodec.TextualFromNative(nil, dataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := loader.AvroMixedCodec.TextualFromNative(nil, dataMedium)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkAvroTextualDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroMixedCodec.NativeFromTextual(serializedMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadAvroMixedData("../../loader/data/mixed.csv", 1)
var serializedSmall, _ = loader.AvroMixedCodec.TextualFromNative(nil, dataSmall)

func BenchmarkAvroTextualSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroMixedCodec.TextualFromNative(nil, dataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := loader.AvroMixedCodec.TextualFromNative(nil, dataSmall)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkAvroTextualDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroMixedCodec.NativeFromTextual(serializedSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
