package database

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
)

type SAddress struct {
	Address common.Address `json:"address"`
}

func PayFine(w http.ResponseWriter, r *http.Request) {
	var addr SAddress
	err := json.NewDecoder(r.Body).Decode(&addr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f, err := variables.Database.DriverFines(variables.DefaultCallOpts(), addr.Address, big.NewInt(0))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if f == nil {
		http.Error(w, "not fines", http.StatusBadRequest)
		return
	}
	acc := variables.ImportAccount(addr.Address)
	tOpts, err := variables.NewTransactOptions(*acc, variables.Ether(10))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = variables.Database.PayFine(tOpts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
