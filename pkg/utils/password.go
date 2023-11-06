package utils

import (
	"errors"
	"github.com/lizongying/go-gua64/gua64"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (hashedPassword string, err error) {
	if len(password) < 1 {
		err = errors.New("too short")
		return
	}
	var hashed []byte
	hashed, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	g := gua64.NewGua64()
	hashedPassword = g.Encode(hashed)
	return
}

func ComparePassword(password string, hashedPassword string) bool {
	g := gua64.NewGua64()
	return bcrypt.CompareHashAndPassword(g.Decode(hashedPassword), []byte(password)) == nil
}
