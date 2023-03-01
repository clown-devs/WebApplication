package model

import (
	"math/rand"
	"strings"
	"time"
)

type Token struct {
	EmployeeId uint64 `json:"employee_id"`
	Token      string `json:"auth_token"`
}

func (t *Token) String() string {
	return t.Token
}

func NewToken(id uint64, salt string) *Token {
	rand.Seed(time.Now().UnixNano() + int64(len(salt)))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 6 + len(strings.Split(salt, ""))

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
		b.WriteRune(rune(salt[rand.Intn(len(salt))]))
	}
	str := b.String()
	return &Token{EmployeeId: id, Token: str}

}
