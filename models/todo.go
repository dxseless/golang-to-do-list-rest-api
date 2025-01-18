package models

type Todo struct {
    ID     int    `json:"id"`
    Task   string `json:"task"`
    Status string `json:"status"` 
}