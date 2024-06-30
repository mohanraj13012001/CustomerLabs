package transform

import "CustomerLabs/model"

func createNestedMap(value, valueType string) map[string]interface{} {
    return map[string]interface{}{
        "value": value,
        "type":  valueType,
    }
}

func TransformRequest(incoming model.IncomingRequest) model.TransformedRequest {
    attributes := make(map[string]interface{})
    traits := make(map[string]interface{})

    if incoming.Atrk1 != "" {
        attributes[incoming.Atrk1] = createNestedMap(incoming.Atrv1, incoming.Atrt1)
    }
    if incoming.Atrk2 != "" {
        attributes[incoming.Atrk2] = createNestedMap(incoming.Atrv2, incoming.Atrt2)
    }
    if incoming.Uatrk1 != "" {
        traits[incoming.Uatrk1] = createNestedMap(incoming.Uatrv1, incoming.Uatrt1)
    }
    if incoming.Uatrk2 != "" {
        traits[incoming.Uatrk2] = createNestedMap(incoming.Uatrv2, incoming.Uatrt2)
    }
    if incoming.Uatrk3 != "" {
        traits[incoming.Uatrk3] = createNestedMap(incoming.Uatrv3, incoming.Uatrt3)
    }

    return model.TransformedRequest{
        Event:           incoming.Ev,
        EventType:       incoming.Et,
        AppId:           incoming.ID,
        UserId:          incoming.UID,
        MessageId:       incoming.MID,
        PageTitle:       incoming.T,
        PageUrl:         incoming.P,
        BrowserLanguage: incoming.L,
        ScreenSize:      incoming.Sc,
        Attributes:      attributes,
        Traits:          traits,
    }
}
