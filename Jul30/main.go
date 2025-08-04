package main

import (
	"fmt"
)

// Program 1
var x = 10
func main() {
	fmt.Println(x)
	x := 20 // Error: no new variables on left side of :=
	fmt.Println(x)
}


// Program 2
/*
const pi = 3.14
pi := 3.1342 
*/

// Program 3
/*
func main() {
	var s string
	var n int
	var b bool
	fmt.Printf("%q %d %t\n", s, n, b) 
}
*/

// Program 4
/*
func test() (a int, b int) {
	a = 1
	b = 2
	return
}

func main() {
	x, y := test()
	fmt.Println(x, y) 
}
*/

// Program 5
/*
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	data := []int{1, 2, 3}
	fmt.Println(sum(data...)) // Outputs: 6
}
*/

// Program 6
/*
const B = 3.13
var y int = B 
*/

// Program 7
/*
func main() {
	x := 42
	fmt.Println(x) 
}
*/

// Program 8
/*
const pi = 3.14

func main() {
	pi = 3.212 
}
*/

// Program 9
/*
func sayHi() {
	fmt.Println("Hi")
}

func main() {
	var f func() = sayHi
	f() // Outputs: Hi
}
*/


// Program 10

// func calc() (int, int) {
// 	return
// }

// func main() {
// 	a, b := calc()
// 	fmt.Println(a, b)
// }

