// Given a string S of parentheses ‘(‘ or ‘)’, return the minimum number of parentheses needed to balance it. Eg “()()()” -> 0, “()))” -> 2
// Example: 
// 		Input: ")))(((", 	Output: 6
// 		Input: "((()", 		Output: 2 
// 		Input: "())))(", 	Output: 4
package main

import "fmt"

type Stack []string

func (s *Stack) Length() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string{
	if s.IsEmpty() {
		return ""
	}
	index := s.Length()-1
	elem := (*s)[index]
	*s = (*s)[:index]
	return elem
}

func (s *Stack) ElemAtTop() string {
	if s.IsEmpty() {
		return ""
	}
	index := s.Length()-1
	elem := (*s)[index]
	return elem
	
}

func balanceStack(str string) int {
	var stack Stack

	for i := 0; i < len(str); i++ {
		// fmt.Println("Length of stack: ", stack)
		// fmt.Println("Element at the top: ", stack.ElemAtTop())
		if (stack.ElemAtTop() == "(" && string(str[i]) == ")" ) {
			fmt.Println("Elem popped: ", stack.Pop())
		} else {
			stack.Push(string(str[i]))
		}
	}

	return stack.Length()
}


func main() {
	fmt.Println(balanceStack(")))((("))
	fmt.Println(balanceStack("((()"))
	fmt.Println(balanceStack("())))("))
}