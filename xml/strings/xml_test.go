package xml_strings

import (
	"encoding/xml"
	"fmt"
	"testing"

	"goFormatsBenchmarking/loader"
)

var dataLarge = loader.LoadStringsData("../../loader/data/strings.csv", 100_000)
var newDataLarge = make([]loader.StringData, 0, 100_000)
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

var dataMedium = loader.LoadStringsData("../loader/data/strings.csv", 100)
var newDataMedium = make([]loader.StringData, 0, 100)
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

var dataSmall = loader.LoadStringsData("../loader/data/strings.csv", 1)
var newDataSmall = make([]loader.StringData, 0, 1)
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
