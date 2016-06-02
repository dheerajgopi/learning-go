/* A simple stack which can push and pop ints */
package stack

// Stack for storing ints
type Stack struct {
	i    int
	data [10]int
}

// An int is pushed to the stack
func (s *Stack) Push(n int) {
	s.data[s.i] = n
	s.i++
}

// An int is popped from the stack and the popped int is returned
func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}
