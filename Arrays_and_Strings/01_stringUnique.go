package Arrays_and_Strings

import "fmt"

func FindUnique() {
	var boolList [128]bool
	result := true
	fmt.Println("Enter a string and I will tell you if this is unique")
	var option string
	fmt.Scanln(&option)
	fmt.Println("You have entered: " + option)
	for _, value := range option {
		fmt.Println(value)
		if boolList[value] {
			fmt.Println("This is already encountered")
			result = false
			break
		} else {
			boolList[value] = true
		}
	}
	if result {
		fmt.Println("Its a unique string")
	} else {
		fmt.Println("Not a unique string")
	}
}
