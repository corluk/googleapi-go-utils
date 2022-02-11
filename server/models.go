package server

import "time"

type ResponseGetUrl struct {
	Time time.Time `json:"time"`
	URL  string    `json:"code"`
}
