package codes

import "math"

func Eight(num int) bool{
	if num<=1{
		return false
	}
	
	lim:=int(math.Sqrt(float64(num)))
	for i:=2;i<=lim;i++{
		if num%i==0{
			return false
		}
	}
	return true
}