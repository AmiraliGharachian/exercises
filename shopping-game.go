package main

import "fmt"

func main() {

	var budget int
	fmt.Print("Enter your budget: ")
	_, err := fmt.Scan(&budget)

	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	starcraftprice := 30
	witcherprice := 60
	cyberpunkprice := 130

	boughtGames := ""

	if budget >= starcraftprice {
		boughtGames += "StarCraft 2"
		budget -= starcraftprice
	}

	if budget >= witcherprice {
		if boughtGames != "" {
			boughtGames += ", "
		}
		boughtGames += "The Witcher 3"
		budget -= witcherprice
	}

	if budget >= cyberpunkprice {
		if boughtGames != "" {
			boughtGames += ", "
		}
		boughtGames += "Cyberpunk 2077"
		budget -= cyberpunkprice
	}

	if boughtGames != "" {
		fmt.Println("Here's what I bought:", boughtGames)
	} else {
		fmt.Println("I couldn't buy anything!")
	}
}
