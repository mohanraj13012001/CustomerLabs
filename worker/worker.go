package worker

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
    "CustomerLabs/model"
    "CustomerLabs/transform"
)

var RequestChannel = make(chan model.IncomingRequest)

func StartWorker() {
    for incomingRequest := range RequestChannel {
        transformedRequest := transform.TransformRequest(incomingRequest)
        sendToWebhook(transformedRequest)
    }
}

func sendToWebhook(request model.TransformedRequest) {
    jsonData, err := json.Marshal(request)
    if err != nil {
        log.Println("Error marshaling JSON:", err)
        return
    }

    _, err = http.Post("https://webhook.site/", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        log.Println("Error sending request to webhook:", err)
    }
}
