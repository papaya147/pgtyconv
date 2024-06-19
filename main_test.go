package pgtyconv

import (
	"github.com/papaya147/randomize"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	randomize.RegisterCustomRandomizer(randomTimeUnit)
	os.Exit(m.Run())
}

func randomTimeUnit() TimeUnit {
	r := rand.New(rand.NewSource(time.Now().UnixMilli()))
	args := []TimeUnit{
		TimeUnitNanoseconds,
		TimeUnitMicroseconds,
		TimeUnitMilliseconds,
		TimeUnitSeconds,
	}
	return args[r.Intn(len(args))]
}
