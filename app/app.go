package app

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

var (
	r = fiber.New()
)

func StartApp(chn chan string) {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")
	Routers()
	r.Listen(PORT)
	chn <- "Server running on port 8082"
}
