package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type TripleTuple struct {
	a, b, c interface{}
}

func main() {
	// dictionaryURL := "https://raw.githubusercontent.com/zeisler/scrabble/master/db/dictionary.csv"
	// airportCodeURL := "https://raw.githubusercontent.com/datasets/airport-codes/master/data/airport-codes.csv"
	// dictionaryPath := download_csv(dictionaryURL)
	// airportCodePath := download_csv(airportCodeURL)
	// six_letter_words := filter_six_letter_words(dictionaryPath)
	// airport_codes := get_airport_codes(airportCodePath)

	six_letter_words := filter_six_letter_words("data/dictionary.csv")
	airport_codes := get_airport_codes("data/airport-codes.csv")

	find_combinations(six_letter_words, airport_codes)

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

func filter_six_letter_words(filePath string) []string {
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

	/*
		for _, word := range six_letter_words {
			fmt.Println(word)
		}
	*/

	return six_letter_words
}

func get_airport_codes(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading records", err)
	}

	local_code_index := -1
	for column := range records[0] {
		if records[0][column] == "local_code" {
			local_code_index = column
			break
		}
	}

	if local_code_index == -1 {
		log.Fatal("Could not find column for local airport code")
	}

	r := regexp.MustCompile("^[[:alpha:]][[:alpha:]][[:alpha:]]$")

	airport_codes := []string{}
	for _, eachrecord := range records[1:] {
		if r.Match([]byte(eachrecord[local_code_index])) {
			airport_codes = append(airport_codes, eachrecord[local_code_index])
		}
	}

	/*
		for _, code := range airport_codes {
			fmt.Println(code)
		}
	*/

	return airport_codes
}

func find_combinations(six_letter_words []string, airport_codes []string) {
	created_words := []TripleTuple{}
	for _, code_one := range airport_codes {
		for _, code_two := range airport_codes {
			concatenated_code := code_one + code_two
			for _, word := range six_letter_words {
				if concatenated_code == word {
					created_words = append(created_words, TripleTuple{concatenated_code, code_one, code_two})
				}
			}
		}
	}

	for _, tuple := range created_words {
		fmt.Printf("%s + %s = %s\n", tuple.b, tuple.c, tuple.a)
	}
}
