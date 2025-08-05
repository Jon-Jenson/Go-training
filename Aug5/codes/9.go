package codes

import "fmt"

func Nine(day int){
	switch day{

	case 1:
		fmt.Print("Monday")

	case 2:
		fmt.Print("Tuesday")
		
	case 3:
		fmt.Print("Wednesday")

	case 4:
		fmt.Print("Thurdasy")
	
	case 5:
		fmt.Print("Friday")
	
	case 6:
		fmt.Print("Saturday")
	
	case 7:
		fmt.Print("Sunday")

	default:
		fmt.Print("Invalid day")
	}	
}