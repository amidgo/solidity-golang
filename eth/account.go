package eth

import (
	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/accounts"
)

func NewAccount(password string) (*accounts.Account, error) {
	acc, err := variables.KeyStore.NewAccount(password)
	variables.KeyStore.Unlock(acc, password)
	return &acc, err
}
