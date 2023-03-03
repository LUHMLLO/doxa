package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	s.Handle_insertUsers(w, r)
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	signInUserReq := &SigninUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&signInUserReq)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading body: %s", err.Error()))
		return
	}

	signinUser := NewSigninUser(
		signInUserReq.Username,
		signInUserReq.Password,
	)

	databaseUser, err := s.store.Query_beforeSigninUsers(signinUser)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	validToken, err := GenerateJWT(databaseUser.Username, databaseUser.Role)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("failed to generate token: %s", err.Error()))
		return
	}

	secretJWTtoken := NewSecretJWTtoken(
		databaseUser.Username,
		databaseUser.Role,
		validToken,
	)
	json.NewEncoder(w).Encode(secretJWTtoken)
}

func (s *Server) CheckAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	switch r.Header.Get("Role") {
	case "admin":
		json.NewEncoder(w).Encode("Welcome, Admin.")
		return
	default:
		json.NewEncoder(w).Encode("Not authorized.")
		return
	}
}

func (s *Server) CheckRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	switch r.Header.Get("Role") {
	case "admin":
		json.NewEncoder(w).Encode(fmt.Sprintf("Role: %s", r.Header.Get("Role")))
		return
	case "user":
		json.NewEncoder(w).Encode(fmt.Sprintf("Role: %s", r.Header.Get("Role")))
		return
	default:
		json.NewEncoder(w).Encode("Not authorized.")
		return
	}
}
