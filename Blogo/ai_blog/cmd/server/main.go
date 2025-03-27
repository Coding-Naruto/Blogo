package main

import (
    "log"
    "net/http"
    "ai_blog/internal/handlers"
    "ai_blog/internal/models"
)

func main() {
    models.InitDB()

    mux := http.NewServeMux()
    mux.HandleFunc("/", handlers.Home)
    mux.HandleFunc("/post", handlers.Post)
    mux.HandleFunc("/edit", handlers.Edit)
    mux.HandleFunc("/delete", handlers.Delete)
    mux.HandleFunc("/view", handlers.ViewPost)

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        log.Fatal(err)
    }
}