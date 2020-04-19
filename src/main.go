package main

import (
	"fmt"
	"golangStudy/src/algorithm"
	"golangStudy/src/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	/*e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))*/

	for n := 1; n <= 5; n++ {
		for r := 0; r <= n; r++ {
			fmt.Printf("%d C %d = %d\n", n, r, algorithm.Qusetion20(n, r))
		}
	}
}
