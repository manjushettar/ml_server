package main

import (
    "fmt"
    "net/http"
    "github.com/manjushettar/ml_server/internal/api"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main(){
    fmt.Println("Starting server")
    router := chi.NewRouter()
    

    router.Use(middleware.RequestID)
    router.Use(middleware.RealIP)
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)
    
    router.Get("/", api.Ping) 
    
    router.Route("/jobs", api.Jobs)

    router.Route("/gpu", api.GPU)

    if err := http.ListenAndServe(":3000", router); err != nil{
        fmt.Println("Something went wrong with the server!")
    }
}
