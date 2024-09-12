package test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ipld/go-ipld-prime/printer"
)

func TestImport(t *testing.T) {
	ctx := context.Background()
	db, _ := indexer.NewMemoryDatabase()
	collection, _ := db.Collection(ctx, "demo", "name")

	resp, err := http.Get("http://127.0.0.1:8080/ndjson")
	if err != nil {
		fmt.Println("请求失败：", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	reader := strings.NewReader(string(body))
	err = collection.IndexNDJSON(ctx, reader)
	if err != nil {
		panic(err)
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
