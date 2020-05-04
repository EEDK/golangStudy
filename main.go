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
	Question36(3 , 1, 1 , 1)
	fmt.Println(cnt)
}

var cnt int = 0
var checked = [21][21][21]int{}

func Question36(n int , a int , b int , c int) {

	if a + b + c == n {
		if a <= b && a <= c && a + b > c && checked[a][b][c] == 0 {
			cnt = cnt + 1
			checked[a][b][c] = 1
		}
		return
	}
	Question36(n, a+1 , b , c)
	Question36(n, a , b+1 , c)
	Question36(n, a , b , c+1)
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}