package codes
import "fmt"

func Four(n int , L int , R int , array[] int) {
	Ls:=0
	var window_number int 
	window_number=0
	
	for i:=0;i<L ; i++{
		Ls+=array[i]
	}

	mx:=Ls
	for i:=L;i<n;i++{
		Ls-=array[i-L]
		Ls+=array[i]
		if Ls>mx{
			window_number=i-L+1
			mx=Ls
		}
	}
	var Rs int 
	size:=0
	Rmx:=0
	for i:=0;i<n;i++{
		if i>=window_number && i<window_number+L{
			Rs=0
			size=0
			continue
		}
		Rs+=array[i]
		size++
		if size==R+1{
			size--
			Rs-=array[i-R]
		}
		Rmx=max(Rmx,Rs)
	}
	fmt.Println(Rmx+mx)
}