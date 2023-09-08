package main

import "flag"

func main() {
	number := flag.Int("t", 2, "Test number")
	flag.Parse()
	switch *number {
	case 0:
		test_Knowledge_free()
	case 1:
		test_omniscient()
	default:
		panic("Wrong number")
	}
}
