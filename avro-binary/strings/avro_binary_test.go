package avro_binary_strings

import (
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadAvroStringsData("../../loader/data/strings.csv", 100_000)
var serializedLarge, _ = loader.AvroStringCodec.BinaryFromNative(nil, dataLarge)

func BenchmarkAvroBinarySerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroStringCodec.BinaryFromNative(nil, dataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroStringCodec.NativeFromBinary(serializedLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadAvroStringsData("../../loader/data/strings.csv", 100)
var serializedMedium, _ = loader.AvroStringCodec.BinaryFromNative(nil, dataMedium)

func BenchmarkAvroBinarySerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroStringCodec.BinaryFromNative(nil, dataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroStringCodec.NativeFromBinary(serializedMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadAvroStringsData("../../loader/data/strings.csv", 1)
var serializedSmall, _ = loader.AvroStringCodec.BinaryFromNative(nil, dataSmall)

func BenchmarkAvroBinarySerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := loader.AvroStringCodec.BinaryFromNative(nil, dataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAvroBinaryDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := loader.AvroStringCodec.NativeFromBinary(serializedSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
