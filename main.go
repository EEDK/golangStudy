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
	fmt.Printf("%d\n" , Question31(17))

}

func Question31(n int) int{
	answer := -1
	for i := 0 ; n >= 0 ; i++ {
		n = n - 2 * i + 1

		if n < 0 {
			answer = i-1
		}
	}
	return answer
}