package codes

func One(nums *[]int){
	var n int = len(*nums)
	for i:=0;i<n/2 ;i++{
		(*nums)[i],(*nums)[n-i-1]=(*nums)[n-i-1],(*nums)[i]
	}
}