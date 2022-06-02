package uuid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
)

// GenerateNewUuid 生成唯一键
// max length: 20 + 26 = 46
func GenerateNewUuid(length int) string {
	randomStr := NewUlidSeq() + NewXidSeq()
	letters := []rune(randomStr)
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NewUlidSeq() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

func NewXidSeq() string {
	guid := xid.New()
	return guid.String()
}

var numbers = []rune("0123456789")

// NewCode 生成验证码, 最大长度为10
func NewCode(n int) string {
	if n > 10 {
		n = 10
	}
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}
