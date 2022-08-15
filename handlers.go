package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (app *App) Login(c *fiber.Ctx) error {
	var model UserModel
	c.BodyParser(&model)
	if model.Username == "" || model.Password == "" {
		return app.failure(c, "Email and Password are required")
	}
	cacheGet, _ := app.cache.Get("userList")
	var userList []UserModel
	json.Unmarshal(cacheGet, &userList)
	for _, user := range userList {
		if user.Username == model.Username && user.Password == model.Password {
			key := app.secretGenerator()
			cachedSecrets,_:= app.cache.Get("secrets")
			var secretsList []SecretModel
			json.Unmarshal(cachedSecrets, &secretsList)
			fmt.Println(secretsList)
			for i, secret := range secretsList {
				if secret.UserName == user.Username {
					secretsList = append(secretsList[:i], secretsList[i+1:]...)
					break
				}
			}
			secretsList = append(secretsList, SecretModel{
				UserName: user.Username,
				Key:      app.secretGenerator(),
			})
			var cacheSet, _ = json.Marshal(secretsList)
			app.cache.Set("secrets", cacheSet)
			return app.success(c, map[string]string{
				"key": key,
			})
		}
	}
	return app.failure(c, "Email or Password is incorrect")
}

func (app *App) Logout(c *fiber.Ctx) error {
	key := c.GetReqHeaders()["Authorization"]
	app.cache.Delete(key)
	return app.success(c, "Logout Success")
}

func (app *App) Register(c *fiber.Ctx) error {
	var bodyModel UserModel
	c.BodyParser(&bodyModel)
	if bodyModel.Username == "" || bodyModel.Password == "" {
		return app.failure(c, "Email and Password are required")
	}
	cacheGet, _ := app.cache.Get("userList")
	var userList []UserModel
	json.Unmarshal(cacheGet, &userList)
	for _, user := range userList {
		if user.Username == bodyModel.Username {
			return app.failure(c, "Email or Username already exists")
		}
	}
	userList = append(userList, bodyModel)
	var cacheSet, _ = json.Marshal(userList)
	app.cache.Set("userList", cacheSet)
	return app.success(c,nil)
}

func (app *App) Dashboard(c *fiber.Ctx) error {
	return app.success(c, "You accessed dashboard, golang is awesome!!!!!!!!")
}

