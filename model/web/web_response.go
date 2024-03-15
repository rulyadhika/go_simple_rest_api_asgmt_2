package web

type WebResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   any    `json:"data"`
}
