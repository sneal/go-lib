package go_lib

import "time"

type App struct {
	GUID      string            `json:"guid"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Name      string            `json:"name"`
	State     string            `json:"state"`
	Links     map[string]string `json:"links"`
}
