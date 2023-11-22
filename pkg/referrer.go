package pkg

import "strings"

type ReferrerPolicy uint8

const (
	DefaultReferrerPolicy ReferrerPolicy = iota
	NoReferrerPolicy
)

func (r ReferrerPolicy) String() string {
	switch r {
	case NoReferrerPolicy:
		return "NoReferrerPolicy"
	default:
		return "DefaultReferrerPolicy"
	}
}

func ReferrerPolicyFromString(referrerPolicy string) ReferrerPolicy {
	switch strings.ToLower(referrerPolicy) {
	case "1":
		return NoReferrerPolicy
	case "NoReferrerPolicy":
		return NoReferrerPolicy
	default:
		return DefaultReferrerPolicy
	}
}
