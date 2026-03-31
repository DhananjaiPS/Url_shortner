package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Url struct {
	Id          string    `json:"id"`
	OriginalUrl string    `json:"originalurl"`
	ShortUrl    string    `json:"shorturl"`
	CreatedAt   time.Time `json:"createdat"`
}

//in map

// 123oxd2 --> {
// 		       	Id: 1,
// 		       	OriginalUrl: "https://www.google.com",
// 		       	ShortUrl: "123oxd2",
// 		       	CreatedAt: time.Now(),
// 		       	}

//creating in memory db

func generateshortUrl(OriginalUrl string) string {

	hasher := md5.New()               //create empty object
	hasher.Write([]byte(OriginalUrl)) // hasher ke liye url write kri
	data := hasher.Sum(nil)           // bytes bana dega
	hash := hex.EncodeToString(data)  // ab byte ko convert kr dega string mein
	fmt.Println(hash)
	fmt.Printf("%T", hash) //string
	return string(hash[:8])

}

func createUrl(originalurl string) string {

	shortUrl := generateshortUrl(originalurl)

	UrlDb[shortUrl] = Url{
		Id:          shortUrl,
		OriginalUrl: originalurl,
		ShortUrl:    shortUrl,
		CreatedAt:   time.Now(),
	}
	return shortUrl

}

func GetUrl(id string) (Url, error) {

	SearchedId, ok := UrlDb[id]
	if !ok {
		return Url{}, errors.New("Id not found")
	}
	return SearchedId, nil

}

var UrlDb = map[string]Url{}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handler function is called on / route ")

}

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	// var ResponseData data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "invaild response ", http.StatusBadRequest)
		return
	}
	shortUrl := createUrl(data.URL)
	// fmt.Fprint(w, shortUrl)
	response := struct {
		ShortUrl string `json:"shorturl"`
	}{ShortUrl: shortUrl}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RedirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]

	fmt.Println("Redirect ID received:", id)
	fmt.Println("URL DB:", UrlDb)

	url, err := GetUrl(id)
	if err != nil {
		http.Error(w, "Invalid req", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, url.OriginalUrl, http.StatusFound)
}

func main() {
	fmt.Println("Welcome to Url shortner")
	// myUrl := "https://docs.google.com/document/d/1MIFBHjWVrwiGzL3CuYD_1zXdXaP_l6TC5La1WqnnBeQ/edit?tab=t.0"

	http.HandleFunc("/", handler)                 // iske bina server on kro ge toh 404 page not found aye ga pr ise hum kya kr rhe h ki use page / pr ek function ko call or print kr dega print statemnet
	http.HandleFunc("/shortner", ShortUrlHandler) //fetch the url insert in struct and then decode it and then return
	http.HandleFunc("/redirect/", RedirectUrlHandler)
	fmt.Println("Server is getting ready on port:3000")
	fmt.Println("Server is running on port:3000")

	error := http.ListenAndServe(":3000", nil)
	if error != nil {
		panic(error)
	}

}
