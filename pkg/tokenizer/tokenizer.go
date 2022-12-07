package tokenizer

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const _defaultMinCost = 4

type Tokenizer struct {
	minCost int
}

func New(opts ...Option) *Tokenizer {
	tk := &Tokenizer{
		minCost: _defaultMinCost,
	}

	// Custom options
	for _, opt := range opts {
		opt(tk)
	}

	return tk
}

func (T *Tokenizer) GenerateFromPassword(password []byte) ([]byte, error) {
	fmt.Println(T.minCost)
	token, err := bcrypt.GenerateFromPassword(password, T.minCost)
	if err != nil {
		return nil, err
	}

	return token, nil
}
