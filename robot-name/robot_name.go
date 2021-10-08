package robotname

import (
	"errors"
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numset = "0123456789"

const maxNamesAmount = 26 * 26 * 10 * 10 * 10

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

var generatedNames map[string]bool = make(map[string]bool)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		err := r.Reset()
		if err != nil {
			return "", err
		}
	}
	return r.name, nil
}

func (r *Robot) Reset() error {
	name, err := genName()
	if err != nil {
		return err
	}
	r.name = name
	return nil
}

func genName() (string, error) {
	var name string
	for {
		if len(generatedNames) == maxNamesAmount {
			return "", errors.New("Exceeded max names\n")
		}
		name = randString(2, charset) + randString(3, numset)
		if !generatedNames[name] {
			generatedNames[name] = true
			return name, nil
		}
	}
}

func randString(l int, set string) string {
	r := make([]byte, l)
	for i := range r {
		r[i] = set[seededRand.Intn(len(set))]
	}
	return string(r)
}
