package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (lst *List[T]) String() string {
	if lst == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{val: %v, next: %s}", lst.val, lst.next)
}
func createList[E any](val E) *List[E] {
	new_lst := new(List[E])
	new_lst.val = val
	return new_lst
}

func main() {
	head := createList("A")

	lst := head
	for i := 1; i < 26; i++ {
		new_lst := createList(string(('A' + byte(i))))
		lst.next = new_lst
		lst = lst.next
	}

	fmt.Println(head)
}
