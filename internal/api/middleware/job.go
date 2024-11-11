package middleware

import (
    "context"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/manjushettar/ml_server/internal/queue"
)

type Job struct{
    ID     string
    Status string
    Config map[string]interface{}
}

func JobCtx(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){ 
        jobID := chi.URLParam(r, "jobID")

        job, err := queue.GetJob(jobID)
        if err != nil{
            http.Error(w, "Job not found", http.StatusNotFound)
            return
        }

        ctx := context.WithValue(r.Context(), "job", job)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
