package database

import (
	"encoding/json"
	"math"
	"math/big"
	"net/http"

	"github.com/amidgo/solidity-ethereum/variables"
)

func CalcInsurance(w http.ResponseWriter, r *http.Request) {
	var addr SAddress
	json.NewDecoder(r.Body).Decode(&addr)
	dr, err := variables.Database.Drivers(variables.DefaultCallOpts(), addr.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	cr, err := variables.Database.DriverCars(variables.DefaultCallOpts(), addr.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if cr.Category == "" {
		http.Error(w, "you not have a car", http.StatusBadRequest)
	}
	amount := InsuranceFormula(float32(cr.Cost), float32(cr.AgeCount), float32(dr.FineCount), float32(dr.FineCount), float32(dr.Exp))
	w.Write([]byte(amount.String()))
}

func InsuranceFormula(carCost float32, carAgeCount float32, fineCount float32, dtpCount float32, driverAgeCount float32) *big.Int {
	first := carCost*float32(math.Abs(float64(1-carAgeCount/10)))*0.1 + 0.2*fineCount + dtpCount - 0.2*driverAgeCount
	return variables.Ether(first)
}
