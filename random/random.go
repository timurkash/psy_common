package random

import (
	"github.com/timurkash/psy_common/config"
	"math/rand"
	"time"
)

var (
	symbols = config.GetEnv("RANDOM_SYMBOLS", "_$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letterRunes = []rune(symbols)
	CODE_SIZE = config.GetEnvInt("CODE_SIZE", 32)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Generate() string {
	random := randStringRunes(CODE_SIZE)
	return random
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


