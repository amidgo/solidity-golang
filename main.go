package main

import (
	"github.com/amidgo/solidity-ethereum/database"
	"github.com/amidgo/solidity-ethereum/routing"
	"github.com/amidgo/solidity-ethereum/variables"
)

func main() {
	variables.Init()
	database.InitAccounts()
	routing.Configure()
}
