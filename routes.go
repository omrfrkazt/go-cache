package main

func (app *App) AddRoutes() {
	main := app.api.Group("/api")
	main.Post("/register", app.Register)
	main.Post("/login", app.Login)
	main.Post("/logout", app.Logout)
	main.Get("/dashboard", app.checkCacheMiddleware, app.Dashboard)
}
