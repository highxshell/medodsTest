package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"medodsTest/pkg/store/mongodb"
	http2 "medodsTest/services/jwt/internal/delivery/http"
	"medodsTest/services/jwt/usecase/jwt"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("failed to load env")
	}
	portStr := os.Getenv("PORT")
	port, _ := strconv.Atoi(portStr)
	storage, err := mongodb.New()
	if err != nil {
		slog.Error("failed to connect to db")
		os.Exit(1)
	}
	coll := storage.OpenCollection("token", "refresh")
	jwtUC := jwt.NewUseCaseImpl(coll)
	delivery := http2.NewJWTHTTP(jwtUC)
	delivery.Start(port)
}
