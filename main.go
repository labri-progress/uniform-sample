package main

import "flag"

func main() {
	number := flag.Int("t", 2, "Test number")
	flag.Parse()
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
