package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

const baseUrl = "https://xkcd.com"
const suffix = "/info.0.json"
const indexDirName = "index"
const oneBigJsonName = "big.json"

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

func makeOrUpdate() {
	current := getCurrent()
	if indexExists() {
		files, err := ioutil.ReadDir(indexDirName)
		if err != nil {
			log.Fatal(err)
		}
		latestInIndex := len(files)
		if uint(latestInIndex) < current {
			updateIndex(uint(latestInIndex)+1, current)
		}
	} else {
		buildIndex(current)
	}
}

func buildIndex(current uint) {
	err := os.Mkdir(indexDirName, 0777)
	if err != nil {
		log.Fatal(err)
	}
	updateIndex(1, current)
}

func updateIndex(start uint, end uint) {
	for i := start; i <= end; i += 50 {
		if i+49 > end {
			i = end
		}
		fmt.Printf("Indexing comics %d through %d.\n", i, i+49)
		updateIndexSubtask(i, i+49)
	}
}

func updateIndexSubtask(start uint, end uint) {
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

	if res.StatusCode == 404 {
		fmt.Printf("Didn't find anything at %s. Continuing.\n", url)
	} else if res.StatusCode > 299 {
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

func updateOneBigJson() {
	_, err := os.Stat(indexDirName + "/" + oneBigJsonName)

	if err != nil {
		files, err := ioutil.ReadDir(indexDirName)
		if err != nil {
			log.Fatal(err)
		}
		latestInIndex := len(files)
		updateJSON
	}
}
