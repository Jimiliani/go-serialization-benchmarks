package avro_textual_strings

import (
	"fmt"
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadAvroStringsData("../../loader/data/strings.csv", 100_000)
var serializedLarge, _ = loader.AvroStringCodec.TextualFromNative(nil, dataLarge)

func BenchmarkAvroTextualSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroStringCodec.TextualFromNative(nil, dataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := loader.AvroStringCodec.TextualFromNative(nil, dataLarge)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkAvroTextualDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroStringCodec.NativeFromTextual(serializedLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadAvroStringsData("../../loader/data/strings.csv", 100)
var serializedMedium, _ = loader.AvroStringCodec.TextualFromNative(nil, dataMedium)

func BenchmarkAvroTextualSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroStringCodec.TextualFromNative(nil, dataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := loader.AvroStringCodec.TextualFromNative(nil, dataMedium)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkAvroTextualDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroStringCodec.NativeFromTextual(serializedMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadAvroStringsData("../../loader/data/strings.csv", 1)
var serializedSmall, _ = loader.AvroStringCodec.TextualFromNative(nil, dataSmall)

func BenchmarkAvroTextualSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroStringCodec.TextualFromNative(nil, dataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := loader.AvroStringCodec.TextualFromNative(nil, dataSmall)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkAvroTextualDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroStringCodec.NativeFromTextual(serializedSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
