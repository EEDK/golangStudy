package main

import (
	"fmt"
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
	fmt.Printf("%d\n" , Question30(4 , 7))

}

func Question30(a int , b int) int{
	if a == b {
		return 0
	} else if b > a {
		return Question30(a, b / 2) + 1
	} else {
		return Question30(a / 2 , b) + 1
	}
}