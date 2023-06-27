// Package main handles the middleware for the web app
package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/joho/godotenv"
)

func getAuthMiddleware() func(*fiber.Ctx) error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")

	authMiddleware := keyauth.New(keyauth.Config{
		Validator: func(c *fiber.Ctx, key string) (bool, error) {
			hashedAPIKey := sha256.Sum256([]byte(apiKey))
			hashedKey := sha256.Sum256([]byte(key))

			if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
				return true, nil
			}
			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	})

	return authMiddleware
}
