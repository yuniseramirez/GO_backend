package main

import (
                "encoding/json"
                "log"
                "net/http"
                "github.com/gorilla/mux"
                "github.com/gorilla/handlers"

)

type Article struct {
                Image string
                Tittle string
                Paragraph string
}

var article = []Article{
                Article{Image: "https://i2.wp.com/blackdoctor.org/wp-content/uploads/2016/09/morris-wife-wiki.png?zoom=2&resize=728%2C388&ssl=1", Tittle: "MORRIS CHESTNUT" , Paragraph: "Actor Morris Chestnut has played opposite a lot of different beautiful leading ladies in his 20-plus year career.  There’s been Vivica Fox, Halle Berry, Monica Calhoun, and in his most recent film, When The Bough Breaks, it was Regina Hall."},
                Article{Image: "image.png", Tittle: "people" , Paragraph: "hello everyone"},
                Article{Image: "image.png", Tittle: "people" , Paragraph: "hello everyone"},
                Article{Image: "image.png", Tittle: "people" , Paragraph: "hello everyone"}}

func GetData(w http.ResponseWriter, r *http.Request) {
                json.NewEncoder(w).Encode(article)
}

func main() {
                router := mux.NewRouter()
                router.HandleFunc("/article", GetData).Methods("GET")
                log.Print("localhost:8000")
                corsObj := handlers.AllowedOrigins([]string{"*"})
                log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsObj)(router)))
}
