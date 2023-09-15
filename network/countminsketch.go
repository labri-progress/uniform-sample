package network

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"log"
	"math/rand"
	"time"
)

type CMS struct {
	s      uint // depth: number of hash functions
	k      uint // width: number of colums : s < k
	mat    [][]uint64
	hasher hash.Hash64
}

var PeerCMS *CMS
var Sample_memory = []string{}
var in_memory = make(map[string]bool)
var C int

func InitCMS(d uint, w uint) (cms *CMS) {

	log.Printf("Initializing the CMS with %d*%d\n", d, w)

	cms, err := New(d, w)
	if err != nil {
		log.Panic("error when initializing the CMS", err)
		return
	}
	return
}
func (cms *CMS) getMinMatrix() (min uint64) {
	min = 0
	for i := uint(0); i < cms.s; i++ {
		for j := uint(0); j < cms.k; j++ {
			if cms.mat[i][j] < min || min == 0 {
				if cms.mat[i][j] > 0 {
					min = cms.mat[i][j]
				}
			}
		}
	}
	return
}

func getMinFreq() (min uint64) {
	min = 0
	for i := range in_memory {
		freq := PeerCMS.EstimateString(i)
		if freq < min || min == 0 {
			min = freq
		}
	}
	return
}

func (cms *CMS) Knowledge_free(peer string) (output_choice string) {
	var cmsRand = rand.New(
		rand.NewSource(time.Now().UnixNano())) /* RNG generator */

	if len(Sample_memory) < C {
		if !in_memory[peer] {
			Sample_memory = append(Sample_memory, peer) // add j
			in_memory[peer] = true
		}

	} else {

		var prob float64
		if !in_memory[peer] {
			prob = 1.0
		} else {
			//min := cms.getMinMatrix()
			min := getMinFreq()
			freq := PeerCMS.EstimateString(peer)
			if freq == 0 {
				log.Panicf("Element %s has not been registered yet !!!\n", peer)
			}
			prob = float64(min) / float64(freq) // aj <= 1
			//fmt.Println("min", min, "prob", prob)
		}

		choice := cmsRand.Float64() // random choice in [0.0, 1.0[
		//fmt.Println("choice", choice)
		if choice < prob && !in_memory[peer] { // try to add only not known elements

			sample_choice_index := cmsRand.Intn(C) //uniform random choice
			k := Sample_memory[sample_choice_index]
			delete(in_memory, k)
			Sample_memory[sample_choice_index] = peer // add j
			in_memory[peer] = true
			//fmt.Println("change", k, peer)
		}
	}
	//fmt.Println("Memory", Sample_memory)
	output_choice_index := cmsRand.Intn(len(Sample_memory)) //uniform random choice
	output_choice = Sample_memory[output_choice_index]      // k'
	//fmt.Printf("Sample memory : %v \n in_memory: %v\n", Sample_memory, in_memory)
	return
}

// New CMS
func New(d uint, w uint) (c *CMS, err error) {
	if d <= 0 || w <= 0 {
		return nil, errors.New("(countminsketch) values of d and w should both be greater than 0")
	}

	c = &CMS{
		s:      d,
		k:      w,
		hasher: fnv.New64(),
	}
	c.mat = make([][]uint64, d)
	for r := uint(0); r < d; r++ {
		c.mat[r] = make([]uint64, w)
	}

	return c, nil
}

// get the two basic hash function values for data.
func (s *CMS) BaseHashes(key []byte) (a uint32, b uint32) {
	s.hasher.Reset()
	/* newkey := string(key)
	sum := s.hasher.Sum([]byte(newkey)) */
	s.hasher.Write(key)
	sum := s.hasher.Sum(nil)
	//fmt.Printf("sum=%d\n", sum)
	upper := sum[0:4]
	lower := sum[4:8]

	//fmt.Printf("upper=%d and lower=%d\n", upper, lower)
	a = binary.BigEndian.Uint32(lower)
	b = binary.BigEndian.Uint32(upper)

	//fmt.Printf("a=%d and b=%d\n", a, b)
	return
}

// Get the _w_ locations to update/Estimate
func (s *CMS) Locations(key []byte) (locs []uint) {
	locs = make([]uint, s.s)
	a, b := s.BaseHashes(key)
	ua := uint(a)
	ub := uint(b)
	//fmt.Printf("ua=%d and ub=%d\n", ua, ub)
	for r := uint(0); r < s.s; r++ {
		locs[r] = (ua + ub*r) % s.k
	}
	//fmt.Printf("locs=%v\n", locs)
	return
}

// Update the frequency of a key
func (s *CMS) Update(key []byte, count uint64) {
	for r, c := range s.Locations(key) {
		s.mat[r][c] += count
	}
}

// UpdateString updates the frequency of a key with a string parameter
func (s *CMS) UpdateString(key string, count uint64) {
	s.Update([]byte(key), count)
}

// Estimate the frequency of a key. It is point query.
func (s *CMS) Estimate(key []byte) uint64 {
	var min uint64
	for r, c := range s.Locations(key) {
		if r == 0 || s.mat[r][c] < min {
			min = s.mat[r][c]
		}
	}
	return min
}

// EstimateString estimate the frequency of a key of string
func (s *CMS) EstimateString(key string) uint64 {
	return s.Estimate([]byte(key))
}

// Merge combines this CountMinSketch with another one
func (s *CMS) Merge(other *CMS) error {
	if s.s != other.s {
		return errors.New("countminsketch: matrix depth must match")
	}

	if s.k != other.k {
		return errors.New("countminsketch: matrix width must match")
	}

	for i := uint(0); i < s.s; i++ {
		for j := uint(0); j < s.k; j++ {
			// s.mat[i][j] += other.mat[i][j] ADD
			var val float64
			val = (float64(s.mat[i][j]) + float64(other.mat[i][j])) / 2 /* MOY */
			if 0.0 < val && val < 1.0 {
				s.mat[i][j] = 1
			} else {
				s.mat[i][j] = uint64(val)
			}
		}
	}

	return nil
}

// Transforms the matrix of the CMS into a string
func (cms *CMS) MatToString() (str string) {
	var s string
	for i := range cms.mat {
		for _, n := range cms.mat[i] {
			s += fmt.Sprintf("%d ", n)
		}
		s += "| "
	}
	return s
}
