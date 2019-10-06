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
	POPULATION = 17//33 // odd number to make room for the elite one.
	MUTATE_LIKELY = 6 // inverse
	GENERATIONS = 100000
	MAXERR = .75
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
	for i := 2; i < POPULATION; i++ {     // the number of breeders for which
		if (i*i-i)>>1 >= POPULATION { //the number of all combinations will equal POPULATION count.
			nBreeders = i
			break
		}
	}
	fmt.Println("nBreeders", nBreeders)
}

func threshold(pop *popu) int {// dud meter
	var nxLowest, mean int = 9999999, 0
	for _, v := range(pop) {
		mean += v.fitness
		if 0 < v.fitness && v.fitness < nxLowest {nxLowest = v.fitness}
	}
	mean /= len(pop)
	if mean == 0 {mean = nxLowest}
	return mean
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
	c := rand.Intn(len(GOAL)*MUTATE_LIKELY)
	if c < len(GOAL) {// if c < len(), mutate. Also use the random c to choose the target.
		i.chrome[c] = LETTERS[rand.Intn(len(LETTERS))]
	}
}

func crossover(a,b, oa, ob *individual) {// two-point. half of genes swapped
	c1 := rand.Intn(len(GOAL)>>1) // {0..len/2}
	c2 := c1+(len(GOAL)>>1)       // c1+len/2
	for j := range(GOAL) {
		if j< c1 ||  j > c2 {
			oa.chrome[j], ob.chrome[j] = b.chrome[j], a.chrome[j]
		} else {
			oa.chrome[j], ob.chrome[j] = a.chrome[j], b.chrome[j]
		}
	}
}

func regen(pop *popu) {
	var breeders []individual
	fitThresh:= threshold(pop)
	sort.Slice(pop[:], func(i, j int) bool {
	  return pop[i].fitness > pop[j].fitness
	})
	for i := 0; i < nBreeders; i++ {
		if pop[i].fitness < fitThresh {break}
		breeders = append(breeders, pop[i])
	}
	ipop := 1 // skip pop[0], the elite
	//if len(breeders) < 2 {goto xit}
	for  {// all populace
		for i := 0; i < len(breeders)-1; i++ {// crossover all combinations of breeders
			for j := i+1; j < len(breeders); j++ {
				crossover(&breeders[i], &breeders[i+1], &pop[ipop], &pop[ipop+1])
				ipop += 2
				if ipop >= len(pop)-1 {goto xit}
			}
		}
	}
	xit:
	for i, _ := range(pop)[1:] { mutate(&pop[i]) }// skip pop[0], the elite
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
