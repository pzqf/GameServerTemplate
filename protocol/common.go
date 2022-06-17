package protocol

import "time"

type TestPingReq struct {
	Id   int32     `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

type TestPingRes struct {
	Id   int32     `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}
