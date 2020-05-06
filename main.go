package main

import (
	"fmt"
	"strings"
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

	fmt.Printf("%d\n" , Question38(84, 60))
}

func Question38(a , b int) int{
	if b == 0 {
		return a
	} else {
		return Question38(b, a % b)
	}
}


func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}