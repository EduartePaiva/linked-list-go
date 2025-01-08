package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// prints the list
func (lst *List[T]) String() string {
	if lst == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{val: %v, next: %s}", lst.val, lst.next)
}

func main() {
	head := new(List[string])
	head.val = "A"

	lst := head
	for i := 1; i < 26; i++ {
		new_lst := new(List[string])
		new_lst.val = string(('A' + byte(i)))
		lst.next = new_lst
		lst = lst.next
	}

	fmt.Println(head)
}
