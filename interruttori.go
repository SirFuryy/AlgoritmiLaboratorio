package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type conf []bool //tipo che rappresenta lo stato di accensione delle luci

//Dato un intero x, modifica lo stato di accensione delle luci passate in input in base al pulsante premuto
func premi(curr conf, x int) conf {
	switch x {
	case 0:
		curr[len(curr)-1] = !curr[len(curr)-1]
		curr[x] = !curr[x]
		curr[x+1] = !curr[x+1]
	case len(curr) - 1:
		curr[x-1] = !curr[x-1]
		curr[x] = !curr[x]
		curr[0] = !curr[0]
	default:
		curr[x-1] = !curr[x-1]
		curr[x] = !curr[x]
		curr[x+1] = !curr[x+1]
	}

	return curr
}

//Data una sequenza x di interruttori da premere, modifica lo stato di accensione delle luci passate in input in base ai pulsanti premuti
func sequenza(curr conf, x []int) conf {
	for _, v := range x {
		curr = premi(curr, v)
	}

	return curr
}

//Data uno stato di accensione delle luci, ritorna una sequenza di interruttori da premere per spegnere tutte le luci
func sequenzaSpegniTutto(curr conf) []int {
	/*
		L'idea è quella di eseguire una bfs per trovare il nodo spento all'interno del grafo di tutti le possibili combinazioni, considerando che per
		ogni combinazione potevamo premere n tasti diversi generando n figli diversi. Il path effettivamente utile alla fine sarà quelo di una dfs, ma
		questa parte la realiziamo con un altra funzione che controlla i padri di ogni nodo e dalla fine ritorna all'inizio. Non è ottimizzata perchè
		non considera eventuali loop di nodi già visitati, ma non ho voglia di aggiungere una eql
	*/

	frontiera := make([]conf, 0) //CODA FIFO PER BFS
	padri := make([]conf, 0)

	//primo caso
	frontiera = append(frontiera, curr)
	padri = append(padri, nil)

	for len(frontiera) > 0 {
		nodo := frontiera[0]
		frontiera = frontiera[1:] //in teoria eliminiamo il nodo

		for c := range curr { //si intende quanti tasti possibili ci sono da premere
			nuovocurr := premi(nodo, c)
			frontiera = append(frontiera, nuovocurr)
			padri = append(padri, nodo)

			if spento(nuovocurr) {
				return variazioni(percMinimo(frontiera, padri))
			}

			premi(nodo, c) //riporto allo stato di accensione di nodo
		}
	}

	return nil
}

//ritorna una configurazione che contiene i passaggi che effettivamente portano allo stato spento, però lavorando alla rovescia cercando ciascun padre di ogni elemento che porta allo stato finale
func percMinimo(sequenza []conf, padri []conf) []conf {
	seqContr := make([]conf, 0)

	//primo padre
	padre := padri[len(padri)-1]
	seqContr = append(seqContr, padre)

	//ricerchiamo all'indietro i padri
	for padre != nil {
		for c, v := range sequenza {
			if uguali(padre, v) {
				padre = padri[c]
				seqContr = append(seqContr, padre)
			}
		}
	}

	//riordiniamo in avanti la sequenza
	seq := make([]conf, 0)
	for i := len(seqContr) - 1; i >= 0; i++ {
		seq = append(seq, seqContr[i])
	}

	return seq
}

//trasforma la sequenza di stati minimi in una sequenza di bottoni da premere
func variazioni(seq []conf) []int {
	seqTasti := make([]int, 0)

	/*
		qua l'idea è quella di controllare se premendo un tasto ho la configurazione successiva, se si è il tasto giusto, se no proseguo
	*/
	for i := 1; i < len(seq); i++ {
		for c := range seq[i] {
			ris := premi(seq[i-1], c)
			if uguali(seq[i], ris) {
				seqTasti = append(seqTasti, c)
			}

			premi(seq[i-1], c) //risistema lo stato delle luci
		}
	}

	return seqTasti
}

//date due configurazioni, controlla che esse siano uguali vagliando che tutti gli elementi siano uguali
func uguali(conf1, conf2 conf) bool {
	for c, v := range conf1 {
		if v != conf2[c] {
			return false
		}
	}

	return true
}

//data una sequenza di partenza, ritorna il numero di pulsanti da premere per arrivare allo stato di spento
func spegniTutto(curr conf) int {
	return len(sequenzaSpegniTutto(curr))
}

//data una configurazione, ritorna se essa è completamente spenta (tutti i valori sono settati false)
func spento(curr conf) bool {
	for _, v := range curr {
		if v {
			return false
		}
	}

	return true
}

func main() {
	rand.Seed(time.Now().Unix())
	var riga string
	var rete conf = make(conf, 0)

	for {
		fmt.Scanln(&riga)

		parole := strings.Split(riga, " ")
		var interrompi bool

		switch parole[0] {
		case "+": //crea rete
			nluci, _ := strconv.Atoi(parole[1])
			if len(rete) != 0 { //se già esistente la ricreo
				for c := range rete {
					rete[c] = false
				}

				rete = rete[:nluci]
			} else { //altrimenti la creo e basta
				for i := 0; i < nluci; i++ {
					rete = append(rete, false)
				}
			}

		case "o": //imposta la rete a questo stato
			for c, v := range parole[1] {
				if v == '1' {
					rete[c] = true
				}
			}

		case "p": //stampa lo stato della rete
			fmt.Println(rete)

		case "s": //preme l'interruttore
			tasto, _ := strconv.Atoi(parole[1])
			rete = premi(rete, tasto)

		case "S": //preme la sequenza di interruttori
			seq := make([]int, 0)
			for _, v := range parole[1] {
				num, _ := strconv.Atoi(string(v))
				seq = append(seq, num)
			}
			rete = sequenza(rete, seq)

		case "x": //stampa il numero minimo di pulsanti per spegnere tutto
			fmt.Println(spegniTutto(rete))

		case "f": //interrompe il ciclo
			interrompi = true
		}

		if interrompi {
			break
		}
	}
}
