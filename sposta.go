package main

import (
	"fmt"
	"strconv"
)

type circNode struct {
	val        int
	next, prec *circNode
}

type listHead struct {
	head *circNode
}

func main() {
	var r string
	var l listHead
	for {
		contr, _ := fmt.Scanln(&r)
		if contr != 1 {
			break
		}

		valore, _ := strconv.Atoi(r)
		createNewNode(&l, valore)
	}

	fmt.Println("elementi:", f2(l.head, 5))
}

func sposta(p *circNode) {
	if p.val==0 {
		return
	}

	if p.val>0 {
		spostamenti:=p.val
		for i:=0; i<spostamenti; i++ {
			p.next.prec=p.prec
			p.prec.next=p.next
			p.prec=p.next
			p.next=p.next.next
		}
	} else {
		spostamenti:=-(p.val)
		for i:=0; i<spostamenti; i++ {
			p.prec.next=p.next
			p.next.prec=p.prec
			p.next=p.prec
			p.prec=p.prec.prec
		}
	}
}

func f2(p *circNode, k int) int {
	var counter int = 1
	var ricorda *circNode = p

	p=p.prec
	for {
		if counter==k {
			fmt.Println(p.val)
		}

		if p==ricorda {
			break
		}

		p=p.prec
		counter++
	}
	return counter
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
		nodo.next = lista.head
		nodo.prec = lista.head.prec
		nodo.prec.next = nodo
		lista.head.prec = nodo
		lista.head = nodo
	}
}

func newNode(val int) *circNode {
	node := new(circNode)
	node.val = val
	return node
}