package database

import (
	"encoding/json"
	"net/http"

	"github.com/amidgo/solidity-ethereum/eth"
)

func Balance(w http.ResponseWriter, r *http.Request) {
	var addr SAddress
	json.NewDecoder(r.Body).Decode(&addr)
	b, err := eth.Balance(addr.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(b.String()))
}
