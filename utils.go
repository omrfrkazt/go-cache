package main

import (
	"encoding/json"
	"math/rand"
	"time"
	"github.com/gofiber/fiber/v2"
)

func (app *App) secretGenerator() string {
	const letter = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var seed *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	keyArr := make([]byte, 10)
	for i := range keyArr {
		keyArr[i] = letter[seed.Intn(len(letter))]
	}
	return string(keyArr)
}

func (app *App) success(c *fiber.Ctx, data interface{}) error {
	response, _ := json.Marshal(BaseResponseModel{
		StatusCode: 200,
		Message:    "success",
		Data:       data,
	})
	return c.Status(200).Send(response)
}

func (app *App) failure(c *fiber.Ctx, message string) error {
	response, _ := json.Marshal(BaseResponseModel{
		StatusCode: 500,
		Message:    message,
	})
	return c.Status(500).Send(response)
}

func removeIndex[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
