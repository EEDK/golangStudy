package main

import (
	"fmt"
	"golangStudy/src/algorithm"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	//getPages()
	fmt.Println(algorithm.Is_n_prime_number(100))
}

func getPages() int {
	res, err := http.Get(baseURL)
	// 에러 발생시
	checkErr(err)
	// 정상 접속이 안될수
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	//doc.Find(".pagination").Each()

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}

}
