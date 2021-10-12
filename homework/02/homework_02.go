package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ComputeHouseEdge(10000)
	//PlayCrapsOnce()
}


func PlayCrapsOnce() bool {
	round := 1
	point := 0
	if round == 1 {
		round ++
		score := SumTwoDice()
		// player wins
		for _, v := range [2]int{7,11} {
			if score == v {
				//fmt.Println("Player wins during round number", round-1)
				return true
			}
		}
		// house wins
		for _, v := range [3]int{2,3,12} {
			if score == v {
				//fmt.Println("House wins during round number", round-1)
				return false
			}
		}
		// carry playing
		point = score
	}
	for round > 1 {
		round ++
		score := SumTwoDice()
		if score == point { // player wins
			//fmt.Println("Player wins during round number", round-1)
			return true
		} else if score == 7 { // house wins
			//fmt.Println("House wins during round number", round-1)
			return false
		}
	}
	return false
}

func SumTwoDice() int {
	roll1 := RollDie()
	roll2 := RollDie()
	return roll1 + roll2
}

func RollDie() int {
	roll := rand.Intn(6)
	return roll + 1
}

func RollWeightedDie() int {
	randomNumber := rand.Float32()
	if randomNumber < 1/10 {
		return 1
	} else if randomNumber < 4/10 {
		return 2
	} else if randomNumber < 5/10 {
		return 3
	} else if randomNumber < 6/10 {
		return 4
	} else if randomNumber < 7/10 {
		return 5
	} else {
		return 6
	}
	return -1
}

func ComputeHouseEdge(n int) float64{
	countWins := 0
	for i := 1; i <= n; i++{
		outcome := PlayCrapsOnce()
		if !outcome {
			countWins ++
		} else {
			countWins --
		}
	}
	edge := float64(countWins)/float64(n)
	fmt.Println("House edge is", edge)
	return edge
}