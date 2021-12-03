package main

func main() {

	// datatype to hold each chunck of this pretend binary
	// I will make an array of theese so it can be accessed in this way:
	//
	// [n] abcde
	// ---------
	// [0] 00100
	// [1] 11110
	// [2] 10110
	//
	// to access leftmost "bit" of the first "chunck" bin[0].a
	type bin struct {
		a int
		b int
		c int
		d int
		e int
	}

}
