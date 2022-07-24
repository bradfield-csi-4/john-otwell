package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

const baseUrl = "https://xkcd.com"
const suffix = "/info.0.json"
const indexDirName = "index"

type XKCDResponse struct {
	Alt        string `json:"alt"`
	Day        string `json:"day"`
	Img        string `json:"img"`
	Link       string `json:"link"`
	Month      string `json:"month"`
	News       string `json:"news"`
	Num        uint   `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Year       string `json:"year"`
}

func buildIndex() {
	err := os.Mkdir(indexDirName, 0777)
	if err != nil {
		log.Fatal(err)
	}
	updateIndex(1, 5)
}

func updateIndex(start uint, end uint) {
	var wg sync.WaitGroup
	for i := start; i <= end; i++ {
		filename := fmt.Sprintf("%s/%d.json", indexDirName, i)
		url := fmt.Sprintf("%s/%d", baseUrl, i)

		wg.Add(1)

		go func() {
			defer wg.Done()
			writeToFileFromUrl(filename, url)
		}()
	}

	wg.Wait()
}

func indexExists() bool {
	_, err := os.Stat(indexDirName)
	if err != nil {
		return false
	} else {
		return true
	}
}

func getCurrent() uint {
	xkcdata := urlToStruct(baseUrl)
	return xkcdata.Num
}

func urlToStruct(url string) XKCDResponse {
	var xkcdata XKCDResponse
	body := getABody(url)
	json.Unmarshal(body, &xkcdata)
	return xkcdata
}

func getABody(url string) []byte {
	res, httpErr := http.Get(url + suffix)
	if httpErr != nil {
		log.Fatal(httpErr)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func writeToFileFromUrl(filename string, url string) {
	body := getABody(url)
	os.WriteFile(filename, body, 0777)
}
