package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	csvFile, _ := os.Open("./train.1.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	property := make(map[string]int)

	property["total"] = 0
	property["survived"] = 0

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		} else {
			property["total"]++
			if row[1] == "1" {
				property["survived"]++
			}
		}
	}

	fmt.Println(property)
}
