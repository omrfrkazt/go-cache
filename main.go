package main

import (
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	api     *fiber.App
	context *fiber.Ctx
	cache  *bigcache.BigCache
}

func main() {
	var app App
	app.api = fiber.New(fiber.Config{
		AppName: "fiber-api",
	})
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		panic(err)
	}
	cache.Reset()
	app.cache = cache
	app.AddRoutes()
	app.api.Listen(":5000")
}
