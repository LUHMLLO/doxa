package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	s.Route_insert_user(w, r)
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
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading request body: %v", err))
		return
	}

	user, err := s.store.Users_beforeSignin(signInUserReq.Username, signInUserReq.Password)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error before signin: %v", err))
		return
	}

	token, err := GenerateJWT(user.Username, user.Role)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error generating token: %v", err))
		return
	}

	if err = s.store.users_update(user.ID, "jwt", token); err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error updating jwt: %v", err))
		return
	}

	if err = s.store.users_update(user.ID, "modified", time.Now().UTC()); err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error updating modified: %v", err))
		return
	}

	cookie := &http.Cookie{
		Name:     "JWT",
		Value:    token,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode(user)
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

func (s *Server) CheckUsername(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	err := s.store.users_readCol("name", r.Header.Get("Username"))
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading column: %v", err))
		return
	}

	json.NewEncoder(w).Encode(r.Header.Get("Username"))
}
