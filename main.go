package main

import (
	"fmt"
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

	algorithm.Question29(3, "a", "b", "c")
}

func Question29(n int , a string , b string , c string){
	if n > 0 {
		Question29(n-1, a , c , b)
		fmt.Printf("%d번 원판을 %s에서 %s 로 옮김\n", n , a , b)
		Question29(n-1, c , b , a)
	}
}