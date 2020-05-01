package main

import "fmt"

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

	Question33('Z')
}

func Question33(n int32){
	if n < 'A' || n > 'Z' {

	} else {
		for c := 'A'; c <= n ; c++{
			fmt.Printf("%2c " , c)
		}
		fmt.Printf("\n")
	}
	next := n - 1
	Question33(next)
}