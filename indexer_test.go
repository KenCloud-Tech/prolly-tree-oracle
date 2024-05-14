package Prolly_Tree_Oracle

import (
	"context"
	"fmt"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/gogo/protobuf/test"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/printer"
	"strings"
)

func Indexer(t *test.T) {
	ctx := context.Background()
	// create memory db
	db, err := indexer.NewMemoryDatabase()
	if err != nil {
		panic(err)
	}

	// input data
	reader := strings.NewReader(`{"name":"Alice", "age": 18}
					{"name":"Bob", "age": 19}
					{"name":"Albert", "age": 20}
					{"name":"Clearance and Steve", "age":18}`)

	// create collection and set primary key
	collection, err := db.Collection("users", "name")
	if err != nil {
		panic(err)
	}

	// create index in age field
	_, err = collection.CreateIndex(ctx, "age")
	if err != nil {
		panic(err)
	}

	// insert data
	err = collection.IndexNDJSON(ctx, reader)
	if err != nil {
		panic(err)
	}

	// iterate all records and print
	records, err := collection.Iterate(ctx)
	if err != nil {
		panic(err)
	}
	for record := range records {
		fmt.Println(printer.Sprint(record.Data))
	}

	// export all data of db as car file
	err = db.ExportToFile(ctx, "./init.car")
	if err != nil {
		panic(err)
	}

	// query record
	query := indexer.Query{
		Equal: map[string]ipld.Node{
			"name": basicnode.NewString("Bob"),
		},
	}
	results, err := collection.Search(ctx, query)
	if err != nil {
		panic(err)
	}

	record := <-results
	fmt.Println(printer.Sprint(record.Data))

	// get proof of record(ProllyTree path)
	proof, err := collection.GetProof(record.Id)
	if err != nil {
		panic(err)
	}

	fmt.Println(proof)

	// query record with index field
	query = indexer.Query{
		Equal: map[string]ipld.Node{
			"age": basicnode.NewInt(18),
		},
	}
	results, err = collection.Search(ctx, query)
	if err != nil {
		panic(err)
	}

	// two records
	for record = range results {
		fmt.Println(printer.Sprint(record.Data))
	}

}
