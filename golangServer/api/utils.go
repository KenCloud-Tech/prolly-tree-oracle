package api

import (
	"Oracle.com/golangServer/config"
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent/qp"
	basicnode "github.com/ipld/go-ipld-prime/node/basicnode"
	"io"
	"log"
	"math/big"
	"strings"
)

func GenTransactOpts(GasLimit uint64) *bind.TransactOpts {
	// Generate TransactOpts from private key
	auth, err := bind.NewKeyedTransactorWithChainID(config.PrivateKey, big.NewInt(config.ChainID))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Set gas prices and gas limits, these can be set more intelligently through client queries
	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
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
