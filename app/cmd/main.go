package main

import (
	"github.com/gofiber/fiber/v2"
	//burada olusturdugumuz docs paketini underscore ile import ediyoruz. Bunu yapma amacimiz docs icerisindeki docs.go dosyasini main ile etkilesime gecirmek. Eger bu islemi yapmazsak init islemleri gerceklesmeyecek.
	_"github.com/omrfrkazt/go-fiber-swagger/app/cmd/docs"
)


// @title Fiber-Swagger API
// @version 0.99dev
// @description Fiber-Swagger API

//ALT KISIMDA AUTHORIZATON EKLEDIK

//BasicAuth kismini istenildigi gibi degistirilebilir
// @securityDefinitions.apikey BasicAuth
// @in header
// @name Authorization
func main() {
	//default bir fiber app olusturuyoruz
	app := fiber.New()
	//routelarimizi olusturuyoruz
	Routes(app)
	//5454 portunda listen ediyoruz - ayaga kaldiriyoruz
	app.Listen(":5454")
}

//NOT: SWAGGERDA HERHANGI BIR DEGISIKLIK OLDUGU ZAMAN MAIN.GO DOSYASININ BULUNDUGU DIZINDE BULUNAN DOCS DOSYASI SILINDIKTEN SONRA TERMINAL ACARAK SWAG INIT YAZILARAK GUNCELLENMESI GEREKMEKTEDIR.

//swagger dokuman linki;
//https://pkg.go.dev/github.com/gofiber/swagger@v0.0.1