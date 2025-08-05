package codes 

func Three(nums1 []int,nums2 [] int) (nums []int){
	i:=0
	j:=0
	n1:=len(nums1)
	n2:=len(nums2)
	for i<n1 && j<n2{
		if nums1[i]<nums2[j]{
			nums=append(nums,nums1[i])
			i++
		}else{
			nums=append(nums,nums2[j])
			j++
		}
	}
	if i<n1{
		nums=append(nums, nums1[i:]...)
	}
	if j<n2{
		nums=append(nums, nums2[j:]...)
	}
	return 
}