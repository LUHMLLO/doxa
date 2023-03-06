package lib

import (
	"encoding/json"
	"fmt"
	"log"
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
		json.NewEncoder(w).Encode(fmt.Sprintf("error reading body: %s", err.Error()))
		return
	}

	user, err := s.store.Query_before_signin_user(signInUserReq.Username, signInUserReq.Password)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	token, err := GenerateJWT(user.Username, user.Role)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("failed to generate token: %s", err.Error()))
		return
	}

	if err = s.store.Query_update_user_column_where_ID(user.ID, "jwt", token); err != nil {
		log.Fatal(err)
	}

	if err = s.store.Query_update_user_column_where_ID(user.ID, "modified", time.Now().UTC()); err != nil {
		log.Fatal(err)
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
