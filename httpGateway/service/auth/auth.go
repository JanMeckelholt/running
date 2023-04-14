package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey []byte

const RunningAppHttpClient = "running_app"
const MasterHttpClient = "master"

var validHttpClientNames = []string{MasterHttpClient, RunningAppHttpClient}

var httpClients = map[string]string{}

func UpsertHttpClients(httpClient, password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password is too short")
	}
	for _, validClient := range validHttpClientNames {
		if validClient == httpClient {
			httpClients[httpClient] = password
			return nil
		}
	}
	return fmt.Errorf("httpClient does not exist %s", httpClient)
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func LoginHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		log.Infof("trying to login")
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			log.Infof("failing...%s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		expectedPassword, ok := httpClients[creds.Username]

		if !ok || expectedPassword != creds.Password {
			log.Infof("not authorized...")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("{\"login\":\"failed - not authorized\"}")
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &Claims{
			Username: creds.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(JwtKey)
		if err != nil {
			log.Infof("internal error...%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("{\"login\":\"failed - internal error\"}")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteStrictMode,
		})
		log.Info("Login Success!")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("{\"login\":\"success\"}")
	})
}
