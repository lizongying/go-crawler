package request

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
)

type WithContext struct {
	pkg.Context `json:"context,omitempty"`
	pkg.Request `json:"request,omitempty"`
}

func (w *WithContext) MarshalWithContext() ([]byte, error) {
	return json.Marshal(JsonWithContext{
		ContextJson: w.Context.ToContextJson(),
		RequestJson: w.Request.ToRequestJson(),
	})
}
