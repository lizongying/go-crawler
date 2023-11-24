package api

import (
	"compress/gzip"
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

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}

func (r *Response) OutJsonGzip(w http.ResponseWriter, code int, msg string, data any) {
	jsonData, err := json.Marshal(Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Encoding", "gzip")

	gw := gzip.NewWriter(w)
	defer func() {
		_ = gw.Close()
	}()

	_, _ = gw.Write(jsonData)
}
