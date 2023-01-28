package database

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/amidgo/solidity-ethereum/eth"
	"github.com/amidgo/solidity-ethereum/variables"
)

type DocNumber struct {
	Number string `json:"number"`
}

func AddDtp(w http.ResponseWriter, r *http.Request) {
	var n DocNumber
	json.NewDecoder(r.Body).Decode(&n)
	acc := variables.ImportAccount(variables.Coinbase)
	tOpts, err := variables.NewTransactOptions(*acc, big.NewInt(0))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = variables.Database.AddDTP(tOpts, n.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	drA, err := variables.Database.DriverDocNumbers(variables.DefaultCallOpts(), n.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if drA == variables.ZeroAddr {
		http.Error(w, "not found", http.StatusBadRequest)
		return
	}

	amount, err := variables.Database.DriverInsurances(variables.DefaultCallOpts(), drA)
	if err != nil {
		http.Error(w, "not found", http.StatusBadRequest)
		return
	}
	fmt.Println(amount)
	fmt.Println(amount.Cmp(big.NewInt(0)))
	fmt.Println(drA)
	if amount.Cmp(big.NewInt(0)) != 0 {
		insB, err := eth.Balance(variables.Insurance)
		if err != nil {
			http.Error(w, "not found", http.StatusBadRequest)
			return
		}
		sum := big.NewInt(0).Mul(amount, big.NewInt(10))
		if insB.Cmp(sum) == -1 {
			diff := big.NewInt(0).Sub(sum, insB)
			err = eth.SendTransaction(variables.Bank, variables.Insurance, variables.BankPwd, big.NewInt(0).Add(diff, big.NewInt(1000000)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			_, err = variables.Database.IncLoan(variables.DefaultTransactOptions(), diff)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		if err := eth.SendTransaction(variables.Insurance, drA, variables.InsurancePwd, sum); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}
