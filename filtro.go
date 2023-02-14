package main

import (
	"fmt"
	"strconv"
)

type circNode struct {
	val        int
	next, prec *circNode
}

type node struct {
	val  int
	next *node
}

type listHead struct {
	head *circNode
}

type list struct {
	head *node
}

func main() {
	var r string
	var l listHead
	var fl list
	for {
		contr, _ := fmt.Scanln(&r)
		if contr != 1 {
			break
		}

		valore, _ := strconv.Atoi(r)
		createNewNode(&l, valore)
		addNewNodePointer(&fl, valore)
	}

	stampaDaZero(l.head)

	arara:=f(fl.head, 4)
	fmt.Println(arara)
}

func createNewNode(lista *listHead, n int) {
	// 0 elementi nella lista
	if lista.head == nil {
		lista.head = newNode(n)
		lista.head.next = lista.head
		lista.head.prec = lista.head
	} else {
		// 1 o più
		nodo := newNode(n)
		nodo.next = lista.head.next
		lista.head.next = nodo
		nodo.next.prec = nodo
		nodo.prec = lista.head
		lista.head = nodo
	}
}

func newNode(val int) *circNode {
	node := new(circNode)
	node.val = val
	return node
}

func stampaDaZero(p *circNode) {
	for p.val != 0 {
		p = p.next
	}

	fmt.Print(p.val, " ")
	p = p.next
	for p.val != 0 {
		fmt.Print(p.val, " ")
		p = p.next
	}
}

func f(p *node, k int) int {
	var a int
	if p == nil {
		return 0
	}

	a = 1 + f(p.next, k)

	if a == k {
		fmt.Println(p.val)
	}

	fmt.Println("quante volte", a)
	return a
}

func addNewNodePointer(list *list, valore int) {
	node := newwNode(valore)
	node.next = list.head
	list.head = node
}

//creo il nuovo nodo e ne inserisco i valori
func newwNode(valore int) *node {
	node := new(node)
	node.val = valore
	return node

	/*oppure possiamo fare così "return &listNode{val, chiav, nil}"*/
}
