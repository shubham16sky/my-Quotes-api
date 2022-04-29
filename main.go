package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type quote struct {
	Title  string
	Author string
}

var quotes []quote = []quote{
	{"Learn to lead", "Sai Vidya Campus"},
	{"You are enough", "Unknown"},
	{"You are great", "Unknown"},
}

func main() {

	//here we are seeding the rand with the time of the system

	rand.Seed(time.Now().Unix())
	api := echo.New()

	api.GET("/quotes", getQuote)

	//here we are making a new get request with randomQuote hadler function
	api.GET("/quotes/random", randomQuote)

	port := os.Getenv("PORT")

	if port == "" {
		port = "32445"
	}

	api.Start(":" + port)

}

func getQuote(c echo.Context) error {
	return c.JSON(http.StatusOK, quotes)
}

//creating hallndler function to generate random quotes from the given array of the quotes

func randomQuote(c echo.Context) error {

	index := rand.Intn(len(quotes))
	return c.JSON(http.StatusOK, quotes[index])
}
