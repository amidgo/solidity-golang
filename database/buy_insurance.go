package database

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
)

type BuyIns struct {
	Address common.Address `json:"address"`
	Amount  *big.Int       `json:"amount"`
}

func BuyInsuranse(w http.ResponseWriter, r *http.Request) {
	var addr BuyIns
	json.NewDecoder(r.Body).Decode(&addr)
	acc := variables.ImportAccount(addr.Address)
	tOpts, err := variables.NewTransactOptions(*acc, addr.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = variables.Database.BuyInsurance(tOpts, addr.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
