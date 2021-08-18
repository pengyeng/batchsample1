package main

import (
	"entsample2/reader"
	"entsample2/writer"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pengyeng/batch103"
)

func main() {

	var myJobLauncher1 batch103.JobLauncher
	var myLaptopReader batch103.ReaderType
	var myLaptopDBWriter batch103.WriterType
	myLaptopReader = &reader.LaptopReader{}
	myLaptopDBWriter = &writer.LaptopDBWriter{}

	//Prepare Processor List
	laptopProcessorList := []batch103.ProcessorType{}

	//Prepare Writer List
	laptopWriterList := []batch103.WriterType{}
	laptopWriterList = append(laptopWriterList, myLaptopDBWriter)
	myJobLauncher1.Run(myLaptopReader, laptopProcessorList, laptopWriterList)

	// Run the automatic migration tool to create all schema resources.
	//if err := client.Schema.Create(ctx); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}

}
