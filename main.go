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

	Question32()

}

func Question32(){
	for i := 'A' ; i <= 'Z' ; i++{
		for j := 'A' ; j <= 'Z' - (i - 65) ; j++{
			fmt.Printf("%c " , j)
		}
		fmt.Printf("\n")
	}
}