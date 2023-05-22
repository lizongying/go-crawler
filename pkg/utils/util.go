package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ExistsDir if dir exists
func ExistsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	if s.IsDir() {
		return true
	}
	s, err = os.Stat(filepath.Dir(path))
	if err != nil {
		return false
	}

	return s.IsDir()
}

// ExistsFile if file exists
func ExistsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// NowStr now str
func NowStr() string {
	return time.Now().Format(time.DateTime)
}

// JsonStr output json string
func JsonStr(i any) string {
	m, _ := json.Marshal(i)
	return string(m)
}

// Struct2JsonKV struct to json key & value
func Struct2JsonKV(i any) (key string, value string) {
	key = reflect.TypeOf(i).Name()
	bs, _ := json.Marshal(i)
	value = string(bs)
	return
}

// InSlice return if in slice
func InSlice[T int | string](t T, sl []T) bool {
	for _, s := range sl {
		if t == s {
			return true
		}
	}
	return false
}

// DumpRead e.g. utils.DumpRead(r)
func DumpRead(reader *io.ReadCloser) {
	byteRes, _ := io.ReadAll(*reader)
	fmt.Println(string(byteRes))
	*reader = io.NopCloser(bytes.NewReader(byteRes))
}

// DumpBytes e.g. utils.DumpBytes(b)
func DumpBytes(b []byte) {
	fmt.Println(string(b))
}

// DumpStr e.g. utils.DumpStr(s)
func DumpStr(s string) {
	fmt.Println(s)
}

// DumpJson e.g. utils.DumpJson(s)
func DumpJson(i interface{}) {
	m, _ := json.Marshal(i)
	fmt.Println(string(m))
}

// DumpFRead e.g. utils.DumpFRead(r, "/tmp/detail.html")
func DumpFRead(reader *io.ReadCloser, path string) {
	byteRes, _ := io.ReadAll(*reader)
	file, _ := os.OpenFile(fmt.Sprintf("%s", path), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	_, _ = io.WriteString(file, string(byteRes))
	*reader = io.NopCloser(bytes.NewReader(byteRes))
}

// DumpFBytes e.g. utils.DumpFBytes(b, "/tmp/detail.html")
func DumpFBytes(b []byte, path string) {
	file, _ := os.OpenFile(fmt.Sprintf("%s", path), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	_, _ = io.WriteString(file, string(b))
}

// DumpFStr e.g. utils.DumpFStr(s, "/tmp/detail.html")
func DumpFStr(s string, path string) {
	file, _ := os.OpenFile(fmt.Sprintf("%s", path), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	_, _ = io.WriteString(file, s)
}

// DumpFJson e.g. utils.DumpFJson(s, "/tmp/data.json")
func DumpFJson(i interface{}, path string) {
	m, _ := json.Marshal(i)
	file, _ := os.OpenFile(fmt.Sprintf("%s", path), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	_, _ = io.WriteString(file, string(m))
}

// StrMd5 md5 string
func StrMd5(sl ...string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(sl, ""))))
}

func Request2Curl(r *pkg.Request) string {
	var args []string
	args = append(args, "curl")
	if r.Method != "GET" {
		args = append(args, "-X", r.Method)
	}
	args = append(args, fmt.Sprintf(`'%s'`, r.Url))
	for k := range r.Header {
		args = append(args, fmt.Sprintf(`-H '%s: %s'`, k, r.Header.Get(k)))
	}
	if r.BodyStr != "" {
		args = append(args, fmt.Sprintf(`--data-raw '%s'`, r.BodyStr))
	}

	return fmt.Sprint(strings.Join(args, " "))
}

// Str2Int Convert string to int
func Str2Int(str string) (i int, err error) {
	i, err = strconv.Atoi(str)
	return
}

// Str2Uint Convert string to uint
func Str2Uint(str string) (i uint, err error) {
	i64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return
	}
	i = uint(i64)
	return
}

// Str2Int8  Convert string to int8
func Str2Int8(str string) (i int8, err error) {
	i64, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return
	}
	i = int8(i64)
	return
}

// Str2Uint8 Convert string to uint8
func Str2Uint8(str string) (i uint8, err error) {
	u64, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return
	}
	i = uint8(u64)
	return
}

// Str2Int16  Convert string to int16
func Str2Int16(str string) (i int16, err error) {
	i64, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return
	}
	i = int16(i64)
	return
}

// Str2Uint16 Convert string to uint16
func Str2Uint16(str string) (i uint16, err error) {
	u64, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return
	}
	i = uint16(u64)
	return
}

// Str2Int32  Convert string to int32
func Str2Int32(str string) (i int32, err error) {
	i64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return
	}
	i = int32(i64)
	return
}

// Str2Uint32 Convert string to uint32
func Str2Uint32(str string) (i uint32, err error) {
	u64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return
	}
	i = uint32(u64)
	return
}

// Str2Int64  Convert string to int64
func Str2Int64(str string) (i int64, err error) {
	i, err = strconv.ParseInt(str, 10, 64)
	return
}

// Str2Uint64 Convert string to uint64
func Str2Uint64(str string) (i uint64, err error) {
	i, err = strconv.ParseUint(str, 10, 64)
	return
}

func Int2Str[T int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64](i T) (str string) {
	str = fmt.Sprintf("%d", i)
	return
}
