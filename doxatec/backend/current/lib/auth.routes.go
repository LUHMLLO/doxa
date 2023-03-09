package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (s *Server) SignUp(w http.ResponseWriter, r *http.Request) {
	s.Route_insert_user(w, r)
}

func (s *Server) SignIn(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, true, ClientURL, "POST")

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

	token, err := GenerateJWT(user.ID.String())
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error generating token: %v", err))
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
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode("user session initialized")
}

func (s *Server) SignedUser(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, true, ClientURL, "GET")

	// for _, cookie := range r.Cookies() {
	// 	fmt.Printf("Cookie: %s=%s\n", cookie.Name, cookie.Value)
	// }

	cookie, err := r.Cookie("jwt")
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("token does not exists: %v", err))
		return
	}

	//log.Println(cookie)

	if cookie.Value == "" {
		log.Println("received an empty cookie")
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretJWTkey), nil
	})
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("not authenticated: %v", err))
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	IssuerUUID, err := uuid.Parse(claims.Issuer)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("error formating issuer: %v", err))
		return
	}

	//log.Println(IssuerUUID)

	user, err := s.store.users_read(IssuerUUID)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("token was valid but issuer was not found: %v", err))
		return
	}

	//log.Println(user)

	json.NewEncoder(w).Encode(user)
}

func (s *Server) SignOut(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w, true, ClientURL, "POST")

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode("user session terminated")
}
