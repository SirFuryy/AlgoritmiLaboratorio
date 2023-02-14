package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type oggetto struct {
	nome string
	val  int // non rilevante se l'indizio e' un'operazione
	dx   string
	op   rune // " " se numero
	sx   string
	tipo string // "num" se l'indizio è un numero, "op" se l'indizio è una operazione
}

func leggiInput() map[string]*oggetto {
	scanner := bufio.NewScanner(os.Stdin)
	fogli := make(map[string]*oggetto)

	for scanner.Scan() {
		parole := strings.Split(scanner.Text(), " ")

		nome := strings.Trim(parole[0], ":")
		if len(parole) == 2 {
			numero, _ := strconv.Atoi(parole[1])
			fogli[nome] = &oggetto{nome, numero, "", ' ', "", "num"}
		} else if len(parole) == 4 {
			fogli[nome] = &oggetto{nome, 0, parole[3], rune(parole[2][0]), parole[1], "op"}
		} else {
			fmt.Println("Errore")
		}
	}

	return fogli
}

type albero struct {
	data       map[string]*oggetto
	entryTesta string
}

func calcolaPrezzo(f foresta, n string) int {
	if f.dati[n].tipo == "num" {
		return f.dati[n].val
	}
	
	switch f.dati[n].op {
	case '+':
		return calcolaPrezzo(f, f.dati[n].sx) + calcolaPrezzo(f, f.dati[n].dx)
	case '-':
		return calcolaPrezzo(f, f.dati[n].sx) - calcolaPrezzo(f, f.dati[n].dx)
	case '*':
		return calcolaPrezzo(f, f.dati[n].sx) * calcolaPrezzo(f, f.dati[n].dx)
	case '/':
		return calcolaPrezzo(f, f.dati[n].sx) / calcolaPrezzo(f, f.dati[n].dx)
	}

	return 0
}

func (ab *albero) peso(nodo string) int {
	node := ab.data[nodo]
	if node == nil {
		return 0
	}

	if node.tipo == "num" {
		return node.val
	} else {
		switch node.op {
		case '+':
			return ab.peso(node.sx) + ab.peso(node.dx)
		case '-':
			return ab.peso(node.sx) - ab.peso(node.dx)
		case '*':
			return ab.peso(node.sx) * ab.peso(node.dx)
		case '/':
			return ab.peso(node.sx) / ab.peso(node.dx)
		}

	}
	return 0
}

type foresta struct {
	dati  map[string]*oggetto
	entry []string
}

func costruisciForesta(mappa map[string]*oggetto) foresta {
	var f foresta
	radix := make(map[string]bool) // True se è ha un padre, false se non ce l'ha
	f.dati = mappa
	for k := range mappa {
		radix[k] = false
	}
	for _, v := range mappa {
		radix[v.dx] = true
		radix[v.sx] = true
	}
	for k, v := range radix {
		if !v {
			f.entry = append(f.entry, k)
		}
	}
	return f
}

func stampaAlbero(f foresta, n string) {
	sx := f.dati[n].sx
	dx := f.dati[n].dx
	if f.dati[n].tipo == "num" {
		fmt.Print(f.dati[n].nome, " (val = ", f.dati[n].val, ")\n")
	} else {
		stampaAlbero(f, sx)
		fmt.Println(f.dati[n].nome)
		stampaAlbero(f, dx)
	}
}

func sx(f foresta, n string) (string, bool) {
	return f.dati[n].sx, true
}

func dx(f foresta, n string) (string, bool) {
	return f.dati[n].dx, true
}

func up(f foresta, n string) (string, bool) {
	for _, v := range f.dati {
		if v.dx == n || v.sx == n {
			return v.nome, true
		}
	}
	return "", false
}

// non modificare
func main() {
	mappa := leggiInput()
	f := costruisciForesta(mappa)
	stampaAlbero(f, "letto")
}
