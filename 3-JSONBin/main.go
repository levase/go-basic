package main

import (
	"fmt"
	"homework/JSONBin/bins"
	"log"
)

func main() {
	bin, err := bins.NewBin("1", "John", false)
	if err != nil {
		log.Fatal(err)
	}

	binList := bins.BinList{}
	binList = append(binList, bin)

	fmt.Println(binList)
}
