package pkg

import "fmt"

type Browser string

const (
	BrowserUnknown Browser = ""
	BrowserChrome  Browser = "chrome"
	BrowserEdge    Browser = "edge"
	BrowserSafari  Browser = "safari"
	BrowserFireFox Browser = "firefox"
)

func (b *Browser) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, *b)), nil
}
func (b *Browser) UnmarshalJSON(bytes []byte) error {
	switch string(bytes) {
	case "":
		*b = BrowserUnknown
	case "chrome":
		*b = BrowserChrome
	case "edge":
		*b = BrowserEdge
	case "safari":
		*b = BrowserSafari
	case "firefox":
		*b = BrowserFireFox
	default:
		return fmt.Errorf("invalid")
	}
	return nil
}
