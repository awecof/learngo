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

type extractedJob struct {
	value    string
	title    string
	location string
	sector   string
}

func Scrape(term string) {
	var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=" + term
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(baseURL, i+1, c)
	}
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}
	writeJobs(jobs)
	fmt.Println("Done Extracting:", len(jobs))
}

func getPages(baseURL string) int {
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

func getPage(baseURL string, page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
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
		go extractJob(card, c)
	})
	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	value, _ := card.Attr("value")
	title := card.Find(".area_job>.job_tit>a").Text()
	location := ""
	card.Find(".area_job>.job_condition>span>a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		if location != "" {
			text = " " + s.Text()
		}
		location += text

	})
	sectors := []string{}
	card.Find(".job_sector>b>a").Each(func(i int, s *goquery.Selection) {
		sectors = append(sectors, s.Text())
	})
	card.Find(".job_sector>a").Each(func(i int, s *goquery.Selection) {
		sectors = append(sectors, s.Text())
	})
	sector := strings.Join(sectors, " ")
	c <- extractedJob{
		value:    value,
		title:    title,
		location: location,
		sector:   sector,
	}

}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)
	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	file.Write(utf8bom)
	w := csv.NewWriter(file)
	defer w.Flush()
	headers := []string{"Link", "Title", "Location", "Sector"}
	err = w.Write(headers)
	checkErr(err)
	for _, job := range jobs {
		jobSlice := []string{
			"https://www.saramin.co.kr/zf_user/jobs/relay/view?rec_idx=" + job.value,
			job.title,
			job.location,
			job.sector,
		}
		err = w.Write(jobSlice)
		checkErr(err)
	}
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
