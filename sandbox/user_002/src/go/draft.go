package main

import "fmt"

type Node struct {
	info int
	next *Node
	prev *Node
}

type DList struct {
	head *Node
}

func NewDList() *DList {
	dlist := DList{}
	dlist.head = &Node{}
	dlist.head.next = dlist.head
	dlist.head.prev = dlist.head
	return &dlist
}

func insert(A *Node, value int) {
	B := A.prev
	C := &Node{
		info: value,
		next: A,
		prev: B,
	}
	B.next = C
	A.prev = C
}

func remove(B *Node){
    A := B.prev
    C := B.next

    A.next = C
    C.prev = A

    B.next = nil
    B.prev = nil
}

func (list *DList) PopBack(){ //remove o último da lista
    remove(list.head.prev)
}

func (list *DList) PopFront(){ //remove o primeiro da lista
    remove(list.head.next)
}

func (list *DList) PushFront(value int) { //insere na frente
	insert(list.head.next, value)
}

func (list *DList) PushBack(value int) { //insere atrás
	insert(list.head, value)
}

func (list *DList) Front() *Node { //primeiro elemento da lista válido
	return list.head.next
}

func (list *DList) Back() *Node { //ultimo elemento da lista válido
	return list.head.prev
}

func (list *DList) End() *Node { //depois do último valido (head)
	return list.head
}


func (list *DList) String() string { //toString
	saida := "[ "
	for it := list.Front(); it != list.End(); it = it.next{
		saida += fmt.Sprintf("%d ", it.info)
	}
	return saida + "]"
}

func main() {
	list := NewDList()
	for i := range 10 {
		list.PushBack(i)
	}

    list.PopBack()
	fmt.Println(list)
}