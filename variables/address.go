package variables

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

var ZeroAddr = common.HexToAddress("0")

func ImportAccount(addr common.Address) *accounts.Account {
	for _, a := range KeyStore.Accounts() {
		if a.Address == addr {
			return &a
		}
	}
	return nil
}
