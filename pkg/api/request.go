package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct{}

func (r *Request) BindJson(w http.ResponseWriter, req *http.Request, data any) {
	fmt.Println(1111111111111, req.Method)
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
