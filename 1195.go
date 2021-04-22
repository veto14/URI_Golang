package main

import (
	"fmt"
)

type Arvore struct {
	raiz *No
}

type No struct {
	chave byte
	esq   *No
	dir   *No
}

func (n *No) insere(data byte) {
	if data <= n.chave {
		if n.esq == nil {
			n.esq = &No{chave: data}
		} else {
			n.esq.insere(data)
		}
	} else {
		if n.dir == nil {
			n.dir = &No{chave: data}
		} else {
			n.dir.insere(data)
		}
	}
}

func (t *Arvore) insere(data byte) {
	if t.raiz == nil {
		t.raiz = &No{chave: data}
	} else {
		t.raiz.insere(data)
	}
}

func prefixa(n *No) {
	if n == nil {
		return
	} else {
		fmt.Printf(" ")
		fmt.Printf("%d", n.chave)
		prefixa(n.esq)
		prefixa(n.dir)
	}
}

func posfixa(n *No) {
	if n == nil {
		return
	} else {
		posfixa(n.esq)
		posfixa(n.dir)
		fmt.Printf(" ")
		fmt.Printf("%d", n.chave)
	}
}

func infixa(n *No) {
	if n == nil {
		return
	} else {
		posfixa(n.esq)
		fmt.Printf(" ")
		fmt.Printf("%d", n.chave)
		posfixa(n.dir)
	}
}

func main() {
	var casos, no int
	fmt.Scan(&casos)
	for i := 0; i < casos; i++ {
		var t Arvore
		var aux byte
		fmt.Scan(&no)
		for j := 0; j < no; j++ {
			fmt.Scan(&aux)
			t.insere(aux)
		}
		fmt.Println("Case 1:")
		fmt.Print("Pre.:")
		prefixa(t.raiz)
		fmt.Printf("\n")
		fmt.Println("Case 2")
		fmt.Print("In.:")
		infixa(t.raiz)
		fmt.Printf("\n")
		fmt.Println("Case 3:")
		fmt.Print("Post.:")
		posfixa(t.raiz)
		fmt.Printf("\n\n")
	}
}
