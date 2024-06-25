package test

import (
	"Oracle.com/golangServer/api"
	"context"
	"fmt"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ipld/go-ipld-prime/printer"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestImport(t *testing.T) {
	ctx := context.Background()
	db, _ := indexer.NewMemoryDatabase()
	collection, _ := db.Collection("demo", "Series_Title")

	resp, err := http.Get("http://127.0.0.1:8080/data.csv")
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	reader := strings.NewReader(string(body))
	err = api.IngestCSV(ctx, reader, collection)
	err = db.ApplyChanges(ctx)

	if err != nil {
		fmt.Println(err)
	}
	records, err := collection.Iterate(ctx)
	if err != nil {
		panic(err)
	}
	i := 1
	for record := range records {
		fmt.Println(printer.Sprint(record.Data), i)
		i++
	}
}
