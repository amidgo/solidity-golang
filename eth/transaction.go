package eth

import (
	"context"
	"math/big"
	"time"

	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func SendTransaction(from, to common.Address, password string, value *big.Int) error {
	nonce, err := variables.Client.PendingNonceAt(context.Background(), from)
	if err != nil {
		return err
	}
	gasLimit := uint64(30000)
	gasPrice, err := variables.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil)
	chainId, err := variables.Client.NetworkID(context.Background())
	if err != nil {
		return err
	}
	account := variables.ImportAccount(from)
	signedTx, err := variables.KeyStore.SignTxWithPassphrase(*account, password, tx, chainId)
	if err != nil {
		return err
	}
	err = variables.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}
	time.Sleep(time.Second * 2)

	return err
}
