package codes

import "fmt"

func Three(rows int , cols int , array [][]int){
	var nums []int
	flag:=false
	for i:=0;i<rows;i++{
		mn:=array[i][0]
		for j:=0;j<cols;j++{
			mn=min(mn,array[i][j])
		}
		nums=append(nums,mn)
	}

	for i:=0;i<cols;i++{
		mx:=array[0][i]
		for j:=0;j<rows;j++{
			mx=max(mx,array[j][i])
		}
		if mx==nums[i]{
			fmt.Println(mx)
			flag=true
			break
		}
	}
	if !flag{
		fmt.Println(-1)
	}

}