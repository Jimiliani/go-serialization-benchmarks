package xml_mixed

import (
	"encoding/xml"
	"fmt"
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.MixedDataSlice{MixedDatas: loader.LoadMixedData("../../loader/data/mixed.csv", 100_000)}
var newDataLarge = loader.MixedDataSlice{MixedDatas: make([]loader.MixedData, 0, 100_000)}
var serializedLarge, _ = xml.Marshal(dataLarge)

func BenchmarkXMLSerializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := xml.Marshal(dataLarge)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := xml.Marshal(dataLarge)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkXMLDeserializeLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := xml.Unmarshal(serializedLarge, &newDataLarge)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataMedium = loader.MixedDataSlice{MixedDatas: loader.LoadMixedData("../../loader/data/mixed.csv", 100)}
var newDataMedium = loader.MixedDataSlice{MixedDatas: make([]loader.MixedData, 0, 100)}
var serializedMedium, _ = xml.Marshal(dataMedium)

func BenchmarkXMLSerializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := xml.Marshal(dataMedium)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := xml.Marshal(dataMedium)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkXMLDeserializeMedium(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := xml.Unmarshal(serializedMedium, &newDataMedium)
		if err != nil {
			b.Fatal(err)
		}
	}
}

var dataSmall = loader.MixedDataSlice{MixedDatas: loader.LoadMixedData("../../loader/data/mixed.csv", 1)}
var newDataSmall = loader.MixedDataSlice{MixedDatas: make([]loader.MixedData, 0, 1)}
var serializedSmall, _ = xml.Marshal(dataSmall)

func BenchmarkXMLSerializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		marshalled, err := xml.Marshal(dataSmall)
		if err != nil {
			fmt.Println(len(marshalled))
			b.Fatal(err)
		}
	}
	b.StopTimer()
	d, _ := xml.Marshal(dataSmall)
	fmt.Printf("Serialized data size: %d bytes (%.2f MB)\n", len(d), float64(len(d))/(1024*1024))
}

func BenchmarkXMLDeserializeSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		err := xml.Unmarshal(serializedSmall, &newDataSmall)
		if err != nil {
			b.Fatal(err)
		}
	}
}
