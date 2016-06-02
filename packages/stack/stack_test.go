package stack

import "testing"

func TestPushPop(t *testing.T) {
	c := new(Stack)
	c.Push(5)
	if c.Pop() != 5 {
		t.Log("Pop() does not give 5")
		t.Fail()
	}
}
