package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// MutationRate is the rate of mutation
var MutationRate = 0.005

// PopSize is the size of the population
var PopSize = 500

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	target := []byte("To be or not to be, that is the question.")
	population := createPopulation(target)

	found := false
	generation := 0
	for !found {
		generation++
		bestDNA := getBest(population)
		fmt.Printf("\r generation: %d | %s | fitness: %2f", generation, string(bestDNA.Gene), bestDNA.Fitness)

		if bytes.Compare(bestDNA.Gene, target) == 0 {
			found = true
		} else {
			maxFitness := bestDNA.Fitness
			pool := createPool(population, target, maxFitness)
			population = naturalSelection(pool, population, target)
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime taken: %s\n", elapsed)
}

// create the reproduction pool that creates the next generation
func createPool(population []DNA, target []byte, maxFitness float64) (pool []DNA) {
	pool = make([]DNA, 0)
	// create a pool for next generation
	for i := 0; i < len(population); i++ {
		population[i].calcFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

// perform natural selection to create the next generation
func naturalSelection(pool []DNA, population []DNA, target []byte) []DNA {
	next := make([]DNA, len(population))

	for i := 0; i < len(population); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]

		child := crossover(a, b)
		child.mutate()
		child.calcFitness(target)

		next[i] = child
	}
	return next
}

// creates the initial population
func createPopulation(target []byte) (population []DNA) {
	population = make([]DNA, PopSize)
	for i := 0; i < PopSize; i++ {
		population[i] = createDNA(target)
	}
	return
}

// Get the best gene
func getBest(population []DNA) DNA {
	best := 0.0
	index := 0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > best {
			index = i
			best = population[i].Fitness
		}
	}
	return population[index]
}

// DNA represents the genotype of the GA
type DNA struct {
	Gene    []byte
	Fitness float64
}

// generates a DNA string
func createDNA(target []byte) (dna DNA) {
	ba := make([]byte, len(target))
	for i := 0; i < len(target); i++ {
		ba[i] = byte(rand.Intn(95) + 32)
	}
	dna = DNA{
		Gene:    ba,
		Fitness: 0,
	}
	dna.calcFitness(target)
	return
}

// calculates the fitness of the DNA to the target string
func (d *DNA) calcFitness(target []byte) {
	score := 0
	for i := 0; i < len(d.Gene); i++ {
		if d.Gene[i] == target[i] {
			score++
		}
	}
	d.Fitness = float64(score) / float64(len(d.Gene))
	return
}

// crosses over 2 DNA strings
func crossover(d1 DNA, d2 DNA) DNA {
	child := DNA{
		Gene:    make([]byte, len(d1.Gene)),
		Fitness: 0,
	}
	mid := rand.Intn(len(d1.Gene))
	for i := 0; i < len(d1.Gene); i++ {
		if i > mid {
			child.Gene[i] = d1.Gene[i]
		} else {
			child.Gene[i] = d2.Gene[i]
		}

	}
	return child
}

// mutate the DNA string
func (d *DNA) mutate() {
	for i := 0; i < len(d.Gene); i++ {
		if rand.Float64() < MutationRate {
			d.Gene[i] = byte(rand.Intn(95) + 32)
		}
	}
}
