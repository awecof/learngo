package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	value    string
	title    string
	location string
	sector   string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=python"

func main() {
	var jobs []extractedJob
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i + 1)
		jobs = append(jobs, extractedJobs...)
	}
	fmt.Println(jobs)
}

func getPage(page int) []extractedJob {
	var jobs []extractedJob
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".item_recruit")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})
	return jobs
}

func extractJob(card *goquery.Selection) extractedJob {
	value, _ := card.Attr("value")
	title := cleanString(card.Find(".area_job>.job_tit>a").Text())
	location := cleanString(card.Find(".area_job>.job_condition>span>a").Text())
	sectors := []string{}
	card.Find(".job_sector>b>a").Each(func(i int, s *goquery.Selection) {
		sectors = append(sectors, s.Text())
	})
	card.Find(".job_sector>a").Each(func(i int, s *goquery.Selection) {
		sectors = append(sectors, s.Text())
	})
	sector := strings.Join(sectors, " ")
	return extractedJob{
		value:    value,
		title:    title,
		location: location,
		sector:   sector,
	}

}

func cleanString(str string) string {
	return strings.Join(strings.Fields(str), " ")
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
		log.Fatalln("request failed with status:", res.StatusCode)
	}
}
