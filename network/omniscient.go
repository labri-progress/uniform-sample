package network

import (
	"fmt"
	"math/rand"
	"time"
)

var occurence = make(map[string]int)
var M int
var occ_min int

func Read_occurence(list []string) (n int) {
	fmt.Printf("length of the list is %d\n", len(list))
	for _, elmt := range list {
		if _, ok := occurence[elmt]; ok {
			occurence[elmt] = occurence[elmt] + 1
		} else {
			occurence[elmt] = 1
		}
	}
	occ_min = getMin()
	n = len(occurence)
	fmt.Printf("length of the occurence map is %d with min %d\n",
		n, occ_min)
	return
}
func getMin() (min int) {
	min = 0
	for _, j := range occurence {
		if j < min || min == 0 {
			if j > 0 {
				min = j
			}
		}
	}
	return
}

func (cms *CMS) Omniscient(peer string) (output_choice string) {
	var cmsRand = rand.New(
		rand.NewSource(time.Now().UnixNano())) /* RNG generator */

	if len(Sample_memory) < C {
		if !in_memory[peer] {
			Sample_memory = append(Sample_memory, peer) // add j
			in_memory[peer] = true
		}
	} else {
		min := occ_min          // min
		freq := occurence[peer] // pj

		prob := float64(min) / float64(freq) // aj <= 1
		choice := cmsRand.Float64()          // random choice in [0.0, 1.0[

		if choice < prob {
			sample_choice_index := cmsRand.Intn(C) //uniform random choice
			//Sample_memory[sample_choice_index] = peer // replace k by peer j
			if !in_memory[peer] {
				Sample_memory[sample_choice_index] = peer // add j
				in_memory[peer] = true
			}
		}
	}
	output_choice_index := cmsRand.Intn(len(Sample_memory)) //uniform random choice
	output_choice = Sample_memory[output_choice_index]      // k'

	return
}
