package lib

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "POST")

	createUserReq := &CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&createUserReq)
	if err != nil {
		// set error status code and message
		log.Println("error reading request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if createUserReq.Username == "" {
		// set error status code and message
		log.Println("username field was empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if createUserReq.Password == "" {
		// set error status code and message
		log.Println("password field was empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if createUserReq.Avatar == "" {
		// set error status code and message
		log.Println("avatar field was empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if createUserReq.Name == "" {
		// set error status code and message
		log.Println("name field was empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if createUserReq.Email == "" {
		// set error status code and message
		log.Println("email field was empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if createUserReq.Phone == "" {
		// set error status code and message
		log.Println("phone field was empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if createUserReq.Role == "" {
		// set error status code and message
		log.Println("role field was empty")
		w.WriteHeader(http.StatusBadRequest)
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
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		// set error status code and message
		log.Println("error hashing password: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.store.users_insert(user)
	if err != nil {
		// set error status code and message
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "user created succesfully"})
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "POST")

	signInUserReq := &SigninUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&signInUserReq); err != nil {
		// set error status code and message
		log.Println("error reading request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if signInUserReq.Username == "" {
		// set error status code and message
		log.Println("username field was empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if signInUserReq.Password == "" {
		// set error status code and message
		log.Println("password field was empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.store.Users_beforeSignin(signInUserReq.Username, signInUserReq.Password)
	if err != nil {
		// set error status code and message
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		// set error status code and message
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode(map[string]string{"message": "user session initialized"})
}

func (s *Server) SignedUser(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "GET")

	//log.Println("after authenticated : ", r.Context().Value(ClaimsContext).(httpClaimsContext))

	IssuerUUID, err := uuid.Parse(r.Context().Value(ClaimsContext).(httpClaimsContext).Issuer)
	if err != nil {
		log.Println("error formating issuer: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := s.store.users_read(IssuerUUID)
	if err != nil {
		log.Println("token was valid but issuer was not found: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (s *Server) SignOut(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "POST")

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Secure:   false,
		HttpOnly: false,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode(map[string]string{"message": "user session terminated"})
}

func (s *Server) UserDevices(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, r, "GET")

	devices, err := s.store.devices_readTableWhereOwner(r.Context().Value(ClaimsContext).(httpClaimsContext).Issuer)
	if err != nil {
		// set error status code and message
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(devices)
}
