package render

import "fmt"

func GetNumberOfColumns(numbersOfColumns *int) {
	for {
		fmt.Print(Bold + "Write number of columns: ")
		fmt.Scan(numbersOfColumns)
		if *numbersOfColumns%2 == 0 || *numbersOfColumns < 2 {
			fmt.Println(Red + "Error: number of columns must be not even and greater than 2" + Reset + Bold)
		} else if *numbersOfColumns > 30 {
			var answer string
			fmt.Print("Are you sure?(y/n): ")
			fmt.Scan(&answer)
			if answer == "y" {
				fmt.Print("Okay :) Resize your window and enter any symbol: ")
				fmt.Scan(&answer)
				break
			}
		} else {
			break
		}
	}
	fmt.Print(Reset)
}
