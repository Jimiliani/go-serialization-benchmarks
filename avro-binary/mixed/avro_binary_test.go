package avro_binary_mixed

import (
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadAvroMixedData("../../loader/data/mixed.csv", 100_000)
var serializedLarge, _ = loader.AvroMixedCodec.BinaryFromNative(nil, dataLarge)

func BenchmarkAvroBinarySerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroMixedCodec.BinaryFromNative(nil, dataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroMixedCodec.NativeFromBinary(serializedLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadAvroMixedData("../../loader/data/mixed.csv", 100)
var serializedMedium, _ = loader.AvroMixedCodec.BinaryFromNative(nil, dataMedium)

func BenchmarkAvroBinarySerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroMixedCodec.BinaryFromNative(nil, dataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroMixedCodec.NativeFromBinary(serializedMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadAvroMixedData("../../loader/data/mixed.csv", 1)
var serializedSmall, _ = loader.AvroMixedCodec.BinaryFromNative(nil, dataSmall)

func BenchmarkAvroBinarySerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroMixedCodec.BinaryFromNative(nil, dataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroMixedCodec.NativeFromBinary(serializedSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
