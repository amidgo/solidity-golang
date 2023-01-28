package variables

import (
	"fmt"
	"log"
	"time"

	"github.com/amidgo/solidity-ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const Node1 string = "http://0.0.0.0:1111"
const Node1KeyStore string = "./eth-net/node1/keystore"

const Node2 string = "http://0.0.0.0:2222"
const Node2KeyStore string = "./eth-net/node2/keystore"

const CoinbaseAddr string = "0x87B601B72D92D5D4064888D51cF84f40b12B9D27"
const CoinbasePwd string = "coinbase"
const BankAddr string = "0x051dBEeCA5FD9CBac8a419Fe236C6503D675b19B"
const BankPwd string = "bank"
const InsuranceAddr string = "0xd6ab6767cfEcAe49Ae8EF15A2C72Accd6Cc9FB9c"
const InsurancePwd string = "insurance"
const IvanAddr string = "0xF47571c4ae4ddb7bFAAAF78C341f3941aaf96f86"
const IvanPwd string = "ivan"
const PetrAddr string = "0x42AA518548229a8465138028d0242BF9CA20ef27"
const PetrPwd string = "petr"
const SemenAddr string = "0x01971e73d7A0437ED839F62B1E9370aB34015C83"
const SemenPwd string = "semen"

var Coinbase common.Address = common.HexToAddress(CoinbaseAddr)
var Bank common.Address = common.HexToAddress(BankAddr)
var Insurance common.Address = common.HexToAddress(InsuranceAddr)
var Ivan common.Address = common.HexToAddress(IvanAddr)
var Petr common.Address = common.HexToAddress(PetrAddr)
var Semen common.Address = common.HexToAddress(SemenAddr)

var CoinbaseAcc *accounts.Account
var BankAcc *accounts.Account
var InsuranceAcc *accounts.Account

var Client *ethclient.Client
var KeyStore *keystore.KeyStore
var DatabaseAddr common.Address
var Database *contracts.Contracts

func Init() {
	c, err := ethclient.Dial(Node2)
	if err != nil {
		log.Fatal(err)
	}
	Client = c
	KeyStore = keystore.NewKeyStore(Node2KeyStore, keystore.StandardScryptN, keystore.StandardScryptP)
	CoinbaseAcc = ImportAccount(Coinbase)
	BankAcc = ImportAccount(Bank)
	InsuranceAcc = ImportAccount(Insurance)
	KeyStore.Unlock(*CoinbaseAcc, CoinbasePwd)
	Deploy()
}

func Deploy() error {
	auth := DefaultTransactOptions()
	a, t, d, err := contracts.DeployContracts(auth, Client, Bank, Insurance)
	DatabaseAddr = a
	Database = d
	_ = t
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	time.Sleep(250 * time.Millisecond)

	return err
}
