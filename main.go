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
	Question37()
}
var N int = 12

func Question37(){
	for n := 0 ; n <= N ; n++{
		for r := 0 ; r <= n; r++{
			fmt.Printf("%d ", Qustion19(n , r))
		}
		fmt.Printf("\n")
	}
}


func Qustion19(n int, r int) int {
	p := 1
	for i := 1; i <= r; i++ {
		p = p * (n - i + 1) / i
	}

	return p
}


func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}