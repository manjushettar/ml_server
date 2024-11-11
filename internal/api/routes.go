package api

import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi/v5"
)

// Ping route to verify server is online.
func Ping(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid request.")
        return
    }

    if _, err := w.Write([]byte("Pinged.")); err != nil{
        http.Error(w, "Couldn't ping server.", http.StatusInternalServerError)
        return
    }
}

// Job routes group
// Jobs/    -> POST method to submit a Job
// Jobs/    -> GET method to list all Jobs
// {jobID}/ -> Sends to a SpecificJob routing group
func Jobs(r chi.Router){
    r.Post("/", submitJob)
    r.Get("/", listJobs)

    r.Route("/{jobID}", SpecificJob)
}

// Specific Job routes group
// {jobID}/        -> GET method to check status of specified Job
// {jobID}/results -> GET method to check results of specified Job
// {jobID}/        -> DELETE method to delete the specified Job
func SpecificJob(r chi.Router){
    r.Use(JobCtx)
    r.Get("/", getJobStatus)
    r.Get("/results", getResults)
    r.Delete("/", cancelJob)
}


func submitJob(w http.ResponseWriter, r *http.Request){
    if r.Method != "POST"{
        fmt.Fprintf(w, "Invalid request")
        return
    }
} 


// GPU routes group
// GPU/       -> GET method to verify GPU is reachable
// GPU/status -> GET method to check status of the GPU
// GPU/memory -> GET method to check memory of the GPU
func GPU(r chi.Router){
    r.Get("/", pingGPU)
    r.Get("/status", getGPUStatus)
    r.Get("/memory", getGPUMemory)
}


func getGPUStatus(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid request.")
        return
    }

    if _, err := w.Write([]byte("Good status.")); err != nil{
        http.Error(w, "Couldn't access GPU status.", http.StatusInternalServerError)
        return
    }
}

func pingGPU(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid request.")
        return
    }

    if _, err := w.Write([]byte("GPU here.")); err != nil{
        http.Error(w, "Couldn't access GPU server.", http.StatusInternalServerError)
        return
    }
}

func getGPUMemory(w http.ResponseWriter, r *http.Request){
    if r.Method != "GET"{
        fmt.Fprintf(w, "Invalid request.")
        return
    }

    if _, err := w.Write([]byte("Lots of memory.")); err != nil{
        http.Error(w, "Couldn't access GPU memory.", http.StatusInternalServerError)
        return
    }
}


