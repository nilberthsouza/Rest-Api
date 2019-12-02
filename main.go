package main

import(
    "github.com/gorilla/mux"
    "net/http"
    "encoding/json"
)


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


