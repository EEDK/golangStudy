package main

import (
	"golangStudy/src/algorithm"
)

/*const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment("jobs.csv", "jobs.csv")
}
*/
func main() {
	/*e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))*/

	a := []int{75, 25, 6, 73, 43, 46, 31, 13, 60, 90, 5, 43, 35, 65, 100, 28, 83, 95, 35, 45}

	algorithm.Question25(a)
}
