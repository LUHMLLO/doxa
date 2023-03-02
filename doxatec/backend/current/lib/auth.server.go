package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

	createUserReq := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		var err AuthError
		err = SetError(err, "error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		log.Fatal(err)
	}

	user := NewUser(
		createUserReq.Username,
		createUserReq.Password,
		createUserReq.Avatar,
		createUserReq.Name,
		createUserReq.Email,
		createUserReq.Phone,
		createUserReq.Role,
	)

	databaseUserByName, err := s.store.Query_readUsers_Username(user.Username)
	if err != nil {
		log.Fatal(err)
	}

	if databaseUserByName.Username != "" {
		var err AuthError
		err = SetError(err, "username already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	databaseUserByEmail, err := s.store.Query_readUsers_Email(user.Email)
	if err != nil {
		log.Fatal(err)
	}

	if databaseUserByEmail.Email != "" {
		var err AuthError
		err = SetError(err, "email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	s.store.Query_insertUsers(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

	authUserReq := &AuthUser{}
	err := json.NewDecoder(r.Body).Decode(&authUserReq)
	if err != nil {
		var err AuthError
		err = SetError(err, "error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	databaseUser, err := s.store.Query_readUsers_Username(authUserReq.Username)
	if err != nil {
		log.Fatal(err)
	}

	if databaseUser.Username == "" {
		var err AuthError
		err = SetError(err, "username is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := CheckPasswordHash(authUserReq.Password, databaseUser.Password)
	if !check {
		var err AuthError
		err = SetError(err, "password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := GenerateJWT(databaseUser.Username, databaseUser.Role)
	if err != nil {
		var err AuthError
		err = SetError(err, "failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token Token
	token.Username = databaseUser.Username
	token.Role = databaseUser.Role
	token.JWT = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
