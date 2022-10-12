package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}

	return records
}

// CreateCSV : controller for generating CSV and sending
func CreateCSV(c *gin.Context) {
	var csvStruct [][]string
	csvStruct = [][]string{
	{"name", "address", "phone"},
	{"Ram","Tokyo","1236524"},
	{"Shaym","Beijing","8575675484"},
}
b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(csvStruct)
	c.Writer.Write(b.Bytes())
}

func main() {
	records := readCsvFile("C://Users/golebiew/OneDrive - TomTom/Desktop/dataBrazil.csv")
	fmt.Println(records)

	// initialize new gin engine (for server)
	r := gin.Default()// routes definition for generating CSV
	r.GET("/csv-download", CreateCSV)// start the server
	r.Run(":5000")
}
