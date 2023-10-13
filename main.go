package main

import (
	"fmt"

	"goFormatsBenchmarking/loader"
)

func main() {

	mixedData := loader.LoadMixedData("../loader/data/mixed.csv", 100_000)
	intData := loader.LoadIntData("../loader/data/int.csv", 100_000)
	stringsData := loader.LoadStringsData("../loader/data/strings.csv", 100_000)

	fmt.Println(mixedData[0])
	fmt.Println(intData[0])
	fmt.Println(stringsData[0])
}
