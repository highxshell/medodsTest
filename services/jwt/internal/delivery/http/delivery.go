package http

import (
	"encoding/json"
	"fmt"
	"log"
	models "medodsTest/services/jwt/domain/jwt"
	"medodsTest/services/jwt/usecase/jwt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type JWTHTTPDelivery struct {
	jwtUC *jwt.UseCaseImpl
}

func NewJWTHTTP(jwtUC *jwt.UseCaseImpl) *JWTHTTPDelivery {
	return &JWTHTTPDelivery{jwtUC: jwtUC}
}

func (d *JWTHTTPDelivery) Start(port int) {
	addr := fmt.Sprintf(":%d", port)
	http.HandleFunc("/auth/login", d.loginHandler)
	http.HandleFunc("/auth/refresh", d.refreshHandler)
	go func() {
		log.Println("Starting server on port", port)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("%s", err)
		}
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("Shutting down gracefully...")
}

func (d *JWTHTTPDelivery) loginHandler(w http.ResponseWriter, r *http.Request) {
	var guid string
	params := r.URL.Query()
	guid = params.Get("guid")

	if guid == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid credential"})
		return
	}

	token, err := d.jwtUC.CreateToken(guid)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	payload, err := json.Marshal(token)
	if err != nil {
		return
	}
	w.Header().Set("Authorization", token.AccessToken+" "+token.RefreshToken)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (d *JWTHTTPDelivery) refreshHandler(w http.ResponseWriter, r *http.Request) {
	var token models.Token
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	guid, err := d.jwtUC.ValidateRefreshToken(token)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "invalid token"})
		return
	}

	token, err = d.jwtUC.CreateToken(guid)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": "unable to create access token"})
		return
	}
	payload, err := json.Marshal(token)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
