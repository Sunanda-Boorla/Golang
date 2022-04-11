package main

import "fmt"

type Stack struct {
	items []string
}

//push
func (s *Stack) Push(i string) {
	s.items = append(s.items, i)

}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() string {
	if s.IsEmpty() {
		return ""
	} else {

		element := s.items[len(s.items)-1]
		return element
	}
}

//pop

func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	} else {
		s.items = s.items[:(len(s.items) - 1)]
		return true
	}

}

func precedence(c string) int {
	if c == "^" {
		return 3
	} else if c == "/" || c == "*" {
		return 2
	} else if c == "+" || c == "-" {
		return 1
	} else {
		return -1
	}
}

func infixToPostfix(s string) string {
	st := Stack{}
	var result string
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			result += string(s[i])
		} else if s[i] == '(' {
			st.Push(string(s[i]))
		} else if s[i] == ')' {
			for st.Peek() != "(" {
				result += string(st.Peek())
				st.Pop()
			}
			st.Pop()
		} else {
			for !st.IsEmpty() && precedence(string(s[i])) <= precedence(string(st.Peek())) {
				result += string(st.Peek())
				st.Pop()
			}
			st.Push(string(s[i]))
		}

	}
	for !(st.IsEmpty()) {
		result += string(st.Peek())
		st.Pop()
	}
	return result
}

func main() {

	fmt.Print("Enter infix expression: ")
	var exp string
	fmt.Scanln(&exp)
	postfix := infixToPostfix(exp)
	fmt.Println("Postfix expression: ", postfix)

}
