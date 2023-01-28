package database

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/amidgo/solidity-ethereum/variables"
)

func RenewInsurance(w http.ResponseWriter, r *http.Request) {
	var addr SAddress
	json.NewDecoder(r.Body).Decode(&addr)
	acc := variables.ImportAccount(addr.Address)
	tOpts, err := variables.NewTransactOptions(*acc, big.NewInt(0))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = variables.Database.RenewInsurance(tOpts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
