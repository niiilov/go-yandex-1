package calculation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calc(expression string) (string, error) {

	expression = strings.ReplaceAll(expression, " ", "")

	if expression == "" || !isCheckParentheses(expression) {
		return "", errors.New("invalid expression")
	}

	result, err := evaluateExpression(expression)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%g", result), nil

}

// Checking The Correct Placement Of Brackets
func isCheckParentheses(expression string) bool {
	dblParentheses := []rune{}

	for _, val := range expression {
		if val == '(' {
			dblParentheses = append(dblParentheses, val)
		} else if val == ')' {
			if len(dblParentheses) == 0 || dblParentheses[len(dblParentheses)-1] != '(' {
				return false
			}
			dblParentheses = dblParentheses[:len(dblParentheses)-1]
		}
	}

	return len(dblParentheses) == 0
}

// Splitting An Expression Into Tokens (Operands)
// Token Processing
// Calculating And Verifying Result
func evaluateExpression(expression string) (float64, error) {

	tokens := []string{}
	currentToken := ""

	for _, val := range expression {
		if val == '+' || val == '-' || val == '*' || val == '/' || val == '(' || val == ')' {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(val))
		} else {
			currentToken += string(val)
		}
	}
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	stack := []float64{}
	operatorStack := []string{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token == "(" {
			operatorStack = append(operatorStack, token)
		} else if token == ")" {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				operator := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				if len(stack) < 2 {
					return 0, errors.New("insufficient operands")
				}
				operand2 := stack[len(stack)-1]
				operand1 := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, calculate(operand1, operand2, operator))
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else if isDigit(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else {
			for len(operatorStack) > 0 && getPrecedence(token) <= getPrecedence(operatorStack[len(operatorStack)-1]) {
				operator := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]

				if len(stack) < 2 {
					return 0, errors.New("insufficient operands")
				}

				operand2 := stack[len(stack)-1]
				operand1 := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, calculate(operand1, operand2, operator))
			}
			operatorStack = append(operatorStack, token)
		}
	}

	for len(operatorStack) > 0 {
		operator := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]

		if len(stack) < 2 {
			return 0, errors.New("insufficient operands")
		}

		operand2 := stack[len(stack)-1]
		operand1 := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		stack = append(stack, calculate(operand1, operand2, operator))
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

// Counting
func calculate(operand1, operand2 float64, operator string) float64 {

	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		if operand2 == 0 {
			return 0
		}
		return operand1 / operand2
	}

	return 0
}

// Determining Whether It Is A Number Or Not
func isDigit(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return false
	}
	return true
}

// Setting Precedence
func getPrecedence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return -1
}
