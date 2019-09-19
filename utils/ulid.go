package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// GenerateULID gera um id único
func GenerateULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}
