package codes

import "fmt"

func Seven(n int){
	for i:=0;i<n;i++{
		if i%3==0 && i%5==0{
			fmt.Print("FizzBuzz, ")
		}else if i%3==0{
			fmt.Print("Fizz, ")
		}else if i%5==0{
			fmt.Print("Buzz, ")
		}else{
			fmt.Print(i,", ")
		}
	}
	fmt.Println()
}