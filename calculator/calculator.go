package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	output, err := toPolish(expression)
	if err != nil {
		return 0, err
	}
	return evaluatePolish(output)
}

func toPolish(expression string) ([]string, error) {
	var output []string
	var stack []rune
	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	i := 0
	for i < len(expression) {
		token := rune(expression[i])

		if unicode.IsDigit(token) || token == '.' {
			number, newIndex := parseNumber(expression, i)
			output = append(output, number)
			i = newIndex + 1
			continue
		}

		if token == '+' || token == '-' || token == '*' || token == '/' {
			for len(stack) > 0 && stack[len(stack)-1] != '(' &&
				precedence[stack[len(stack)-1]] >= precedence[token] {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		} else if token == '(' {
			stack = append(stack, token)
		} else if token == ')' {
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errors.New("troubles with parentheses")
			}
			stack = stack[:len(stack)-1]
		} else if !unicode.IsSpace(token) {
			return nil, fmt.Errorf("invalid character")
		}
		i++
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top == '(' || top == ')' {
			return nil, errors.New("troubles with parentheses")
		}
		output = append(output, string(top))
		stack = stack[:len(stack)-1]
	}

	return output, nil
}


func parseNumber(expression string, start int) (string, int) {
	end := start
	for end < len(expression) && (unicode.IsDigit(rune(expression[end])) || rune(expression[end]) == '.') {
		end++
	}
	return expression[start:end], end - 1
}

func evaluatePolish(polish []string) (float64, error) {
	var stack []float64

	for _, token := range polish {
		switch token {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 { return 0, errors.New("you cannot divide by zero") }
				result = a / b
			}
			stack = append(stack, result)
		default:
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("%s is not a number", token)
			}
			stack = append(stack, value)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}
	return stack[0], nil
}
