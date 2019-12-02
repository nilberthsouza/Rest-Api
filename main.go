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

func getPosts(w http.Response.Writer, r *http.Request){
    w.Heade().Set("Content-Type","application/json")
    params := mux.Vars(r)
    for _, item := range posts{
        if item.ID == params["id"]{
            json.NewEnconder(w).Encode(item)
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
    json.NewEncoder(w).Enconde(&post)

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
            json.NewEnconder(w).Encode(&post)

            return
            
        
        }
    }
    json.NewEncoder(w).Encode(post)
}
