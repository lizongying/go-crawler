package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func (r *Response) OutJson(w http.ResponseWriter, code int, msg string, data any) {
	jsonData, err := json.Marshal(Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}
