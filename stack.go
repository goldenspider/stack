package stack

import (
	"container/list"
	"fmt"
)

type Stack struct {
	*list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (stack *Stack) Push(value interface{}) {
	stack.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.Back()
	if e != nil {
		stack.Remove(e)
		return e.Value
	}
	return nil
}

//Look up the top element in the stack, but not pop.
func (stack *Stack) LookTop() interface{} {
	e := stack.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (stack *Stack) Len() int {
	return stack.List.Len()
}

func (stack *Stack) Empty() bool {
	return stack.Len() == 0
}

func (stack *Stack) Clone() *Stack {
	clone := NewStack()
	clone.PushBackList(stack.List)
	return clone
}

func (stack *Stack) String() string {
	var output string
	for i, e := stack.List.Len(), stack.Back(); i > 0; i, e = i-1, e.Prev() {
		output = output + fmt.Sprintf("%v ", e.Value)
	}
	return output
}
