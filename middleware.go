package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func (app *App) checkCacheMiddleware(c *fiber.Ctx) error {
	key := c.GetReqHeaders()["Authorization"]
	if key == "" {
		return c.Status(401).SendString("Unauthorized")
	}
	cacheKey, _ := app.cache.Get("secrets")
	var secretsList []SecretModel
	json.Unmarshal(cacheKey, &secretsList)
	for _, secret := range secretsList {
		if secret.Key == key {
			return c.Next()
		}
	}
	return c.Status(401).SendString("Unauthorized")
}
