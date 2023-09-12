package main

import (
	"bufio"
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
type QuadTuple struct {
	a, b, c, d interface{}
}
type QuintTuple struct {
	a, b, c, d, e interface{}
}

func main() {
	dictionaryURL := "https://raw.githubusercontent.com/zeisler/scrabble/master/db/dictionary.csv"
	airportCodeURL := "https://raw.githubusercontent.com/datasets/airport-codes/master/data/airport-codes.csv"
	dictionaryPath := download_csv(dictionaryURL)
	airportCodePath := download_csv(airportCodeURL)
	six_letter_words := filter_x_letter_words(dictionaryPath, 6)
	nine_letter_words := filter_x_letter_words(dictionaryPath, 9)
	twelve_letter_words := filter_x_letter_words(dictionaryPath, 12)
	airport_codes := get_airport_codes(airportCodePath)

	//six_letter_words := filter_x_letter_words("data/dictionary.csv", 6)
	//nine_letter_words := filter_x_letter_words("data/dictionary.csv", 9)
	//twelve_letter_words := filter_x_letter_words("data/dictionary.csv", 12)
	//airport_codes := get_airport_codes("data/airport-codes.csv")

	find_combinations_brute_force(six_letter_words, nine_letter_words, twelve_letter_words, airport_codes)

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

func filter_x_letter_words(filePath string, x int) []string {
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
		if len(eachrecord[0]) == x {
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

func find_combinations_brute_force(six_letter_words []string, nine_letter_words []string, twelve_letter_words []string, airport_codes []string) {
	created_six_letter_words := []TripleTuple{}
	created_nine_letter_words := []QuadTuple{}
	created_twelve_letter_words := []QuintTuple{}
	for _, code_one := range airport_codes {
		for _, code_two := range airport_codes {
			concatenated_code_one := code_one + code_two
			for _, word := range six_letter_words {
				if strings.EqualFold(concatenated_code_one, word) {
					created_six_letter_words = append(created_six_letter_words, TripleTuple{concatenated_code_one, code_one, code_two})
				}
			}
			for _, code_three := range airport_codes {
				concatenated_code_two := concatenated_code_one + code_three
				for _, word := range nine_letter_words {
					if strings.EqualFold(concatenated_code_two, word) {
						created_nine_letter_words = append(created_nine_letter_words, QuadTuple{concatenated_code_two, code_one, code_two, code_three})
					}
				}

				for _, code_four := range airport_codes {
					concatenated_code_three := concatenated_code_two + code_four
					for _, word := range twelve_letter_words {
						if strings.EqualFold(concatenated_code_three, word) {
							created_twelve_letter_words = append(created_twelve_letter_words, QuintTuple{concatenated_code_three, code_one, code_two, code_three, code_four})
						}
					}

				}
			}
		}
	}

	to_write_six := []string{}
	to_write_nine := []string{}
	to_write_twelve := []string{}

	for _, tuple := range created_six_letter_words {
		combination_code := fmt.Sprintf("%s + %s = %s\n", tuple.b, tuple.c, tuple.a)
		to_write_six = append(to_write_six, combination_code)
		fmt.Printf("%s + %s = %s\n", tuple.b, tuple.c, tuple.a)
	}
	writeLines(to_write_six, "data/results_6.txt")

	for _, tuple := range created_nine_letter_words {
		combination_code := fmt.Sprintf("%s + %s + %s = %s\n", tuple.b, tuple.c, tuple.d, tuple.a)
		to_write_nine = append(to_write_nine, combination_code)
		fmt.Printf("%s + %s + %s = %s\n", tuple.b, tuple.c, tuple.d, tuple.a)
	}
	writeLines(to_write_six, "data/results_9.txt")

	for _, tuple := range created_twelve_letter_words {
		combination_code := fmt.Sprintf("%s + %s + %s + %s = %s\n", tuple.b, tuple.c, tuple.d, tuple.e, tuple.a)
		to_write_twelve = append(to_write_twelve, combination_code)
		fmt.Printf("%s + %s + %s + %s = %s\n", tuple.b, tuple.c, tuple.d, tuple.e, tuple.a)
	}
	writeLines(to_write_six, "data/results_12.txt")
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}
