package msgpack_mixed

import (
	"fmt"
	"testing"

	"github.com/vmihailenco/msgpack/v5"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadMixedData("../../loader/data/mixed.csv", 100_000)
var newDataLarge = make([]loader.MixedData, 0, 100_000)
var serializedLarge, _ = msgpack.Marshal(dataLarge)

func BenchmarkMsgpackSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := msgpack.Marshal(dataLarge)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := msgpack.Marshal(dataLarge)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkMsgpackDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := msgpack.Unmarshal(serializedLarge, &newDataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadMixedData("../../loader/data/mixed.csv", 100)
var newDataMedium = make([]loader.MixedData, 0, 100)
var serializedMedium, _ = msgpack.Marshal(dataMedium)

func BenchmarkMsgpackSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := msgpack.Marshal(dataMedium)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := msgpack.Marshal(dataMedium)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkMsgpackDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := msgpack.Unmarshal(serializedMedium, &newDataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadMixedData("../../loader/data/mixed.csv", 1)
var newDataSmall = make([]loader.MixedData, 0, 1)
var serializedSmall, _ = msgpack.Marshal(dataSmall)

func BenchmarkMsgpackSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := msgpack.Marshal(dataSmall)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := msgpack.Marshal(dataSmall)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkMsgpackDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := msgpack.Unmarshal(serializedSmall, &newDataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
