package router

import (
	"github.com/Concerned-Doggo/Go-React-Blog/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
    router := mux.NewRouter()

    router.HandleFunc("/api", nil)
    router.HandleFunc("/api/blogs", controller.GetRecentBlogs).Methods("GET")
    router.HandleFunc("/api/blog/{id}", controller.GetBlogById).Methods("GET")
    router.HandleFunc("/api/createBlog", controller.CreateBlog).Methods("POST")
    router.HandleFunc("/api/updateBlog/{id}", controller.UpdateBlogById).Methods("PUT")
    router.HandleFunc("/api/deleteBlog/{id}", controller.DeleteBlogById).Methods("DELETE")

    return router
}
