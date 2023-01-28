package eth

import (
	"context"
	"math/big"

	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
)

func Balance(addr common.Address) (*big.Int, error) {
	bl, err := variables.Client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	balance, err := variables.Client.BalanceAt(context.Background(), addr, big.NewInt(int64(bl)))
	return balance, err
}
