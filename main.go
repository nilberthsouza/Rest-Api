package main

import(
    "github.com/gorilla/mux"
    "net/http"
    "encoding/json"
)

type Post struct{
    ID string `json:"id"`
    Title string `json:"title"`
    Body string `json:"body"`
}

var posts []Post

// a sintaxe para criar um endpoint é assim:
//router.HandleFunc("/<sua-url>",<nome-da-função>).Methods("<method>")

func main(){
    router := mux.NewRouter()
    
    router.HandleFunc("/posts", getPosts).Methods("GET")
    router.HanfleFunc("/posts", createPost).Methods("POST")
    router.HandleFunc("/posts/{id}",getPost).Methods("GET")
    router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
    router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")

    http.ListenAndServe(":8000",router)

}


func getPost(w http.Response.Writer, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    json.NewEncoder(w).Encode(posts)

}
