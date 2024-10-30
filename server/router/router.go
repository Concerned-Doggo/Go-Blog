package router

import (
	"net/http"

	"github.com/Concerned-Doggo/Go-React-Blog/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
    router := mux.NewRouter()

    router.HandleFunc("/", nil)
    router.HandleFunc("/blogs", controller.GetRecentBlogs).Methods(http.MethodGet, http.MethodOptions)
    router.HandleFunc("/blog/{id}", controller.GetBlogById).Methods(http.MethodGet, http.MethodOptions)
    router.HandleFunc("/createBlog", controller.CreateBlog).Methods(http.MethodPost, http.MethodOptions)
    router.HandleFunc("/updateBlog/{id}", controller.UpdateBlogById).Methods(http.MethodPut, http.MethodOptions)
    router.HandleFunc("/deleteBlog/{id}", controller.DeleteBlogById).Methods(http.MethodDelete, http.MethodOptions)

    router.Use(mux.CORSMethodMiddleware(router))   

    router.Use(corsMiddleware)

    return router
}


func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for demonstration
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        next.ServeHTTP(w, r)
    })
}
