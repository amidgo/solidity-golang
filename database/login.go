package database

import (
	"encoding/json"
	"net/http"

	"github.com/amidgo/solidity-ethereum/variables"
	"github.com/ethereum/go-ethereum/common"
)

type LoginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Addr     common.Address `json:"addr"`
	Password string         `json:"password"`
	Role     string         `json:"role"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var lf LoginForm
	err := json.NewDecoder(r.Body).Decode(&lf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	st, err := variables.Database.Logins(variables.DefaultCallOpts(), lf.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if st.Addr == variables.ZeroAddr {
		http.Error(w, "login not found", http.StatusBadRequest)
		return
	}
	a := variables.ImportAccount(st.Addr)
	if err = variables.KeyStore.Unlock(*a, lf.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lr := &LoginResponse{Addr: st.Addr, Password: st.Password, Role: st.Role}
	j, err := json.Marshal(lr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(j)
}
