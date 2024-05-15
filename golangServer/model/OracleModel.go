package model

import (
	"context"
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
)

type OracleModel struct {
	DbName string
	Ctx    context.Context
	Db     *indexer.Collection
}
