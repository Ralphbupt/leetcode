package main

import (
	"fmt"
)

func main(){
	l := longestValidParentheses(")()())()()(")
	fmt.Println(l)
}

type stack struct {
	list []int
}
func NewStack()stack{
	s := stack{}
	s.init()
	return  s
}
func (s *stack) init(){
	s.list = make([]int,0)
}

func (s *stack) Push(val int){
	s.list = append(s.list, val)
}

func (s *stack) Pop()int{
	length := len(s.list)
	top := s.list[length-1]
	s.list = s.list[:length-1]
	return top
}

func (s *stack) Top()int{
	return  s.list[len(s.list)-1]
}

func (s *stack) Empty() bool{
	return  len(s.list) == 0
}

func longestValidParentheses(s string) int {

	st := NewStack()
	maxLen := 0
	for i,v := range  s{
		if v == 40{
			st.Push(i)
		}else {
			if !st.Empty()&&s[st.Top()] == 40 {
				st.Pop()
			}else {
				st.Push(i)
			}
		}
	}
	n := len(s)
	if st.Empty(){
		return  n
	}else{
		 a := n
		 b := 0;
		for !st.Empty(){
			b = st.Top();
			st.Pop();
			if a-b- 1 > maxLen{
				maxLen = a-b-1
			}
			a = b;
		}
		if a > maxLen{
			maxLen = a
		}
	}
	return  maxLen
}


