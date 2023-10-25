package pkg

import (
	"fmt"
)

type Client string

const (
	ClientUnknown Client = ""
	ClientGo      Client = "go"
	ClientBrowser Client = "browser"
)

func (c *Client) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, *c)), nil
}
func (c *Client) UnmarshalJSON(bytes []byte) error {
	switch string(bytes) {
	case "":
		*c = ClientUnknown
	case "go":
		*c = ClientGo
	case "browser":
		*c = ClientBrowser
	default:
		return fmt.Errorf("invalid")
	}
	return nil
}
