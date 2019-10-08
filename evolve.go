package main

/// By Joe Bumstead ___________ Licence: MIT

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

const (
	LETTERS = "abcdefghijklmnopqrstuvwxyz:' -;.,() ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	POPULATION = 21  // odd number to make room for the elite one.
	MUTATE_LIKELY = .27
	GENERATIONS = 77777
	MAXERR = .77
	GOAL = "There are only three kinds of people in the world: Those who can count and those who can't."
)

type individual struct {
	fitness int
	chrome [len(GOAL)]byte
}
type popu  [POPULATION]individual

var nBreeders int

func init() {
	fmt.Println("init, GOAL:", GOAL)
	rand.Seed(time.Now().UnixNano())
	for i := 2; i < POPULATION; i++ { // the number of breeders for which
		if (i*i-i)>>1 > POPULATION { //the number of all combinations will equal POPULATION count.
			nBreeders = i
			break
		}
	}
	fmt.Println("nBreeders", nBreeders)
}

func fitness(pop *popu) (e float32) {// the number of chars in individual that match chars in GOAL.
	for j, _ := range(pop) {
		pop[j].fitness = 0
		for i := 0; i < len(GOAL); i++ {
			if pop[j].chrome[i] == GOAL[i] {pop[j].fitness++}
		}
		e += float32(pop[j].fitness)
	}
	return float32(len(GOAL))-e/float32(len(pop)) // average error count
}

func mutate(i *individual) {
	if rand.Float32() > MUTATE_LIKELY {return}
	c := rand.Intn(len(GOAL))
	i.chrome[c] = LETTERS[rand.Intn(len(LETTERS))]
}

func crossover(a,b, oa, ob *individual) {// two-point. half of genes swapped
	c1 := rand.Intn(len(GOAL)>>1) // {0..len/2}
	c2 := c1+(len(GOAL)>>1) 		// c1+len/2
	for j := range(GOAL) {
		if c1 <= j &&  j < c2 {
			oa.chrome[j], ob.chrome[j] = b.chrome[j], a.chrome[j]
		} else {
			oa.chrome[j], ob.chrome[j] = a.chrome[j], b.chrome[j]
		}
	}
}

func regen(pop *popu) {
	var breeders []individual
	sort.Slice(pop[:], func(i, j int) bool {
	  return pop[i].fitness > pop[j].fitness
	})
	breeders = append(pop[:0:0], pop[0:nBreeders]...)
	ipop := 1 // skip pop[0], the elite
	for  {// all populace
		for i := 0; i < len(breeders)-1; i++ {// crossover all combinations of breeders
			for j := i+1; j < len(breeders); j++ {
				crossover(&breeders[i], &breeders[j], &pop[ipop], &pop[ipop+1])
				ipop += 2
				if ipop >= len(pop)-1 {goto xit}
			}
		}
	}
	xit:
	for i := 1; i < POPULATION; i++ { mutate(&pop[i]) }// skip pop[0], the elite
}

func randChrome() *[len(GOAL)]byte {
	var i [len(GOAL)]byte
	for j := range(GOAL) { i[j] = LETTERS[rand.Intn(len(LETTERS))] }
	return &i
}

func main() {
	fmt.Println("____________________ START ____________________")
	var populace  popu
	for i := range(populace) { populace[i].chrome = *randChrome() }
	
	for i := 0; i < GENERATIONS; i++ {
		if fitness(&populace) <= MAXERR { fmt.Println("Early fit; Generations:", i); break }// early out?
		regen(&populace)
	}
	for i := 0; i < POPULATION; i++ {// show populace
		fmt.Println(string(populace[i].chrome[:]), "      ", len(GOAL)-populace[i].fitness, "errors.")
	}
}
