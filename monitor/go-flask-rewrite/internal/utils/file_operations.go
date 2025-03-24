package utils

import (
    "encoding/json"
    "os"
)

// ReadJSONFile reads the content of a JSON file and unmarshals it into the provided interface.
func ReadJSONFile(filePath string, out interface{}) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    return decoder.Decode(out)
}

// WriteJSONFile marshals the provided interface into JSON and writes it to the specified file.
func WriteJSONFile(filePath string, data interface{}) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "    ")
    return encoder.Encode(data)
}