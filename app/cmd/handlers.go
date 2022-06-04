package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

//burada handlerlari olusturuyoruz.

//asagidaki parametreler swagger parametreleridir. swagger olusturulurken bu commentler uzerinden generate edecektir
//aciklama yapmaya gerek yok diye dusunuyorum

//@Summary Ping Me!
//@Success 200 "pong"
//@Router /api/Ping [get]
func Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("pong")
}

//@Summary Login
//@Accept json
//@Produce json
//@Param user body LoginModel true "Login"
//@Success 200 "Login Successful"
//@Router /api/Login [post]
func Login(c *fiber.Ctx) error {
	//bodyi almamiz icin bir login modele ihtiyacimiz vardi
	var body LoginModel
	//bodyi yukaridaki modelle parse etmeye calis edemezsen;
	if err := c.BodyParser(&body); err != nil {
		//hata ver
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON")
	}
	//dolu gelsin buralar dolu
	if body.Username == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Username and password are required")
	}
	//sorun yoksa login successful yap.
	return c.Status(fiber.StatusOK).SendString("Login successful")
}

//@Summary Get User By Id
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} UserModel
//@Failure 404 "User Not Found"
//@Router /api/user/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	//id parametresini int olarak almaya calis
	id, err := c.ParamsInt("id")
	//eger bir hata varsa - parse edemiyorsan
	if err != nil {
		//hata ver
		return c.Status(fiber.StatusBadRequest).SendString("Invalid id")
	}
	//user listesini al
	userList := userGenerator()
	//id ye gore useri bul
	for _, user := range userList {
		fmt.Println(user)
		if user.Id == id {
			return c.Status(fiber.StatusOK).JSON(user)
		}
	}
	//bulamazsan bulamamissindir :D
	return c.Status(fiber.StatusNotFound).SendString("User not found")
}

//@Summary Secret
//@Success 200 "Hello Boss"
//@Router /api/Secret [get]
//security tagiyle BasicAuth kullanacagimizi belirttik.
//@Security BasicAuth
func SecretHandler(c *fiber.Ctx) error {
	if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	return c.Status(fiber.StatusOK).SendString("You access me!")
}

////burada user verisi olusturuyoruz.
func userGenerator() []UserModel {
	var userList = make([]UserModel, 100)
	for i := 0; i < 100; i++ {
		userList[i].Id = i
		userList[i].Username = "user" + strconv.Itoa(i)
		userList[i].Password = "password" + strconv.Itoa(i)
		userList[i].Name = "name" + strconv.Itoa(i)
		userList[i].Phone = "530" + strconv.Itoa(i)
		userList[i].Address = "address" + strconv.Itoa(i)
	}
	return userList
}
