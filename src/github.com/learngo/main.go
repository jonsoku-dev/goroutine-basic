package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// dto인가보다
type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	// main에 있는 jobs은 []extractedJob이 모여있는 []extractedJob이다.
	// 그러므로 아래서 spread문법으로 풀어준다.
	var jobs []extractedJob
	totalPages := getPages()

	// 각각의 페이지를 탐색한다. (페이지네이션에서 페이지를 얻은만큼 !)
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		// 배열 안에 배열이 들어가있으면 안되므로 풀어서 넣는다.
		jobs = append(jobs, extractedJobs...)
	}

	fmt.Println(jobs)
}

func getPage(i int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&start=" + strconv.Itoa(i*50)
	fmt.Println("Requesting...", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// Card 전체를 찾는다. Find는 복수로 가져오나보다. (*Selection)
	searchCards := doc.Find(".jobsearch-SerpJobCard")

	// job 정보를 추출한다.
	searchCards.Each(func(i int, card *goquery.Selection) {
		// extractedJob struct
		job := extractJob(card)
		// []extractedJob struct에 개별 extractedJob struct를 주입
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find("salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	// struct 에 주입하여 리턴한다.
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		// page anchor size = 총 페이지 개수
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", res.StatusCode)
	}
}
