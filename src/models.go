package main

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
type sumResponse struct {
	Sum int `json:"sum"`
}
