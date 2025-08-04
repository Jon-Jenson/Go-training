package codes


import "fmt"

func One(array []int,A int , N int){
	s:=0
	for i,v:=range array{
		if checkExist(v,A){
			s+=i
		}
	}
	fmt.Println(s)
}

func checkExist(num,dig int) bool{
	for num>0{
		d:=num/10
		if d==dig{
			return true 
		}
		num=num/10
	}
	
	return false
}