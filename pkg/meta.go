package pkg

type Meta struct {
	Cookies  map[string]string `json:"cookies,omitempty"`
	Referrer string            `json:"referrer,omitempty"`
}
