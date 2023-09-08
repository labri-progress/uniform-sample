package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"network"
	"os"
)

func test_Knowledge_free() {

	path := "data/resultJul95"
	var listpeers, err = readLines(path)
	if err != nil {
		log.Println("(check) Unable to read config file ", path)
		return
	}

	m := len(listpeers)
	log.Println("Input of size ", m)

	n := network.Read_occurence(listpeers)
	/* CMS Parameters */

	//k := 0.01 * float64(N) // other test to do
	k := int(math.Ceil(math.Log(float64(n)))) // round to the next integer
	s := 10
	network.C = k

	network.PeerCMS = network.InitCMS(uint(s), uint(k))

	/*  Knowledge free algorithm */
	var output = []string{}

	for _, elm := range listpeers {
		//println(">elmt", i)
		network.PeerCMS.UpdateString(elm, 1)

		value := network.PeerCMS.Knowledge_free(elm)
		output = append(output, value)

	}

	/* summary */
	fmt.Printf("A matrix of size %d*%d with sample memory length of %d\n Output of size %d\n",
		s, k, len(network.Sample_memory), len(output))

	out_path := "data/output" // path of the unbiaised output stream
	err = writeLines(output, out_path)
	if err != nil {
		log.Println("(Write) Unable to read config file ", out_path)
		return
	}

	return
}

func test_omniscient() {

	network.C = 300

	path := "data/resultJul95"
	var listpeers, err = readLines(path)
	if err != nil {
		log.Println("(check) Unable to read config file ", path)
		return
	}
	network.Read_occurence(listpeers)

	var output = []string{}

	for _, elm := range listpeers {
		//println(">elmt", i)

		value := network.PeerCMS.Omniscient(elm)
		output = append(output, value)
		//fmt.Println("output", output)
	}

	out_path := "data/output"
	err = writeLines(output, out_path)
	if err != nil {
		log.Println("(Write) Unable to read config file ", out_path)
		return
	}

	return
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
