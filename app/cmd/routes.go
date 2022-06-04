package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

//burada fiber uzerinden kullanmak istedigimiz routelari ekledik.
// app.Group yazarak swagger haricindeki servisleri /api konumu altina gruplamis olduk.
// routelari burada olusturdugumuz icin swagger icin de bir adet route olusturduk ve handler kisminda swagger.New() fonksiyonunu cagirarak swaggerin hangi noktada olusmasi gerektigini belirttik. Default config ile olusturduk ama tabi istenirse custom olarak da eklenebiliyor.
func Routes(app *fiber.App) {
	// swagger tam olarak bu noktada ekleniyor.
	app.Get("/swagger/*", swagger.New(swagger.ConfigDefault))
	apiGroup := app.Group("/api")
	apiGroup.Get("/ping", Ping)
	apiGroup.Post("/login", Login)
	apiGroup.Get("/user/:id", GetUserById)
	apiGroup.Get("/secret", SecretHandler)
}
