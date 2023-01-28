package database

import (
	"encoding/json"
	"net/http"

	"github.com/amidgo/solidity-ethereum/variables"
)

func AddFine(w http.ResponseWriter, r *http.Request) {
	var n DocNumber
	json.NewDecoder(r.Body).Decode(&n)
	_, err := variables.Database.AddFine(variables.DefaultTransactOptions(), n.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
