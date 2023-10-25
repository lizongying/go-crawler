package utils

import (
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Format("2006-01-02") + `"`), nil
}

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte(`0`), nil
	}
	return []byte(fmt.Sprintf("%d", t.Unix())), nil
}
func (t *Timestamp) UnmarshalJSON(bytes []byte) error {
	i64, err := Str2Int64(string(bytes))
	if err != nil {
		return err
	}
	t.Time = time.Unix(i64, 0)
	return nil
}

type TimestampNano struct {
	time.Time
}

func (t *TimestampNano) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte(`0`), nil
	}
	return []byte(fmt.Sprintf("%d", t.UnixNano())), nil
}
func (t *TimestampNano) UnmarshalJSON(bytes []byte) error {
	i64, err := Str2Int64(string(bytes))
	if err != nil {
		return err
	}
	t.Time = time.Unix(0, i64)
	return nil
}

type DurationSecond struct {
	time.Duration
}

func (d *DurationSecond) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", int(d.Duration/time.Second))), nil
}
func (d *DurationSecond) UnmarshalJSON(bytes []byte) error {
	u64, err := Str2Uint64(string(bytes))
	if err != nil {
		return err
	}
	d.Duration = time.Duration(u64) * time.Second
	return nil
}

type DurationNano struct {
	time.Duration
}

func (d *DurationNano) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", int(d.Duration))), nil
}
func (d *DurationNano) UnmarshalJSON(bytes []byte) error {
	i64, err := Str2Int64(string(bytes))
	if err != nil {
		return err
	}
	d.Duration = time.Duration(i64)
	return nil
}
