package database

import (
	"fmt"

	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
)

type PoliceMan struct {
	Addr      common.Address `json:"addr"`
	Login     string         `json:"login"`
	Password  string         `json:"password"`
	Fio       string         `json:"fio"`
	Exp       uint8          `json:"exp"`
	DtpCount  uint32         `json:"dtp_count"`
	FineCount uint32         `json:"fine_count"`
}

func AddPoliceMan(p *PoliceMan) {
	account := variables.ImportAccount(p.Addr)
	variables.KeyStore.Unlock(*account, p.Password)
	trOpts, err := variables.NewTransactOptions(*account, variables.Ether(0))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = variables.Database.AddPoliceMan(trOpts, p.Fio, p.Exp, p.DtpCount, p.FineCount, p.Login, p.Password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
