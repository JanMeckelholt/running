package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/JanMeckelholt/running/common/models"
	"github.com/JanMeckelholt/running/common/utils"
	log "github.com/sirupsen/logrus"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey []byte

const RunningAppHttpClient = "running_app"
const MasterHttpClient = "master"

var LatestWebsitePing string

var validHttpClientNames = []string{MasterHttpClient}
var validRunningAppNames = []string{RunningAppHttpClient}

var httpClients = map[string]string{}
var runningAppClients = map[string]string{}

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

func UpsertRunningClients(runningAppClient, password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password is too short")
	}
	for _, validClient := range validRunningAppNames {
		if validClient == runningAppClient {
			runningAppClients[runningAppClient] = password
			return nil
		}
	}
	return fmt.Errorf("runningAppClient does not exist %s", runningAppClient)
}

func LoginHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var creds models.Credentials
		log.Infof("trying to login")
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			log.Infof("failing...%s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		expectedPassword, ok := httpClients[*creds.Username]

		if !ok || expectedPassword != *creds.Password {
			log.Infof("not authorized...")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&models.LoginResponse{Login: "failed - not authorized"})
			return
		}

		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &models.Claims{
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
			json.NewEncoder(w).Encode(&models.LoginResponse{Login: "failed - internal error"})
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteNoneMode,
		})
		log.Info("Login Success!")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&models.LoginResponse{Login: "success"})
	})
}

func WebsiteHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			LatestWebsitePing = utils.RandStr(100)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&models.WebsiteResponse{LatestWebsitePing: &LatestWebsitePing})

		case http.MethodPost:
			var creds models.CredentialsWebsite
			log.Infof("trying to verify website")
			err := json.NewDecoder(r.Body).Decode(&creds)
			if err != nil {
				log.Infof("failing...%s", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Infof("creds: %s - %s", *creds.Username, *creds.LatestPingEncrypted)

			latestPingEncrypted := creds.LatestPingEncrypted
			log.Infof("original: %s", LatestWebsitePing)
			d := utils.Decrypt(runningAppClients[*creds.Username], *creds.LatestPingEncrypted)
			log.Infof("d: %s", d)

			if latestPingEncrypted == nil || d != LatestWebsitePing {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(&models.LoginResponse{Login: "failed - not authorized"})
				return
			}

			expirationTime := time.Now().Add(5 * time.Minute)

			claims := &models.Claims{
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
				json.NewEncoder(w).Encode(&models.LoginResponse{Login: "failed - internal error"})
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    tokenString,
				Expires:  expirationTime,
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteNoneMode,
			})
			log.Info("Login Success!")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&models.LoginResponse{Login: "success"})
		}
	})
}
