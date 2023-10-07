package data

import (
    "time"
)

type Artifact struct {
    ID        int64     `json:"id"`
    CreatedAt time.Time `json:"-"`
    Name      string    `json:"name"`
    Origin    string    `json:"origin"`
    Year      Year      `json:"year,omitempty"`
    Type      string    `json:"type"`
    Version   int32     `json:"version"`
}
