/* A simple stack which can push and pop ints */
package stack

// Stack for storing ints
type Stack struct {
	Height int
	Data   [10]int
}

// An int is pushed to the stack
func (s *Stack) Push(n int) {
	s.Data[s.Height] = n
	s.Height++
}

// An int is popped from the stack and the popped int is returned
func (s *Stack) Pop() int {
	s.Height--
	return s.Data[s.Height]
}
