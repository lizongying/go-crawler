package pkg

import "fmt"

type Platform string

const (
	PlatformUnknown Platform = ""
	PlatformWindows Platform = "windows"
	PlatformMac     Platform = "mac"
	PlatformAndroid Platform = "android"
	PlatformIphone  Platform = "iphone"
	PlatformIpad    Platform = "ipad"
	PlatformLinux   Platform = "linux"
)

func (p *Platform) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, *p)), nil
}
func (p *Platform) UnmarshalJSON(bytes []byte) error {
	switch string(bytes) {
	case "":
		*p = PlatformUnknown
	case "windows":
		*p = PlatformWindows
	case "mac":
		*p = PlatformMac
	case "android":
		*p = PlatformAndroid
	case "iphone":
		*p = PlatformIphone
	case "ipad":
		*p = PlatformIpad
	case "linux":
		*p = PlatformLinux
	default:
		return fmt.Errorf("invalid")
	}
	return nil
}
