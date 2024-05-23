package model

import (
	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/ipfs/go-cid"
)

type OracleModel struct {
	DbOwner string
	Cid     cid.Cid
	Db      *indexer.Database
}
