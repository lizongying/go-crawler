package utils

import (
	"fmt"
	"strings"
)

type Uint64 struct {
	u64 uint64
}

func NewUint64(u64 uint64) *Uint64 {
	return &Uint64{u64: u64}
}

func (u *Uint64) Uint64() uint64 {
	if u == nil {
		return 0
	}
	return u.u64
}
func (u *Uint64) MarshalJSON() ([]byte, error) {
	if u == nil {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%d"`, u.u64)), nil
}
func (u *Uint64) UnmarshalJSON(bytes []byte) error {
	if string(bytes) == `""` {
		return nil
	}

	u64, err := Str2Uint64(strings.Trim(string(bytes), `"`))
	if err != nil {
		return err
	}
	u.u64 = u64
	return nil
}
