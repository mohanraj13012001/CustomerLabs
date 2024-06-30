package main

import (
    "log"
    "net/http"
    "CustomerLabs/handle"
    "CustomerLabs/worker"
)

func main() {
    go worker.StartWorker()

    http.HandleFunc("/event", handler.HandleRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
