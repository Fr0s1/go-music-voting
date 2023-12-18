package main

import (
	"encoding/json"
	"log"
	database "user-service/pkg/db/mysql"

	"net/http"
	"user-service/pkg/jwt"
	"user-service/pkg/users"
)

type AuthResult struct {
	Token        string `json:"token"`
	ErrorMessage string `json:"error"`
}

func signup(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user users.User
	err := decoder.Decode(&user)

	var result AuthResult

	w.Header().Set("Content-Type", "application/json")

	var response []byte

	if err != nil {
		result.ErrorMessage = err.Error()
		response, _ = json.Marshal(&result)

		http.Error(w, string(response), http.StatusBadRequest)

		return
	}

	err = user.Save()

	if err != nil {
		result.ErrorMessage = err.Error()
		response, _ = json.Marshal(&result)

		http.Error(w, string(response), http.StatusBadRequest)

		return
	}

	token, err := jwt.GenerateToken(user.Username)

	result = AuthResult{
		Token:        token,
		ErrorMessage: "",
	}

	response, _ = json.Marshal(&result)

	w.Write(response)
}

func signin(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user users.User
	err := decoder.Decode(&user)

	var result AuthResult

	w.Header().Set("Content-Type", "application/json")

	var response []byte

	if err != nil {
		result.ErrorMessage = err.Error()
		response, _ = json.Marshal(&result)

		http.Error(w, string(response), http.StatusBadRequest)

		return
	}

	isValidUser, err := user.Authenticate()

	if err != nil {
		result.ErrorMessage = err.Error()
		response, _ = json.Marshal(&result)

		http.Error(w, string(response), http.StatusBadRequest)

		return
	}

	if isValidUser {
		token, _ := jwt.GenerateToken(user.Username)

		result = AuthResult{
			Token:        token,
			ErrorMessage: "",
		}

		response, _ = json.Marshal(&result)

		w.Write(response)

		return
	} else {
		result = AuthResult{
			Token:        "",
			ErrorMessage: "Wrong username or password",
		}

		response, _ = json.Marshal(&result)

		w.Write(response)

		return
	}
}

func main() {
	database.InitDB()

	defer database.CloseDB()

	http.HandleFunc("/signup/", signup)

	http.HandleFunc("/signin/", signin)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
