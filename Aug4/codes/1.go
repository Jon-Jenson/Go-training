package codes

import (
	"fmt"
)

func inc(n *int) {
	*n++
}

func One() {
	x := 10
	inc(&x)
	fmt.Println(*(&x))
}