package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DNS        string
	SECRET     string
	PORT       string
	ORDERURL   string
	PRODUCTURL string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env")
	}
}

func SetupConfig() (AppConfig, error) {
	dns := os.Getenv("DNS")
	if len(dns) <= 0 {
		return AppConfig{}, errors.New("couldnt load dns")
	}

	secret := os.Getenv("SECRET")
	if len(secret) <= 0 {
		return AppConfig{}, errors.New("couldnt load secret")
	}
	port := os.Getenv("PORT")
	if len(port) <= 0 {
		return AppConfig{}, errors.New("couldnt load port")
	}
	orderUrl := os.Getenv("ORDERURL")
	if len(orderUrl) <= 0 {
		return AppConfig{}, errors.New("couldnt load orderUrl")
	}
	productUrl := os.Getenv("PRODUCTURL")
	if len(productUrl) <= 0 {
		return AppConfig{}, errors.New("couldnt load productUrl")
	}
	return AppConfig{DNS: dns, SECRET: secret, PORT: port, ORDERURL: orderUrl, PRODUCTURL: productUrl}, nil
}
