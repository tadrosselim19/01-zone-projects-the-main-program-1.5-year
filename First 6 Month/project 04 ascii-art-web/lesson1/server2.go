package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// middleware that logs requests
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
    })
}

// actual handler
func greetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintf(w, "Greetings!\n")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/greet", greetHandler)

    wrappedMux := loggingMiddleware(mux)

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      wrappedMux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    log.Println("Server starting on :8080")
    err := srv.ListenAndServe()
    if err != nil {
        log.Fatal(err)
    }
}
