package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Concerned-Doggo/Go-React-Blog/model"
	"github.com/gorilla/mux"
)

func GetRecentBlogs(w http.ResponseWriter, r *http.Request) {
    blogs := getRecentBlogs(10)

    w.Header().Set("content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Methods", "GET")
    json.NewEncoder(w).Encode(blogs)
}


func GetBlogById(w http.ResponseWriter, r *http.Request){

    blogId := mux.Vars(r)["id"]
    blog := getBlogById(blogId)

    w.Header().Set("Content-Type", "application/json")
    // w.Header().Set("Access-Control-Allow-Methods", "GET")
    json.NewEncoder(w).Encode(blog)
}


func CreateBlog(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Methods", "POST")

    var blog model.Blog
    err := json.NewDecoder(r.Body).Decode(&blog)
    checkNilErr(err, "unable to decode request body")

    id := createBlog(blog)

    json.NewEncoder(w).Encode(id)
}


func DeleteBlogById(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Methods", "DELETE")

    blogId := mux.Vars(r)["id"]
    deleteCount := deleteBlogById(blogId)

    json.NewEncoder(w).Encode(deleteCount)
}


func UpdateBlogById(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Methods", "PUT")

    var blog model.Blog
    err := json.NewDecoder(r.Body).Decode(&blog)
    checkNilErr(err, "unable to decode request body for update")

    blogId := mux.Vars(r)["id"]
    modifierCount := updateBlog(blogId, blog)

    json.NewEncoder(w).Encode(modifierCount)
}
