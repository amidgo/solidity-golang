package routing

import (
	"io"
	"net/http"
	"os"
	"text/template"

	"github.com/amidgo/solidity-ethereum/database"
)

func TemplatePath(fileName string) string {
	return "/home/amidman/.go/src/github.com/amidgo/solidity-golang/templates/" + fileName
}

func NewTemplate(fileName string) (*template.Template, error) {
	file, _ := os.ReadFile(
		TemplatePath(fileName),
	)
	fileString := string(file)
	return template.New(
		"data",
	).Parse(fileString)
}

func Configure() {
	http.HandleFunc("/info", database.Info)
	http.HandleFunc("/renew", database.RenewInsurance)
	http.HandleFunc("/add-dtp", database.AddDtp)
	http.HandleFunc("/add-fine", database.AddFine)
	http.HandleFunc("/buy-insurance", database.BuyInsuranse)
	http.HandleFunc("/calc-insurance", database.CalcInsurance)
	http.HandleFunc("/pay-fine", database.PayFine)
	http.HandleFunc("/add-document", database.AddDocument)
	http.HandleFunc("/add-car", database.AddCar)
	http.HandleFunc("/register", database.RegisterDriver)
	http.HandleFunc("/login", database.Login)
	http.HandleFunc("/balance", database.Balance)
	http.HandleFunc("/test", Test)
	http.ListenAndServe(":8888", nil)
}

func Test(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(body)
}
