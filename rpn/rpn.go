package rpn

import (
	"strconv"
	"unicode"
	"LearningGo/stack"
)

//Very basic RPN calculator
type Calculator struct {
	stack  stack.Stack
}

func NewCalc() (calc Calculator){
	calc.stack.Push(0)
	return
}

//Process input string
func(calc *Calculator) Process(s string) {
	token := ""
	for _,c := range s {
		switch {
		case c == ' ':
			calc.push(token)
			token = ""
		case unicode.IsDigit(c):
			token += string(c)
		case c == '+':
			calc.op(add)
		case c == '-':
			calc.op(sub)
		case c == '*':
			calc.op(mult)
		case c == '/':
			calc.op(div)
		}
	}
	calc.push(token)
}

//Return current top of the stack
func(calc *Calculator) Result() int {
	return calc.stack.Top()
}

func(calc *Calculator) push(token string) {
	if len(token) > 0 {
		v,_ := strconv.Atoi(token)
		calc.stack.Push(v)
	}
}

func(calc *Calculator) op(f func(int, int) int) {
	v2 := calc.stack.Pop()
	v1 := calc.stack.Pop()
	calc.stack.Push(f(v1, v2))
}

func add(v1,v2 int) int {
	return v1 + v2
}

func sub(v1,v2 int) int {
	return v1 - v2
}

func mult(v1,v2 int) int {
	return v1 * v2
}

func div(v1,v2 int) int {
	return v1 / v2
}
