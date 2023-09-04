package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	dictionaryURL := "https://raw.githubusercontent.com/zeisler/scrabble/master/db/dictionary.csv"
	airportCodeURL := "https://raw.githubusercontent.com/datasets/airport-codes/master/data/airport-codes.csv"
	dictionaryPath := download_csv(dictionaryURL)
	_ = download_csv(airportCodeURL)

	filter_six_letter_words(dictionaryPath)

}

func download_csv(fileURL string) string {

	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	path := parsedURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	fmt.Println("Creating file: ", fileName)

	dataFolderName := "data"
	if _, err := os.Stat(dataFolderName); os.IsNotExist(err) {
		err := os.Mkdir(dataFolderName, 0655)
		if err != nil {
			log.Fatal(err)
		}
	}

	fullPath := dataFolderName + "/" + fileName
	file, err := os.Create(fullPath)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)

	return fullPath
}

func filter_six_letter_words(filePath string) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	six_letter_words := []string{}
	for _, eachrecord := range records {
		if len(eachrecord[0]) == 6 {
			six_letter_words = append(six_letter_words, eachrecord[0])
		}
	}

	for _, eachrecord := range records {
		fmt.Println(eachrecord)
	}

}
