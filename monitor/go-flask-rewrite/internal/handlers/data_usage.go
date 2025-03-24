package handlers

import (
    "encoding/json"
    "math"
    "net/http"
    "os"
    "strconv"

    "go-flask-rewrite/internal/models"
)

func GetDataUsage(w http.ResponseWriter, r *http.Request) {
    file, err := os.ReadFile("data.json")
    if err != nil {
        http.Error(w, "Unable to read data file", http.StatusInternalServerError)
        return
    }

    var data models.Data
    if err := json.Unmarshal(file, &data); err != nil {
        http.Error(w, "Error parsing JSON data", http.StatusInternalServerError)
        return
    }

    recvValues := make([]float64, 0, len(data.Recv))
    for _, value := range data.Recv {
        recvValue, _ := strconv.ParseFloat(value, 64)
        recvValues = append(recvValues, recvValue)
    }

    newRecvValues := make(map[string]string)
    for i := 0; i < len(recvValues)-1; i++ {
        newRecvValues[string('a'+i)] = strconv.Itoa(int(math.Ceil(recvValues[i+1] - recvValues[i])))
    }

    sentValues := make([]float64, 0, len(data.Sent))
    for _, value := range data.Sent {
        sentValue, _ := strconv.ParseFloat(value, 64)
        sentValues = append(sentValues, sentValue)
    }

    newSentValues := make(map[string]string)
    for i := 0; i < len(sentValues)-1; i++ {
        newSentValues[string('a'+i)] = strconv.Itoa(int(math.Ceil(sentValues[i+1] - sentValues[i])))
    }

    newData := map[string]interface{}{
        "recv": newRecvValues,
        "sent": newSentValues,
    }

    response, err := json.Marshal(newData)
    if err != nil {
        http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}