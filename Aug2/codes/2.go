package codes

import "fmt"

func Two(n int , array[] int) {
	mx:=0
	for i:=1;i<n;i++{
		mx=max(mx,array[i]+array[i-1])
	}
	fmt.Println(mx)
}