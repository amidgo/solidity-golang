package database

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/amidgo/solidity-ethereum/eth"
	"github.com/amidgo/solidity-ethereum/variables"
)

type DriverInfo struct {
	Balance        *big.Int `json:"balance"`
	Fio            string   `json:"fio"`
	Exp            uint8    `json:"exp"`
	DtpCount       uint32   `json:"dtp_count"`
	FineCount      uint32   `json:"fine_count"`
	InsuranceValue *big.Int `json:"insurance_value"`
	DocNumber      string   `json:"doc_number"`
	ValidateTime   uint64   `json:"validate_time"`
	DocCategory    Ctg      `json:"doc_category"`
	CarCost        uint32   `json:"car_cost"`
	CarAge         uint32   `json:"car_age"`
}

func Info(w http.ResponseWriter, r *http.Request) {
	var addr SAddress
	json.NewDecoder(r.Body).Decode(&addr)
	dr, err := variables.Database.Drivers(variables.DefaultCallOpts(), addr.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	car, err := variables.Database.DriverCars(variables.DefaultCallOpts(), addr.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dc, err := variables.Database.DriverDocuments(variables.DefaultCallOpts(), addr.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, err := eth.Balance(addr.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ins, err := variables.Database.DriverInsurances(variables.DefaultCallOpts(), addr.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	j, err := json.Marshal(
		DriverInfo{
			Balance:        b,
			Fio:            dr.Fio,
			Exp:            dr.Exp,
			FineCount:      dr.FineCount,
			DtpCount:       dr.DtpCount,
			InsuranceValue: ins,
			DocNumber:      dc.Number,
			DocCategory:    Ctg(dc.Category),
			ValidateTime:   dc.ValidateTime,
			CarCost:        car.Cost,
			CarAge:         car.AgeCount,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(j)
}
