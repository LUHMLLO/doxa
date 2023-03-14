package lib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, true, ClientURL, "POST")

	createUserReq := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		// set error status code and message
		http.Error(w, "error reading request body", 200)
		return
	}

	if createUserReq.Username == "" {
		// set error status code and message
		http.Error(w, "username field was empty", 400)
		return
	}

	if createUserReq.Password == "" {
		// set error status code and message
		http.Error(w, "password field was empty", 400)
		return
	}

	if createUserReq.Avatar == "" {
		// set error status code and message
		http.Error(w, "avatar field was empty", 400)
		return
	}

	if createUserReq.Name == "" {
		// set error status code and message
		http.Error(w, "name field was empty", 400)
		return
	}

	if createUserReq.Email == "" {
		// set error status code and message
		http.Error(w, "email field was empty", 400)
		return
	}

	if createUserReq.Phone == "" {
		// set error status code and message
		http.Error(w, "phone field was empty", 400)
		return
	}

	if createUserReq.Role == "" {
		// set error status code and message
		http.Error(w, "role field was empty", 400)
		return
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

	_, err = s.store.users_beforeInsert(user)
	if err != nil {
		// set error status code and message
		http.Error(w, err.Error(), 400)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		// set error status code and message
		http.Error(w, fmt.Sprintf("error hashing password: %v", err), 400)
		return
	}

	err = s.store.users_insert(user)
	if err != nil {
		// set error status code and message
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, true, ClientURL, "POST")

	signInUserReq := &SigninUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&signInUserReq); err != nil {
		// set error status code and message
		http.Error(w, "error reading request body", 200)
		return
	}

	if signInUserReq.Username == "" {
		// set error status code and message
		http.Error(w, "username field was empty", 400)
		return
	}

	if signInUserReq.Password == "" {
		// set error status code and message
		http.Error(w, "password field was empty", 400)
		return
	}

	user, err := s.store.Users_beforeSignin(signInUserReq.Username, signInUserReq.Password)
	if err != nil {
		// set error status code and message
		http.Error(w, err.Error(), 400)
		return
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		// set error status code and message
		http.Error(w, err.Error(), 400)
		return
	}

	// if err = s.store.users_update(user.ID, "modified", time.Now().UTC()); err != nil {
	// 	json.NewEncoder(w).Encode(fmt.Sprintf("error updating modified: %v", err))
	// 	return
	// }

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode("user session initialized")
}

func (s *Server) SignedUser(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, true, ClientURL, "GET")

	//log.Println("after authenticated : ", r.Context().Value(ClaimsContext).(httpClaimsContext))

	IssuerUUID, err := uuid.Parse(r.Context().Value(ClaimsContext).(httpClaimsContext).Issuer)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error formating issuer: %v", err))
		return
	}

	user, err := s.store.users_read(IssuerUUID)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("token was valid but issuer was not found: %v", err))
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) SignOut(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, true, ClientURL, "POST")

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode("user session terminated")
}

func (s *Server) UserDevices(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, true, ClientURL, "GET")

	devices, err := s.store.devices_readTableWhereOwner(r.Context().Value(ClaimsContext).(httpClaimsContext).Issuer)
	if err != nil {
		// set error status code and message
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(devices)
}
