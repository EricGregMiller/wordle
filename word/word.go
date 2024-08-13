package word

type Word struct {
	value     string
	frequency int
}

func (word Word) relativeProbability() float64 {
	return 1.0
}

func (word Word) matchGuess(guess Guess) bool {

	return true
}
