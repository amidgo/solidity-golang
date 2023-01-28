package database

import (
	"encoding/json"
	"net/http"

	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
)

type DriverCar struct {
	Owner    common.Address `json:"owner"`
	Category Ctg            `json:"category"`
	Cost     uint32         `json:"cost"`
	AgeCount uint32         `json:"age_count"`
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	var drC DriverCar
	json.NewDecoder(r.Body).Decode(&drC)
	st, err := variables.Database.DriverDocuments(variables.DefaultCallOpts(), drC.Owner)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if st.Category != string(drC.Category) {
		http.Error(w, "wrong category", http.StatusBadRequest)
		return
	}
	acc := variables.ImportAccount(drC.Owner)
	if acc == nil {
		http.Error(w, "account not found", http.StatusBadRequest)
		return
	}
	tOpts, err := variables.NewTransactOptions(*acc, variables.Ether(0))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = variables.Database.AddDriverCar(tOpts, string(drC.Category), drC.Cost, drC.AgeCount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
