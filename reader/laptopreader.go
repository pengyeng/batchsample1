package reader

import (
	"io"
	"log"

	"github.com/pengyeng/batch103"
)

type LaptopReader struct {
	batch103.FileReader
}

func (r *LaptopReader) Read() []batch103.BatchData {
	log.Println("laptop reader")
	r.SetFileName("laptop.csv")
	csvFileReader := r.OpenCSVFile()
	var result []batch103.BatchData
	for {

		record, err := csvFileReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		//Inserting File Record into Result set
		input := []string{record[0], record[1], record[2], record[3]}
		var batchData = &batch103.BatchData{}
		batchData = batchData.Create(input)
		result = append(result, *batchData)
	}
	log.Println("No Of Record Retrieved ", len(result))
	return result
}
