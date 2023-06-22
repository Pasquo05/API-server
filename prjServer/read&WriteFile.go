package prjServer

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ConvertStructToSlideBook(records []Book) [][]string {

	books := [][]string{}

	for _, record := range records {
		book := []string{}
		book = append(book, record.Id)
		book = append(book, record.Title)
		book = append(book, record.Desc)

		books = append(books, book)
	}

	return books

}

func ReadFile(nameFile string) [][]string {

	file, err := os.Open(nameFile)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(records)

	return records

}

func WriteIntoFile(records [][]string) {

	csvFile, err := os.Create("books.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	err = csv.NewWriter(csvFile).WriteAll(records)
	csvFile.Close()

	if err != nil {
		log.Fatal(err)
	}

}

func updateFile(records []Book) {

	book := ConvertStructToSlideBook(records)
	WriteIntoFile(book)

}
