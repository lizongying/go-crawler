package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type Request struct{}

func (r *Request) BindJson(w http.ResponseWriter, req *http.Request, data any) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(body, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
