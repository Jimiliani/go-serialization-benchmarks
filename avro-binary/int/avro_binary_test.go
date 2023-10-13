package avro_binary_int

import (
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadAvroIntData("../../loader/data/int.csv", 100_000)
var serializedLarge, _ = loader.AvroIntCodec.BinaryFromNative(nil, dataLarge)

func BenchmarkAvroBinarySerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroIntCodec.BinaryFromNative(nil, dataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroIntCodec.NativeFromBinary(serializedLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadAvroIntData("../../loader/data/int.csv", 100)
var serializedMedium, _ = loader.AvroIntCodec.BinaryFromNative(nil, dataMedium)

func BenchmarkAvroBinarySerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroIntCodec.BinaryFromNative(nil, dataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroIntCodec.NativeFromBinary(serializedMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadAvroIntData("../../loader/data/int.csv", 1)
var serializedSmall, _ = loader.AvroIntCodec.BinaryFromNative(nil, dataSmall)

func BenchmarkAvroBinarySerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroIntCodec.BinaryFromNative(nil, dataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroIntCodec.NativeFromBinary(serializedSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
