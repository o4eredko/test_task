package calculator

import "errors"

var operatorWeights = map[string]uint{
	"+": 1, "-": 1, "*": 2, "/": 2,
}

type (
	Operator interface {
		Evaluate(op1, op2 float64) float64
		GetWeight() uint
	}

	addition struct {
		symbol string
	}
	subtraction struct {
		symbol string
	}
	multiplication struct {
		symbol string
	}
	division struct {
		symbol string
	}
)

func NewOperator(token string) (Operator, error) {
	switch token {
	case "+":
		return &addition{symbol: token}, nil
	case "-":
		return &subtraction{symbol: token}, nil
	case "*":
		return &multiplication{symbol: token}, nil
	case "/":
		return &division{symbol: token}, nil
	default:
		return nil, errors.New("invalid operator")
	}
}

func (a *addition) Evaluate(op1, op2 float64) float64 {
	return op1 + op2
}

func (a *addition) GetWeight() uint {
	return operatorWeights[a.symbol]
}

func (s *subtraction) Evaluate(op1, op2 float64) float64 {
	return op1 - op2
}

func (s *subtraction) GetWeight() uint {
	return operatorWeights[s.symbol]
}

func (m *multiplication) Evaluate(op1, op2 float64) float64 {
	return op1 * op2
}

func (m *multiplication) GetWeight() uint {
	return operatorWeights[m.symbol]
}

func (d *division) Evaluate(op1, op2 float64) float64 {
	return op1 / op2
}

func (d *division) GetWeight() uint {
	return operatorWeights[d.symbol]
}
