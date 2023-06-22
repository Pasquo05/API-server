package prjServer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var Books = []Book{}

func AddBookIntoBooksFromCSV(nameFile string) {

	records := ReadFile(nameFile)
	Books = ConvertSlideToStructBook(records)

}

func ConvertSlideToStructBook(records [][]string) []Book {

	Books := []Book{}

	for _, record := range records {

		book := Book{}

		book.Id = record[0]
		book.Title = record[1]
		book.Desc = record[2]

		Books = append(Books, book)
	}
	return Books
}

func getBooks(_ interface{}) (interface{}, error) {
	return Books, nil
}

func EmptyDecoder(r *http.Request) (interface{}, error) {
	return nil, nil
}

func GetKey(r *http.Request) (interface{}, error) {

	vars := mux.Vars(r)
	key := vars["Id"]

	return key, nil
}

func getBook(key interface{}) (interface{}, error) {

	bookReturn := Book{}

	for _, book := range Books {
		if book.Id == key {
			bookReturn = book
		}
	}

	fmt.Println(bookReturn)

	return bookReturn, nil

}

func deleteBook(key interface{}) (interface{}, error) {

	for index, book := range Books {
		if book.Id == key {

			Books = append(Books[:index], Books[index+1:]...)
		}
	}

	return Books, nil

}

func addBook(bookInput interface{}) (interface{}, error) {

	book, ok := bookInput.(Book)

	if !ok {
		return nil, fmt.Errorf("errore")
	}

	Books = append(Books, book)

	updateFile(Books)

	return Books, nil
}

func getBody(r *http.Request) (interface{}, error) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var book Book
	err = json.Unmarshal(reqBody, &book)
	if err != nil {
		fmt.Println(err)
	}

	return book, nil
}
