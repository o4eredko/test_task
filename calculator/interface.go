package calculator

type Calculator interface {
	Evaluate() (float64, error)
}
