package writer

import (
	"context"
	"entsample2/ent"
	"entsample2/model"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pengyeng/batch103"
)

type LaptopDBWriter struct {
	batch103.DBWriter
}

func (d *LaptopDBWriter) Write(data []batch103.BatchData) {

	// Create an ent.Client with MySQL database.
	client, err := ent.Open(dialect.MySQL, "root:p@ssw0rd@tcp(127.0.0.1:3306)/go-data")
	if err != nil {
		log.Fatalf("failed opening connection to MySQL: %v", err)
	}
	defer client.Close()
	ctx := context.Background()

	for i := 0; i < len(data); i++ {
		var currRec = &model.Laptop{}
		input := data[i].GenericData
		currRec = currRec.Create(input)

		if data[i].IsActive() {

			laptop, err := client.Laptop.Create().SetBrand(currRec.Brand()).SetModel(currRec.Model()).SetCPU(currRec.Cpu()).SetRAM(currRec.Ram()).SetHarddisk("").Save(ctx)
			if err != nil {
				log.Fatalf("failed creating a laptop: %v", err)
			}
			println(laptop)

		}
	}
}
