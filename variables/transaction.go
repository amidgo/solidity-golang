package variables

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func DefaultTransactOptions() *bind.TransactOpts {
	t, _ := NewTransactOptions(*CoinbaseAcc, big.NewInt(0))
	return t
}

func NewTransactOptions(from accounts.Account, value *big.Int) (*bind.TransactOpts, error) {
	chainId, err := Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	gasPrice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyStoreTransactorWithChainID(KeyStore, from, chainId)
	auth.Value = value
	auth.GasPrice = gasPrice

	return auth, err
}
