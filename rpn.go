package rpn

import (
	"strconv"
	"strings"
)

func main() {
}

func pop2(stack []float64) (left float64, right float64, newStack []float64) {
	left = stack[len(stack)-2]
	right = stack[len(stack)-1]
	newStack = stack[:len(stack)-1]
	return
}

type operation func(l float64, r float64) float64

func plus(l float64, r float64) float64     { return l + r }
func minus(l float64, r float64) float64    { return l - r }
func divide(l float64, r float64) float64   { return l / r }
func multiply(l float64, r float64) float64 { return l * r }

var operations = map[string]operation{
	"+": plus,
	"-": minus,
	"/": divide,
	"*": multiply,
}

func Rpn(input string) float64 {
	tokens := strings.Split(input, " ")
	var stack []float64
	for _, token := range tokens {
		current_operation, is_operation := operations[token]
		if is_operation {
			left, right, stack := pop2(stack)
			stack = append(stack, current_operation(left, right))
		} else {
			value, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, value)
		}
	}

	return stack[len(stack)-1]
}
