package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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

	Question35("asdsa")
}

func Question35(n string){
	length := utf8.RuneCountInString(n) // 2: 문자열의 실제 길이를 구함

	isRight := true
	for i := 0 ; i < length / 2 ; i++{
		if n[i] != n[length -i - 1]{
			isRight = false
		}
	}
	if isRight{
		fmt.Print("it is palindrome")
	} else {
		fmt.Print("it's not palindrome")
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}