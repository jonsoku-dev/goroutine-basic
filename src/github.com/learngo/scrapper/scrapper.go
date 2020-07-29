package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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

func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50"

	// main에 있는 jobs은 []extractedJob이 모여있는 []extractedJob이다.
	// 그러므로 아래서 spread문법으로 풀어준다.
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)

	// 각각의 페이지를 탐색한다. (페이지네이션에서 페이지를 얻은만큼 !)
	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, c)
	}

	for i := 0; i < totalPages; i++ {
		// 여기에서의 c는 채널들의 집합인 (즉, 배열)이 들어오기때문에
		extractedJob := <-c
		// 여기서 ...로 풀어줘야한다.
		jobs = append(jobs, extractedJob...)
	}

	writeJobs(jobs)
	fmt.Println("파일 작성 완료 🔥 ", len(jobs))
}

// page ( 0 ~ lastLnegth )
func getPage(page int, url string, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := url + "&start=" + strconv.Itoa(page*50)
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
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		// card의 개수마다 요게 반복해서 채널에 콕콕찔렸으니 그 채널을 job에 담고
		job := <-c
		// job에 담긴 채널을 jobs에 담는다. 즉, jobs는 채널들의 배열이라해야되나?
		jobs = append(jobs, job)
	}

	// 채널이 담긴 jobs를 mainC채널에 담는다. 즉 병렬을 또 병렬에 담는다
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find("salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())

	// card의 개수마다 요게 반복해서 채널에 콕콕!
	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
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

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	// 함수가 끝날때 저장하여 파일을 생성한다.
	defer w.Flush()

	headers := []string{"ID", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	/*
		전체 jobs를 순회하면서 map을 하나 만든다.
		map안에는 id, title, location, salary, summary 순으로 작성을 하고
		해당하는 jobSlice를  csv파일에 한줄씩 저장한다 (나중에 이곳도 채널링해야할듯)
	*/
	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.location, job.salary, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
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
