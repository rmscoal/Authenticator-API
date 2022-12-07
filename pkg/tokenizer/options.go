package tokenizer

type Option func(*Tokenizer)

func MinCost(cost int) Option {
	return func(t *Tokenizer) {
		t.minCost = cost
	}
}
