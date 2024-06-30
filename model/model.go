package model

type IncomingRequest struct {
    Ev     string `json:"ev"`
    Et     string `json:"et"`
    ID     string `json:"id"`
    UID    string `json:"uid"`
    MID    string `json:"mid"`
    T      string `json:"t"`
    P      string `json:"p"`
    L      string `json:"l"`
    Sc     string `json:"sc"`
    Atrk1  string `json:"atrk1"`
    Atrv1  string `json:"atrv1"`
    Atrt1  string `json:"atrt1"`
    Atrk2  string `json:"atrk2"`
    Atrv2  string `json:"atrv2"`
    Atrt2  string `json:"atrt2"`
    Uatrk1 string `json:"uatrk1"`
    Uatrv1 string `json:"uatrv1"`
    Uatrt1 string `json:"uatrt1"`
    Uatrk2 string `json:"uatrk2"`
    Uatrv2 string `json:"uatrv2"`
    Uatrt2 string `json:"uatrt2"`
    Uatrk3 string `json:"uatrk3"`
    Uatrv3 string `json:"uatrv3"`
    Uatrt3 string `json:"uatrt3"`
}

type TransformedRequest struct {
    Event           string                 `json:"event"`
    EventType       string                 `json:"event_type"`
    AppId           string                 `json:"app_id"`
    UserId          string                 `json:"user_id"`
    MessageId       string                 `json:"message_id"`
    PageTitle       string                 `json:"page_title"`
    PageUrl         string                 `json:"page_url"`
    BrowserLanguage string                 `json:"browser_language"`
    ScreenSize      string                 `json:"screen_size"`
    Attributes      map[string]interface{} `json:"attributes"`
    Traits          map[string]interface{} `json:"traits"`
}
type OutputEvent struct {
    Event          string                 `json:"event"`
    EventType      string                 `json:"event_type"`
    AppID          string                 `json:"app_id"`
    UserID         string                 `json:"user_id"`
    MessageID      string                 `json:"message_id"`
    PageTitle      string                 `json:"page_title"`
    PageURL        string                 `json:"page_url"`
    BrowserLanguage string                 `json:"browser_language"`
    ScreenSize     string                 `json:"screen_size"`
    Attributes     map[string]interface{} `json:"attributes"`
    Traits         map[string]interface{} `json:"traits"`
}