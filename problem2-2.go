package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Enter command: ")
		scanner.Scan()
		input := scanner.Text()
		commands := strings.Fields(input)

		if len(commands) == 0 {
			continue
		}
		command := commands[0]

		switch command {
		case "add":

			if len(commands) != 2 {
				fmt.Println("Invalid command")
				continue
			}
			num, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("Invalid number")
				continue
			}
			numbers = append(numbers, num)
			fmt.Println("OK")

		case "inc":

			if len(commands) != 2 {
				fmt.Println("Invalid command")
				continue
			}
			index, err := strconv.Atoi(commands[1])
			if err != nil || index < 0 || index >= len(numbers) {
				fmt.Println("invalid index")
				continue
			}
			numbers[index]++
			fmt.Println("OK")

		case "sub":

			if len(commands) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			index, err := strconv.Atoi(commands[1])
			num, err2 := strconv.Atoi(commands[2])
			if err != nil || err2 != nil || index < 0 || index >= len(numbers) {
				fmt.Println("invalid index")
				continue
			}
			numbers[index] -= num
			fmt.Println("OK")

		case "mul":

			if len(commands) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			index, err := strconv.Atoi(commands[1])
			num, err2 := strconv.Atoi(commands[2])
			if err != nil || err2 != nil || index < 0 || index >= len(numbers) {
				fmt.Println("invalid index")
				continue
			}
			numbers[index] *= num
			fmt.Println("OK")

		case "acc":

			if len(commands) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			index, err := strconv.Atoi(commands[1])
			num, err2 := strconv.Atoi(commands[2])
			if err != nil || err2 != nil || index < 0 || index >= len(numbers) {
				fmt.Println("invalid index")
				continue
			}
			numbers[index] += num
			fmt.Println("OK")

		case "show":

			for _, num := range numbers {
				fmt.Printf("%d ", num)
			}
			fmt.Println()

		case "exit":

			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid command")
		}
	}
}
