package main

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres"
)

// Creates a sessions store using GoFiber session middleware
func NewStore(config postgres.Config) *session.Store {
	storage := postgres.New(config)

	store := session.New(session.Config{
		Storage:        storage,
		Expiration:     time.Hour * 1,
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: "none",
	})

	return store
}
