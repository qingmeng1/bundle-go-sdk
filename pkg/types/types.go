package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Hashes []common.Hash

type BundleStatus uint8

const (
	BundleStatusPending BundleStatus = iota
	BundleStatusConfirmed
	BundleStatusFailed
)

type Bundle struct {
	Hash                 common.Hash
	Txs                  Hashes
	MaxBlockNumber       uint64
	MaxTimestamp         uint64
	Status               BundleStatus
	GasFee               *hexutil.Big
	Builder              common.Address
	ConfirmedBlockNumber uint64
	ConfirmedDate        uint64
}
