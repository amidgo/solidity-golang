package database

import (
	"fmt"

	"github.com/amidgo/solidity-ethereum/variables"
)

func InitAccounts() {
	p := &PoliceMan{
		Addr:      variables.Ivan,
		Login:     "ivan",
		Password:  variables.IvanPwd,
		Fio:       "Иванов Иван Иванович",
		Exp:       2,
		DtpCount:  0,
		FineCount: 0,
	}
	AddPoliceMan(p)
	fmt.Println("police man success")
	sd := &DriverData{
		Login:     "semen",
		Fio:       "Семенов Семен Семенович",
		Password:  variables.SemenPwd,
		Exp:       5,
		FineCount: 0,
		DtpCount:  0,
	}
	RegisterDriverWithStruct(sd, &variables.Semen)
	fmt.Println("semen success")
	pd := &DriverData{
		Login:     "petr",
		Fio:       "Петров Петр Петрович",
		Password:  variables.PetrPwd,
		Exp:       10,
		FineCount: 0,
		DtpCount:  3,
	}
	RegisterDriverWithStruct(pd, &variables.Petr)
	fmt.Println("petr success")
}
