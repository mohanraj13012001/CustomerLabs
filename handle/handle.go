package handler

import (
    "encoding/json"
    "net/http"
   model "CustomerLabs/model"
    "fmt"
    "io/ioutil"
)

// Helper function to create a nested map
func createNestedMap(key, value, valueType string) map[string]interface{} {
    return map[string]interface{}{
        "value": value,
        "type":  valueType,
    }
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
    // Get the request URL
    url := r.URL.String()
    fmt.Printf("Request URL: %s\n", url)

    // Read the request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Unmarshal the JSON data
    var inputEvent model.IncomingRequest
    err = json.Unmarshal(body, &inputEvent)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Log the JSON data
    fmt.Printf("Request Body: %+v\n", inputEvent)

    // Transform the data
    outputEvent :=model. OutputEvent{
        Event:           inputEvent.Ev,
        EventType:       inputEvent.Et,
        AppID:           inputEvent.ID,
        UserID:          inputEvent.UID,
        MessageID:       inputEvent.MID,
        PageTitle:       inputEvent.T,
        PageURL:         inputEvent.P,
        BrowserLanguage: inputEvent.L,
        ScreenSize:      inputEvent.Sc,
        Attributes: map[string]interface{}{
            inputEvent.Atrk1: createNestedMap(inputEvent.Atrk1, inputEvent.Atrv1, inputEvent.Atrt1),
            inputEvent.Atrk2: createNestedMap(inputEvent.Atrk2, inputEvent.Atrv2, inputEvent.Atrt2),
        },
        Traits: map[string]interface{}{
            inputEvent.Uatrk1: createNestedMap(inputEvent.Uatrk1, inputEvent.Uatrv1, inputEvent.Uatrt1),
            inputEvent.Uatrk2: createNestedMap(inputEvent.Uatrk2, inputEvent.Uatrv2, inputEvent.Uatrt2),
            inputEvent.Uatrk3: createNestedMap(inputEvent.Uatrk3, inputEvent.Uatrv3, inputEvent.Uatrt3),
        },
    }

    // Marshal the transformed data back to JSON
    responseBody, err := json.Marshal(outputEvent)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Write the response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(responseBody)
}
