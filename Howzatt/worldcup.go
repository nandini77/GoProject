package main

import "fmt"

var IndiaScore int = 0
var SrilankaScore int = 0

func India() {

	IndPlayers := make(map[string]int)
	IndPlayers["Indian player A"] = 20
	IndPlayers["Indian player B"] = 23
	IndPlayers["Indian player C"] = 75
	IndPlayers["Indian player D"] = 5
	IndPlayers["Indian player E"] = 35
	IndPlayers["Indian player F"] = 115
	IndPlayers["Indian player G"] = 0
	IndPlayers["Indian player H"] = 24
	IndPlayers["Indian player I"] = 85
	IndPlayers["Indian player J"] = 95
	IndPlayers["Indian player K"] = 55

	fmt.Println("Names and Scores of Indian Players are : ", IndPlayers)

	SriBowlers := make(map[string]int)
	SriBowlers["Srilanka player C"] = 32
	SriBowlers["Srilanka player D"] = 67
	SriBowlers["Srilanka player G"] = 8
	SriBowlers["Srilanka player I"] = 45
	SriBowlers["Srilanka player K"] = 3

	fmt.Println("Names and runs conceded by srilanka bowlers are : ", SriBowlers)

	for _, value := range IndPlayers {

		IndiaScore = value + IndiaScore

	}
	fmt.Println("Total score of Indian team is : ", IndiaScore, "/ 8")

}

func Srilanka() {
	SriPlayers := make(map[string]int)
	SriPlayers["Srilanka player A"] = 45
	SriPlayers["Srilanka player B"] = 86
	SriPlayers["Srilanka player C"] = 74
	SriPlayers["Srilanka player D"] = 20
	SriPlayers["Srilanka player E"] = 65
	SriPlayers["Srilanka player F"] = 35
	SriPlayers["Srilanka player G"] = 27
	SriPlayers["Srilanka player H"] = 6
	SriPlayers["Srilanka player I"] = 33
	SriPlayers["Srilanka player J"] = 40
	SriPlayers["Srilanka player K"] = 95

	fmt.Println("Names and Scores of Srilanka Players are : ", SriPlayers)

	IndBowlers := make(map[string]int)
	IndBowlers["Indian player B"] = 9
	IndBowlers["Indian player C"] = 19
	IndBowlers["Indian player F"] = 13
	IndBowlers["Indian player H"] = 22
	IndBowlers["Indian player J"] = 6
	fmt.Println("Names and runs conceded by Indian bowlers are : ", IndBowlers)

	for _, value := range SriPlayers {

		SrilankaScore = value + SrilankaScore

	}
	fmt.Println("Total score of Srilanka team is : ", SrilankaScore, "/ 10")

}
func finalResult() {

	var difference int = IndiaScore - SrilankaScore
	fmt.Println("Difference in score between two teams is : ", difference)

	if IndiaScore > SrilankaScore {

		fmt.Println("INDIA WON MATCH BY", difference, "RUNS")
	} else {
		fmt.Println("SRILANKA WON MATCH BY", difference, "RUNS")
	}
}
func main() {

	India()
	Srilanka()
	finalResult()

}
