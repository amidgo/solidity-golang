package database

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
)

type DriverDocument struct {
	Owner        common.Address `json:"owner"`
	Number       string         `json:"number"`
	ValidateTime string         `json:"validate_time"`
	Category     Ctg            `json:"category"`
}

var DocumentList = [7]DriverDocument{
	{variables.ZeroAddr, "000", "11.01.2021", A},
	{variables.ZeroAddr, "111", "12.05.2025", B},
	{variables.ZeroAddr, "222", "09.09.2020", C},
	{variables.ZeroAddr, "333", "13.02.2027", A},
	{variables.ZeroAddr, "444", "11.12.2026", B},
	{variables.ZeroAddr, "555", "24.06.2029", C},
	{variables.ZeroAddr, "666", "31.03.2030", A},
}

func IsExist(dc *DriverDocument) bool {
	for _, d := range DocumentList {
		if d.isEquals(dc) {
			return true
		}
	}
	return false
}

func (d *DriverDocument) isEquals(d1 *DriverDocument) bool {
	return d.Category == d1.Category && d.Number == d1.Number && d.ValidateTime == d1.ValidateTime
}

func AddDocument(w http.ResponseWriter, r *http.Request) {
	var dc DriverDocument
	err := json.NewDecoder(r.Body).Decode(&dc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !IsExist(&dc) {
		http.Error(w, "document does not exist", http.StatusBadRequest)
		return
	}
	a := variables.ImportAccount(dc.Owner)
	tOpts, err := variables.NewTransactOptions(*a, variables.Ether(0))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	variables.Database.AddDriverDocument(tOpts, dc.Number, ParseDate(dc.ValidateTime), string(dc.Category))
}

func ParseDate(date string) uint64 {
	s := strings.Split(date, ".")
	day, _ := strconv.ParseInt(s[0], 10, 64)
	mn, _ := strconv.ParseInt(s[1], 10, 64)
	yr, _ := strconv.ParseInt(s[2], 10, 64)
	return uint64(time.Date(int(yr), time.Month(mn), int(day), 0, 0, 0, 0, time.UTC).Unix())
}
