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
