package database

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amidgo/solidity-ethereum/eth"
	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
)

type DriverData struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	Fio       string `json:"fio"`
	Exp       uint8  `json:"exp"`
	DtpCount  uint32 `json:"dtp_count"`
	FineCount uint32 `json:"fine_count"`
}

func RegisterDriver(w http.ResponseWriter, r *http.Request) {
	var drData DriverData
	json.NewDecoder(r.Body).Decode(&drData)
	st, err := variables.Database.Logins(variables.DefaultCallOpts(), drData.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if st.Addr != variables.ZeroAddr {
		http.Error(w, "Login already exist", http.StatusBadRequest)
		return
	}
	acc, err := eth.NewAccount(drData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = eth.SendTransaction(variables.Coinbase, acc.Address, variables.CoinbasePwd, variables.Ether(50)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	addr, err := RegisterDriverWithStruct(&drData, &acc.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(addr))
}

func RegisterDriverWithStruct(d *DriverData, addr *common.Address) (string, error) {
	acc := variables.ImportAccount(*addr)
	variables.KeyStore.Unlock(*acc, d.Password)
	tOpts, err := variables.NewTransactOptions(*acc, variables.Ether(0))
	if err != nil {
		fmt.Println(err.Error())
		return variables.ZeroAddr.Hex(), err
	}
	_, err = variables.Database.AddDriver(tOpts, d.Fio, d.Exp, d.DtpCount, d.FineCount, d.Login, d.Password)
	if err != nil {
		fmt.Println(err.Error())
		return variables.ZeroAddr.Hex(), err
	}
	return addr.Hex(), nil
}
