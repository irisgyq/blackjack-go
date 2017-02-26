package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
)

var (
	random = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {

	initcards := []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 9,
		10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

	playerCard := make([]int, 0)
	dealerCard := make([]int, 0)

	PisBust := false
	DisBust := false

	fmt.Println("Game begins...")
	cards := shuffle(initcards)

	playerCard = append(playerCard, pop(&cards))
	fmt.Print("The player's first card is: ")
	fmt.Println(playerCard[0])
	dealerCard = append(dealerCard, pop(&cards))
	fmt.Print("The dealer's first card is: ")
	fmt.Println(dealerCard[0])
	playerCard = append(playerCard, pop(&cards))
	fmt.Print("The player's second card is: ")
	fmt.Println(playerCard[1])

	if blackjack(playerCard) {
		fmt.Println("The player has blackjack!")
		fmt.Println("Game is over, the player is the winner.")
	} else {

		dealerCard = append(dealerCard, pop(&cards))

		if blackjack(dealerCard) {
			fmt.Print("The dealer's second card is: ")
			fmt.Println(dealerCard[1])
			fmt.Println("The dealer has blackjack!")
			fmt.Println("Game is over, the dealer is the winner.")

		} else {

			playerSum := playerCard[0] + playerCard[1]
			dealerSum := dealerCard[0] + dealerCard[1]

			fmt.Print("The sum of player's cards is:")
			fmt.Println(playerSum)

			isPValid := true
			for (isPValid) {
				fmt.Println("Does the player want one more card?")
				inputReader := bufio.NewReader(os.Stdin)
				input, err := inputReader.ReadString('\n')

				if err != nil {
					fmt.Println("Your input is wrong.")
					return
				}

				switch input {
				case "yes\n":{
					playerCard = append(playerCard, pop(&cards))
					fmt.Print("This card is:")
					fmt.Println(playerCard[len(playerCard) - 1])
					playerSum += playerCard[len(playerCard) - 1]
					fmt.Print("The sum of player's card is:")
					fmt.Println(playerSum)

					if blackjack(playerCard) {
						fmt.Println("The player has 21 points!")
						isPValid = false
						break
					} else if playerSum > 21 {
						fmt.Println("Player's cards are busting.")
						PisBust = true
						isPValid = false
						break
					}
					break
				}
				case "hit\n" :{
					isPValid = false
					break
				}

				}
			}

			fmt.Print("The dealer's second card is: ")
			fmt.Println(dealerCard[1])
			fmt.Print("The sum of dealer's cards is:")
			fmt.Println(dealerSum)

			isDValid := true
			for (isDValid) {
				for dealerSum < 17 {
					fmt.Println("Because the sum of dealer's cards is less than 17, he must add one more card.")
					dealerCard = append(dealerCard, pop(&cards))
					fmt.Print("The new card is:")
					fmt.Println(dealerCard[len(dealerCard) - 1])
					dealerSum += dealerCard[len(dealerCard) - 1]
					fmt.Print("The sum of dealer's cards is:")
					fmt.Println(dealerSum)

					if dealerSum == 21 {
						fmt.Println("The dealer has 21 points!")
						break
					}
					if dealerSum > 21 {
						fmt.Println("dealer's cards are busting.")
						DisBust = true
						break
					}
				}

				if dealerSum < 21 {
					fmt.Println("Dose the dealer want one more card?")
					inputReader := bufio.NewReader(os.Stdin)
					input, err := inputReader.ReadString('\n')

					if err != nil {
						fmt.Println("Your input is wrong.")
						return
					}

					switch input {
					case "yes\n":{
						dealerCard = append(dealerCard, pop(&cards))
						fmt.Print("This card is:")
						fmt.Println(dealerCard[len(dealerCard) - 1])
						dealerSum += dealerCard[len(dealerCard) - 1]
						fmt.Print("The sum of dealer's card is:")
						fmt.Println(dealerSum)

						if blackjack(dealerCard) {
							fmt.Println("The dealer has 21 points!")
							isDValid = false
							break
						} else if dealerSum > 21 {
							fmt.Println("dealer's cards are busting.")
							DisBust = true
							isDValid = false
							break
						}
					}
					case "hit\n" :{
						isDValid = false
						break
					}

					}

				} else {
					isDValid = false
				}
			}

			if (DisBust && PisBust) || (!DisBust && !PisBust && (dealerSum == playerSum)) {
				fmt.Println("Game is over, it's a push")
			} else if (DisBust && !PisBust) || (!DisBust && !PisBust && (dealerSum < playerSum)) {
				fmt.Println("Game is over, the player wins")
			} else if (!DisBust && PisBust) || (!DisBust && !PisBust && (dealerSum > playerSum)) {
				fmt.Println("Game is over, the dealer wins")
			}

		}
	}
}


func shuffle (cards []int) []int {
	temp := [52]int{}
	l := len(cards)
	for i := l-1; i>0; i-- {
		r := random.Intn(i+1)
		cards[r], cards[i] = cards[i], cards[r]
	}
	temp[cards[0]] += 1
	return cards
}

//deal cards randomly
func pop (cards *[]int) int  {
	pos := rand.Intn(len(*cards)-1)
	card := (*cards)[pos]
	*cards = append((*cards)[1:pos],(*cards)[pos:]...)
	return card
}

//judge if it is blackjack
func blackjack (a []int) bool {
	sum := 0
	hasOne := false
	for i :=0; i<len(a);i++ {
		sum += a[i]
		if(a[i]==1) {
			hasOne = true
		}
	}

	if sum == 21{
		return true;
	} else if hasOne && sum+10==21 {
		return true;
	}
	return false;
}

