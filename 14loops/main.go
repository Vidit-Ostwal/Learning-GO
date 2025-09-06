package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Println(index, day)
	}

	rougurValue := 1

	for rougurValue < 10 {
		if rougurValue == 5 {
			rougurValue++
		}

		fmt.Println("Value is: ", rougurValue)
		rougurValue++
	}
}
