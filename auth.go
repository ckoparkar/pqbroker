package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Credentials struct {
	Username string `json:username`
	Password string `json:password`
}

func credentials() Credentials {
	data, err := Asset("config/auth.json")
	if err != nil {
		panic(err)
	}
	var creds Credentials
	json.Unmarshal(data, &creds)
	return creds
}

func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const basicAuthPrefix string = "Basic "
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, basicAuthPrefix) {
			actual := credentials()
			payload, err := base64.StdEncoding.DecodeString(auth[len(basicAuthPrefix):])
			if err == nil {
				got := strings.Split(string(payload), ":")
				if actual.Username == got[0] && actual.Password == got[1] {
					h(w, r, ps)
					return
				}
			}
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
