package main

import(
    "github.com/gorilla/mux"
    "net/http"
    "encoding/json"
    "math/rand"
    "strconv"
)

type Post struct{
    ID string `json:"id"`
    Title string `json:"title"`
    Body string `json:"body"`
}

var posts []Post

// a sintaxe para criar um endpoint é assim:
//router.HandleFunc("/<sua-url>",<nome-da-função>).Methods("<method>")



func getPosts(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    json.NewEncoder(w).Encode(posts)

}

func getPost(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    for _, item := range posts{
        if item.ID == params["id"]{
            json.NewEncoder(w).Encode(item)
            break
        }
        return
    }
    json.NewEncoder(w).Encode(&Post{})
}

func createPost(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    var post Post
    _ = json.NewDecoder(r.Body).Decode(post)
    post.ID = strconv.Itoa(rand.Intn(1000000))
    posts = append(posts, post)
    json.NewEncoder(w).Encode(&post)

}

func updatePost(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    for index, item := range posts{
        if item.ID == params["id"]{
            posts = append(posts[:index],posts[index+1:]...)
            
            var post Post
            _ = json.NewDecoder(r.Body).Decode(post)
            post.ID = params["id"]
            posts = append(posts, post)
            json.NewEncoder(w).Encode(&post)

            return
            
        
        }
    }
    json.NewEncoder(w).Encode(posts)
}

func deletePost(w http.ResponseWriter, r * http.Request){
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range posts{
        if item.ID == params["id"]{
            posts = append(posts[:index], posts[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(posts)
}


func main(){
    router := mux.NewRouter()
    //apenas para fim de testes
    posts = append(posts, Post{ID: "1",Title:"My first post",Body:
    "this is the content of my first post"})
    posts = append(posts, Post{ID:"2",Title: "My second post",Body:
    "This is the content of my second post"})



    router.HandleFunc("/posts", getPosts).Methods("GET")
    router.HandleFunc("/posts", createPost).Methods("POST")
    router.HandleFunc("/posts/{id}",getPost).Methods("GET")
    router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
    router.HandleFunc("/post/{id}", deletePost).Methods("DELETE")

    http.ListenAndServe(":8000",router)

}
