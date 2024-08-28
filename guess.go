package main

import(
	"fmt"
	"math/rand"
"time"
)

func main()  {
	fmt.Println("Game : Guess a number between 0 and 100")
	fmt.Println("You have three (3) tries")

	//? Give try count 
	var tryCount int
	fmt.Println("Please input your try count ")
	fmt.Scan(&tryCount)
	
	//? Generate a randome number
	source := rand.NewSource(time.Now().UnixNano())

	//? Generate numbers between 0 and 100 
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(100)


	var guess int

	for try := 1;try <= tryCount; try++{
		fmt.Printf("TRIAL %d\n",try)

		fmt.Println("Please enter your number ")

		fmt.Scan(&guess)

		if guess < secretNumber{
			fmt.Println("Sorry. wrong guess ; number is to small \n")
		}else if guess > secretNumber{
			fmt.Println("Sorry, wrong guess ; number is to large \n")
		}else{
			fmt.Println("Wou win! \n")
			break
		}

		if try == tryCount {
			fmt.Printf("Game over!! \n")
			fmt.Printf("The correct number is %d\n",secretNumber)
			break
		}
	}
}