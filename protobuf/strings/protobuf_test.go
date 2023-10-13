package protobuf_strings

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"

	"goFormatsBenchmarking/loader"
	protoSchema "goFormatsBenchmarking/protobuf/schema"
)

var dataLarge = loader.LoadMixedProtoData("../../loader/data/strings.csv", 100_000)
var newDataLarge = protoSchema.StringDataSlice{Data: make([]*protoSchema.StringData, 100_000)}
var serializedLarge, _ = proto.Marshal(dataLarge)

func BenchmarkProtobufSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := proto.Marshal(dataLarge)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
}

func BenchmarkProtobufDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(serializedLarge, &newDataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.LoadMixedProtoData("../../loader/data/strings.csv", 100)
var newDataMedium = protoSchema.StringDataSlice{Data: make([]*protoSchema.StringData, 100)}
var serializedMedium, _ = proto.Marshal(dataMedium)

func BenchmarkProtobufSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := proto.Marshal(dataMedium)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
}

func BenchmarkProtobufDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(serializedMedium, &newDataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.LoadMixedProtoData("../../loader/data/strings.csv", 1)
var newDataSmall = protoSchema.StringDataSlice{Data: make([]*protoSchema.StringData, 1)}
var serializedSmall, _ = proto.Marshal(dataSmall)

func BenchmarkProtobufSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := proto.Marshal(dataSmall)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
}

func BenchmarkProtobufDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(serializedSmall, &newDataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
