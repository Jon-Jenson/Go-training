package codes

func Five(s string)(mp map[rune]int){
	mp = make(map[rune]int)
	for _,v := range s{
		mp[v] +=1 
	}
	return
}