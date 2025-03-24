package handlers

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
    "os"
    "sync"
)

var mu sync.Mutex

func GetDataRecord(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    jsonDataNow, err := getInfo()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    recvSpeed := jsonDataNow["recv_speed"].(string)
    sentSpeed := jsonDataNow["sent_speed"].(string)

    file, err := ioutil.ReadFile("data.json")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var jsonDataLocal map[string]interface{}
    if err := json.Unmarshal(file, &jsonDataLocal); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    keys := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y"}
    for i := 0; i < len(keys)-1; i++ {
        jsonDataLocal["sent"].(map[string]interface{})[keys[i]] = jsonDataLocal["sent"].(map[string]interface{})[keys[i+1]]
        jsonDataLocal["recv"].(map[string]interface{})[keys[i]] = jsonDataLocal["recv"].(map[string]interface{})[keys[i+1]]
    }

    jsonDataLocal["sent"].(map[string]interface{})["y"] = sentSpeed
    jsonDataLocal["recv"].(map[string]interface{})["y"] = recvSpeed

    updatedData, err := json.MarshalIndent(jsonDataLocal, "", "    ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := ioutil.WriteFile("data.json", updatedData, 0644); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(updatedData)
}

func getInfo() (map[string]interface{}, error) {
    // This function should implement the logic to retrieve the current network data usage
    // and return it as a map. This is a placeholder for the actual implementation.
    return map[string]interface{}{
        "recv_speed": "100", // Placeholder value
        "sent_speed": "200", // Placeholder value
    }, nil
}