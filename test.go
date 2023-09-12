package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"network"
	"os"
	"strconv"
)

func test() {

	d := 50
	w := 40
	network.C = 5 //int(d)
	//network.Verbose = true
	network.PeerCMS = network.InitCMS(uint(d), uint(w))
	//fmt.Println(network.PeerCMS.MatToString())

	//var listpeers = []string{"A", "B", "C", "D", "E", "F", "G", "AB", "CD", "EF", "GH", "K"}
	var listpeers = []string{"AZ", "BY", "CXW", "DV", "EU", "FT", "GS", "HR", "IQ", "JP", "GH", "K"}

	for i, elm := range listpeers {
		println(">>>>>>>>>>>>> ELEMENT", i+1, " ", elm)
		//fmt.Println("output", output)

		network.PeerCMS.UpdateString(elm, 1)
		//fmt.Println(network.PeerCMS.MatToString())
	}
	for _, elm := range listpeers {
		freq := network.PeerCMS.EstimateString(elm)
		fmt.Printf("Element %s has freq %d\n", elm, freq)
	}
}

func summary(list []string) {
	var allelmt = make(map[string]uint)

	for _, elmt := range list {
		if _, ok := allelmt[elmt]; ok {
			allelmt[elmt] = allelmt[elmt] + 1
		} else {
			allelmt[elmt] = 1
		}
	}
	fmt.Println(allelmt)
}

func test0() {

	d := 50
	w := 40
	network.C = 10 //int(d)
	network.PeerCMS = network.InitCMS(uint(d), uint(w))
	fmt.Println(network.PeerCMS.MatToString())

	//network.Sample_memory = make([]string, network.C)
	//var listpeers = []string{"A", "B", "A", "A", "K", "S", "S", "B", "C", "A", "B", "K"}
	var listpeers = []string{"A", "B", "A", "A", "K", "S", "S", "B", "C", "A", "A", "B", "A", "A", "K", "S", "S", "B", "C", "A"}
	//var listpeers = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	var output = []string{}
	fmt.Println("Memory", network.Sample_memory)

	for i, elm := range listpeers {
		println(">>>>>>>>>>>>> ELEMENT", i+1, " ", elm)
		//fmt.Println("output", output)

		/* network.PeerCMS.UpdateString(elm, 1)
		//fmt.Println(network.PeerCMS.MatToString())

		value := network.PeerCMS.Knowledge_free(elm)
		output = append(output, value) */
		value := network.PeerCMS.Omniscient(elm)
		output = append(output, value)
		fmt.Println("Memory", network.Sample_memory)
		//fmt.Println("output", output)
	}
	fmt.Println("*******************output", output)

	fmt.Println("summary(listpeers)")
	summary(listpeers)
	fmt.Println("summary(output)")
	summary(output)
	fmt.Println("Memory", network.Sample_memory)
	for _, elm := range listpeers {
		freq := network.PeerCMS.EstimateString(elm)
		fmt.Printf("Element %s has freq %d\n", elm, freq)
	}
}

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

	//k := int(math.Ceil(0.01 * float64(n))) // other test to do
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

	out_path := "data/output" + strconv.Itoa(num_expe) // path of the unbiaised output stream
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

	out_path := "data/output" + strconv.Itoa(num_expe)
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
