package api

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"strings"

	"Oracle.com/golangServer/config"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	basicnode "github.com/ipld/go-ipld-prime/node/basicnode"
)

func GenTransactOpts(ctx context.Context, GasLimit uint64) *bind.TransactOpts {
	// Generate TransactOpts from private key
	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, big.NewInt(config.ChainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set gas prices and gas limits, these can be set more intelligently through client queries
	gasPrice, err := config.Client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatalf("Failed to suggest gas price: %v", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = GasLimit
	return auth
}

func nodeTobyte(node ipld.Node) []byte {
	var buf bytes.Buffer
	if err := dagjson.Encode(node, &buf); err != nil {
		log.Fatalf("Trans node to json fail: %v", err)
	}

	var formatted bytes.Buffer
	if err := json.Indent(&formatted, buf.Bytes(), "", "  "); err != nil {
		log.Fatalf("Format json fail: %v", err)
	}
	return formatted.Bytes()
}

func IngestCSV(ctx context.Context, source io.Reader, collection *indexer.Collection) error {
	reader := csv.NewReader(source)

	headers, err := reader.Read()

	numFields := int64(len(headers) + 1)

	if err != nil {
		return err
	}

	index := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		node, err := qp.BuildMap(basicnode.Prototype.Any, numFields, func(ma datamodel.MapAssembler) {
			qp.MapEntry(ma, "index", qp.Int(int64(index)))
			for fieldIndex, fieldValue := range record {
				nb := basicnode.Prototype__Any{}.NewBuilder()
				err := dagjson.Decode(nb, strings.NewReader(fieldValue))

				// If it wasn't json, it's just a string
				if err != nil {
					qp.MapEntry(ma, headers[fieldIndex], qp.String(fieldValue))
				} else {
					value := nb.Build()
					qp.MapEntry(ma, headers[fieldIndex], qp.Node(value))
				}
			}
		})

		if err != nil {
			return err
		}

		err = collection.Insert(ctx, node)

		if err != nil {
			return err
		}

		index++
	}

	return nil
}

type Eq struct {
	IndexName    string
	IndexVal     any
	IndexValType string
}
type Cmp struct {
	//	GreaterThan Op = "GreaterThan"
	//	LessThan    Op = "LessThan"
	Op           string
	IndexName    string
	IndexVal     any
	IndexValType string
}
type Queryer struct {
	Equal   Eq
	Compare Cmp
	Sort    string
	Limit   int
	Skip    int
}

func (q Queryer) Queryer2indexerQ() (nq indexer.Query) {
	if q.Equal.IndexName != "" {
		val := q.Equal.IndexVal
		switch q.Equal.IndexValType {
		case "string":
			nq = indexer.Query{
				Equal: map[string]ipld.Node{q.Equal.IndexName: basicnode.NewString(val.(string))},
				Sort:  q.Sort,
				Skip:  q.Skip,
				Limit: q.Limit,
			}

		case "int":
			nq = indexer.Query{
				Equal: map[string]ipld.Node{q.Equal.IndexName: basicnode.NewInt(int64(val.(float64)))},
				Sort:  q.Sort,
				Skip:  q.Skip,
				Limit: q.Limit,
			}

		case "float":
			nq = indexer.Query{
				Equal: map[string]ipld.Node{q.Equal.IndexName: basicnode.NewFloat(val.(float64))},
				Sort:  q.Sort,
				Skip:  q.Skip,
				Limit: q.Limit,
			}

		case "byte":
			nq = indexer.Query{
				Equal: map[string]ipld.Node{q.Equal.IndexName: basicnode.NewBytes(val.([]byte))},
				Sort:  q.Sort,
				Skip:  q.Skip,
				Limit: q.Limit,
			}

		}
	}
	if q.Compare.Op != "" {
		val := q.Compare.IndexVal
		switch q.Compare.IndexValType {
		case "string":
			nd := basicnode.NewString(val.(string))
			nq = indexer.Query{
				Compare: &indexer.CompareCondition{
					Cmp:       indexer.Op(q.Compare.Op),
					IndexName: q.Compare.IndexName,
					IndexVal:  nd,
				},
				Sort:  q.Sort,
				Skip:  q.Skip,
				Limit: q.Limit,
			}
		case "int":
			nd := basicnode.NewInt(int64(val.(float64)))
			nq = indexer.Query{
				Compare: &indexer.CompareCondition{
					Cmp:       indexer.Op(q.Compare.Op),
					IndexName: q.Compare.IndexName,
					IndexVal:  nd,
				},
				Sort:  q.Sort,
				Skip:  q.Skip,
				Limit: q.Limit,
			}
		case "float":
			nd := basicnode.NewFloat(val.(float64))
			nq = indexer.Query{
				Compare: &indexer.CompareCondition{
					Cmp:       indexer.Op(q.Compare.Op),
					IndexName: q.Compare.IndexName,
					IndexVal:  nd,
				},
				Sort:  q.Sort,
				Skip:  q.Skip,
				Limit: q.Limit,
			}
		case "byte":
			nd := basicnode.NewBytes(val.([]byte))
			nq = indexer.Query{
				Compare: &indexer.CompareCondition{
					Cmp:       indexer.Op(q.Compare.Op),
					IndexName: q.Compare.IndexName,
					IndexVal:  nd,
				},
				Sort:  q.Sort,
				Skip:  q.Skip,
				Limit: q.Limit,
			}
		}
	}
	return
}
