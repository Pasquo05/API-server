package prjServer

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/books", Wrapper(getBooks, EmptyDecoder)).Methods("GET")
	myRoute.HandleFunc("/book/{id}", Wrapper(getBook, GetKey)).Methods("GET")
	myRoute.HandleFunc("/book/delete/{id}", Wrapper(deleteBook, GetKey)).Methods("GET")
	myRoute.HandleFunc("/book/post", Wrapper(addBook, getBody)).Methods("POST")
	http.Handle("/", myRoute)
}

func Wrapper(fn func(interface{}) (interface{}, error), dec func(*http.Request) (interface{}, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		payload, err := dec(r)
		if err != nil {
			//todo esci
		}

		resp, _ := fn(payload)
		//todo se err ... tornare qualcosa altro

		jsonData, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Scrive il JSON come risposta
		w.Write(jsonData)
	}
}
