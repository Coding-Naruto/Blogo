package ai

import (
    "encoding/json"
    "log"
    "net/http"
    "bytes"
    "io/ioutil"
)

const hfAPIURL = "https://api-inference.huggingface.co/models/facebook/bart-large-cnn"

type HFResponse struct {
    Summary string `json:"summary_text"`
}

func SummarizeText(text string) (string, error) {
    client := &http.Client{}
    requestBody, err := json.Marshal(map[string]string{
        "inputs": text,
    })
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", hfAPIURL, bytes.NewBuffer(requestBody))
    if err != nil {
        return "", err
    }
    req.Header.Set("Authorization", "Bearer YOUR_HUGGING_FACE_API_KEY")
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var hfResp []HFResponse
    err = json.Unmarshal(body, &hfResp)
    if err != nil {
        return "", err
    }

    if len(hfResp) > 0 {
        return hfResp[0].Summary, nil
    }
    return "", nil
}