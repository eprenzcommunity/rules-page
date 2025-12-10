package main

import (
    "fmt"
    "net/http"
    "time"
)

type statusRecorder struct {
    http.ResponseWriter
    status int
}

func (r *statusRecorder) WriteHeader(code int) {
    r.status = code
    r.ResponseWriter.WriteHeader(code)
}

func main() {
    fs := http.FileServer(http.Dir("."))

    logged := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        rec := &statusRecorder{ResponseWriter: w, status: 200}

        fs.ServeHTTP(rec, r)

        duration := time.Since(start)
        t := time.Now().Format(time.RFC3339)

        fmt.Printf("%s  %s  %s  %d  %s\n",
            t,
            r.Method,
            r.URL.Path,
            rec.status,
            duration,
        )
    })

    http.ListenAndServe(":8000", logged)
}

