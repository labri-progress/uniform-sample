package main

import "flag"

var num_expe int
var input_path string

func main() {
	number := flag.Int("t", 2, "Test number")
	expe := flag.Int("n", 0, "experience number")
	path := flag.String("p", "data/resultJul95", "Input path")
	flag.Parse()
	print(expe)
	num_expe = *expe
	input_path = *path
	switch *number {
	case 0:
		//test0()
		check_cms()
	case 1:
		test_omniscient()
	case 2:
		test_Knowledge_free()
	default:
		panic("Wrong number")
	}

}
