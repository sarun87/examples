package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    osName,_ := os.Hostname()
    t := time.Now()
    log.Printf("Request url: %s Sender: %s\n", r.URL.Path[1:], r.RemoteAddr)
    fmt.Fprintf(w, "[%s] Hello world from %s", t.UTC().Format(time.UnixDate), osName)
}

func main() {
    port, ok := os.LookupEnv("PORT")
    if !ok {
        port = "8080"
    }
    fmt.Println("Starting server on port", port)
    port = fmt.Sprintf(":%s", port)
    http.HandleFunc("/", handler)
    http.ListenAndServe(port, nil)
}
