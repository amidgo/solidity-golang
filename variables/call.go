package variables

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func NewCallOptions(from common.Address) *bind.CallOpts {
	blockNumber, _ := Client.BlockNumber(context.Background())
	opts := bind.CallOpts{Pending: true, From: from, BlockNumber: big.NewInt(int64(blockNumber)), Context: context.Background()}
	return &opts
}

func DefaultCallOpts() *bind.CallOpts {
	return NewCallOptions(Coinbase)
}
