package utils

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func UUIDV1WithoutHyphens() string {
	u, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	uuidStr := u.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	return uuidStr
}
func UUIDV4WithoutHyphens() string {
	u := uuid.New()
	uuidStr := u.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	return uuidStr
}

func StrToUUID(str string)  {
	u, err := uuid.Parse(str)
	fmt.Println(u, err)
}
