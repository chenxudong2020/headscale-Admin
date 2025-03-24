package models

type NetworkData struct {
    Sent map[string]string `json:"sent"`
    Recv map[string]string `json:"recv"`
}