package calculator

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"

	"github.com/golang-collections/collections/stack"
)

const (
	openingBracket = "("
	closingBracket = ")"
)

var (
	operatorRegexp = regexp.MustCompile(`^([+\-*/])($|\s)`)
	numberRegexp   = regexp.MustCompile(`^([+-]?\d+)`)
)

type calculator struct {
	expression string
	tokens     []string
	stack      *stack.Stack
}

func NewCalculator(expression string) Calculator {
	return &calculator{
		expression: expression,
		stack:      stack.New(),
	}
}

func (c *calculator) Evaluate() (float64, error) {
	err := c.tokenizeExpressionRPN(c.expression)
	if err != nil {
		return 0, err
	}
	log.Printf("RPN expression: {%s}\n", c.tokens)
	for _, token := range c.tokens {
		switch {
		case isOperator(token):
			operator, _ := NewOperator(token)
			result, err := c.evaluateOneOperation(operator)
			if err != nil {
				return 0, err
			}
			c.stack.Push(result)
		default:
			number, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			c.stack.Push(number)
		}
	}
	if c.stack.Len() != 1 {
		return 0, fmt.Errorf(
			"%w stack doesn't contain result after evaluation: %v",
			ErrInvalidExpression, c.stack,
		)
	}
	return c.stack.Pop().(float64), nil
}

func (c *calculator) evaluateOneOperation(operator Operator) (float64, error) {
	if c.stack.Len() < 2 {
		return 0, fmt.Errorf(
			"%w: not enough elements in the stack to evaluate: %s",
			ErrInvalidExpression, operator,
		)
	}
	op2, op1 := c.stack.Pop().(float64), c.stack.Pop().(float64)
	return operator.Evaluate(op1, op2), nil
}

func (c *calculator) tokenizeExpressionRPN(expression string) error {
	for i := 0; i < len(expression); i++ {
		char := string(expression[i])
		switch {
		case unicode.IsSpace(rune(expression[i])):
			continue
		case operatorRegexp.MatchString(expression[i:]):
			c.addTokensSinceLastOperator(char)
			c.stack.Push(char)
		case numberRegexp.MatchString(expression[i:]):
			number := numberRegexp.FindString(expression[i:])
			c.tokens = append(c.tokens, number)
			i += len(number) - 1
		case char == openingBracket:
			c.stack.Push(char)
		case char == closingBracket:
			c.addTokensBetweenBrackets()
		default:
			return fmt.Errorf("%w: invalid symbol: %s", ErrInvalidExpression, expression[i:])
		}
	}
	c.addAllTokensFromStack()
	return nil
}

func (c *calculator) addAllTokensFromStack() {
	for c.stack.Len() > 0 {
		token, _ := c.stack.Pop().(string)
		c.tokens = append(c.tokens, token)
	}
}

func (c *calculator) addTokensSinceLastOperator(char string) {
	for c.stack.Len() > 0 {
		top := c.stack.Peek().(string)
		if top == openingBracket || !hasHigherPrecedence(top, char) {
			break
		}
		c.tokens = append(c.tokens, top)
		c.stack.Pop()
	}
}

func (c *calculator) addTokensBetweenBrackets() {
	for c.stack.Len() > 0 {
		top := c.stack.Peek().(string)
		if top == openingBracket {
			break
		}
		c.tokens = append(c.tokens, top)
		c.stack.Pop()
	}
	c.stack.Pop()
}

func isOperator(char string) bool {
	_, err := NewOperator(char)
	return err == nil
}

func hasHigherPrecedence(op1, op2 string) bool {
	operator1, err := NewOperator(op1)
	if err != nil {
		return false
	}
	operator2, err := NewOperator(op2)
	if err != nil {
		return true
	}
	return operator1.GetWeight() >= operator2.GetWeight()
}
