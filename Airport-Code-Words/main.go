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

type QuintTuple struct {
	a, b, c, d, e interface{}
}

func main() {
	// dictionaryURL := "https://raw.githubusercontent.com/zeisler/scrabble/master/db/dictionary.csv"
	// airportCodeURL := "https://raw.githubusercontent.com/datasets/airport-codes/master/data/airport-codes.csv"
	// dictionaryPath := download_csv(dictionaryURL)
	// airportCodePath := download_csv(airportCodeURL)
	// six_letter_words := filter_x_letter_words(dictionaryPath, 6)
	// nine_letter_words := filter_x_letter_words(dictionaryPath, 9)
	// twelve_letter_words := filter_x_letter_words(dictionaryPath, 12)
	// airport_codes := get_airport_codes(airportCodePath)

	six_letter_words := filter_x_letter_words("data/dictionary.csv", 6)
	nine_letter_words := filter_x_letter_words("data/dictionary.csv", 9)
	twelve_letter_words := filter_x_letter_words("data/dictionary.csv", 12)
	airport_codes := get_airport_codes("data/airport-codes.csv")

	// find_combinations_brute_force(six_letter_words, nine_letter_words, twelve_letter_words, airport_codes)
	find_combinations_prefixing(six_letter_words, nine_letter_words, twelve_letter_words, airport_codes)

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
	created_six_letter_words := []QuintTuple{}
	created_nine_letter_words := []QuintTuple{}
	created_twelve_letter_words := []QuintTuple{}
	for _, code_one := range airport_codes {
		for _, code_two := range airport_codes {
			concatenated_code_one := code_one + code_two
			for _, word := range six_letter_words {
				if strings.EqualFold(concatenated_code_one, word) {
					created_six_letter_words = append(created_six_letter_words, QuintTuple{concatenated_code_one, code_one, code_two, "", ""})
				}
			}
			for _, code_three := range airport_codes {
				concatenated_code_two := concatenated_code_one + code_three
				for _, word := range nine_letter_words {
					if strings.EqualFold(concatenated_code_two, word) {
						created_nine_letter_words = append(created_nine_letter_words, QuintTuple{concatenated_code_two, code_one, code_two, code_three, ""})
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

	output_results_described(created_six_letter_words, created_nine_letter_words, created_twelve_letter_words)
}

func find_combinations_prefixing(six_letter_words []string, nine_letter_words []string, twelve_letter_words []string, airport_codes []string) {

	six_matches := []string{}
	nine_matches := []string{}
	twelve_matches := []string{}

	three_prefix_matches_for_six := []string{}
	three_prefix_matches_for_nine := []string{}
	three_prefix_matches_for_twelve := []string{}
	for _, code_one := range airport_codes {
		three_prefix_matches_for_six = append(three_prefix_matches_for_six, get_infix_matches(six_letter_words, code_one, 3, 0)...)
		three_prefix_matches_for_nine = append(three_prefix_matches_for_nine, get_infix_matches(nine_letter_words, code_one, 3, 0)...)
		three_prefix_matches_for_twelve = append(three_prefix_matches_for_twelve, get_infix_matches(twelve_letter_words, code_one, 3, 0)...)
	}

	six_prefix_matches_for_nine := []string{}
	six_prefix_matches_for_twelve := []string{}
	for _, code_two := range airport_codes {
		six_matches = append(six_matches, get_infix_matches(three_prefix_matches_for_six, code_two, 3, 3)...)
		six_prefix_matches_for_nine = append(six_prefix_matches_for_nine, get_infix_matches(three_prefix_matches_for_nine, code_two, 3, 3)...)
		six_prefix_matches_for_twelve = append(six_prefix_matches_for_twelve, get_infix_matches(three_prefix_matches_for_twelve, code_two, 3, 3)...)
	}

	nine_prefix_matches_for_twelve := []string{}
	for _, code_three := range airport_codes {
		nine_matches = append(nine_matches, get_infix_matches(six_prefix_matches_for_nine, code_three, 3, 6)...)
		nine_prefix_matches_for_twelve = append(nine_prefix_matches_for_twelve, get_infix_matches(six_prefix_matches_for_twelve, code_three, 3, 6)...)
	}

	for _, code_three := range airport_codes {
		twelve_matches = append(twelve_matches, get_infix_matches(nine_prefix_matches_for_twelve, code_three, 3, 9)...)
	}

	output_results(six_matches, nine_matches, twelve_matches)
}

func get_infix_matches(words []string, prefix string, prefix_length int, skip int) []string {
	matches := []string{}
	for _, word := range words {
		word_prefix := n_letters(word, prefix_length, skip)
		if strings.EqualFold(prefix, word_prefix) {
			matches = append(matches, word)
		}
	}

	return matches
}

func output_results_described(results_six []QuintTuple, results_nine []QuintTuple, results_twelve []QuintTuple) {
	to_write_six := []string{}
	to_write_nine := []string{}
	to_write_twelve := []string{}

	for _, tuple := range results_six {
		combination_code := fmt.Sprintf("%s + %s = %s\n", tuple.b, tuple.c, tuple.a)
		to_write_six = append(to_write_six, combination_code)
		fmt.Printf("%s + %s = %s\n", tuple.b, tuple.c, tuple.a)
	}
	write_lines(to_write_six, "results_six.txt")

	for _, tuple := range results_nine {
		combination_code := fmt.Sprintf("%s + %s + %s = %s\n", tuple.b, tuple.c, tuple.d, tuple.a)
		to_write_nine = append(to_write_nine, combination_code)
		fmt.Printf("%s + %s + %s = %s\n", tuple.b, tuple.c, tuple.d, tuple.a)
	}
	write_lines(to_write_nine, "results_nine.txt")

	for _, tuple := range results_twelve {
		combination_code := fmt.Sprintf("%s + %s + %s + %s = %s\n", tuple.b, tuple.c, tuple.d, tuple.e, tuple.a)
		to_write_twelve = append(to_write_twelve, combination_code)
		fmt.Printf("%s + %s + %s + %s = %s\n", tuple.b, tuple.c, tuple.d, tuple.e, tuple.a)
	}
	write_lines(to_write_twelve, "results_twelve.txt")
}

func output_results(results_six []string, results_nine []string, results_twelve []string) {

	write_lines(results_six, "results_six.txt")

	write_lines(results_nine, "results_nine.txt")

	write_lines(results_twelve, "results_twelve.txt")
}

func write_lines(lines []string, path string) error {
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

func n_letters(s string, n int, k int) string {
	i := 0
	for j := range s {
		if i == (n + k) {
			return s[k:j]
		}
		i++
	}
	return s[k:]
}
