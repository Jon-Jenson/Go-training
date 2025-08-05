package codes

func Four(nums []int) (arr [] int){
	mp:=make(map[int]bool)
	for _,v :=range nums{
		if !mp[v]{
			arr=append(arr, v)
		}
		mp[v]=true

	}
	return
}