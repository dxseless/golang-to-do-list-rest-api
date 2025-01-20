package models

import "time"

type Todo struct {
    ID        int       `json:"id"`
    Task      string    `json:"task"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}