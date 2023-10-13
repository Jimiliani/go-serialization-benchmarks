package msgpack_int

import (
	"fmt"
	"testing"

	"github.com/vmihailenco/msgpack/v5"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadIntData("../../loader/data/int.csv", 100_000)
var newDataLarge = make([]loader.IntData, 0, 100_000)
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

var dataMedium = loader.LoadIntData("../../loader/data/int.csv", 100)
var newDataMedium = make([]loader.IntData, 0, 100)
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

var dataSmall = loader.LoadIntData("../../loader/data/int.csv", 1)
var newDataSmall = make([]loader.IntData, 0, 1)
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
