package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome to the home page")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request)  {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteSingleArticle(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id := vars["id"]
	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			json.NewEncoder(w).Encode("Success")
		}
	}

}

func updateSingleArticle(w http.ResponseWriter, r *http.Request)  {
	var article Article
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &article)
	for _, oldArticle := range Articles {
		if oldArticle.Id == article.Id {
			oldArticle.Title = article.Title
			oldArticle.Desc = article.Desc
			oldArticle.Content = article.Content
			json.NewEncoder(w).Encode(oldArticle)
		}
	}
}



func handleRequest()  {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteSingleArticle).Methods("DELETE")
	myRouter.HandleFunc("/article", updateSingleArticle).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}

func main()  {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequest()
}
