package codes

func Two(nums *[]int ,k int){
	k=k%len(*nums)
	(*nums)=append((*nums)[k:],(*nums)[0:k]... )
}