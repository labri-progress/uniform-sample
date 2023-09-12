package main

import "flag"

var num_expe int

func main() {
	number := flag.Int("t", 2, "Test number")
	expe := flag.Int("n", 0, "experience number")
	flag.Parse()

	num_expe = *expe

	switch *number {
	case 0:
		test0()
	case 1:
		test_omniscient()
	case 2:
		test_Knowledge_free()
	default:
		panic("Wrong number")
	}

}
