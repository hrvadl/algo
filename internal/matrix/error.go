package matrix

type DivideByZeroError struct{}

func (e DivideByZeroError) Error() string {
	return "divide by zero"
}
