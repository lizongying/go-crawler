package utils

import (
	"fmt"
	"net/url"
	"strings"
)

type Url struct {
	*url.URL
}

func (u *Url) MarshalJSON() ([]byte, error) {
	if u.URL == nil {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, u.String())), nil
}
func (u *Url) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == `""` {
		return nil
	}
	proxy, err := url.Parse(strings.Trim(string(bytes), `"`))
	if err != nil {
		return err
	}
	if proxy.Host == "" {
		return fmt.Errorf("loss host")
	}
	u.URL = proxy
	return nil
}
