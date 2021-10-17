package main

import (
	"fmt"
	GCD "github.com/vincehow/learning-go/homework/00"
	"math"
	"math/rand"
	"time"
)

func main() {
	//ComputeHouseEdge(10000)
	//PlayCrapsOnce()
	// create random seed
	//rand.Seed(time.Now().UnixNano())
	//// read in electoral votes data
	//testPolls := map[string]float64{"A":0.55, "B":0.49}
	//testElectoralVotes := map[string]int{"A":10, "B":15}
	//SimulateMultipleElections(testPolls, testElectoralVotes, 1, 0.1)


	// #1 WeightedDie
	//fmt.Println(WeightedDie())

	// #2
	//TrivialGCD(4,8)
	//compareSingleRun(100,50)
	//compareMultipleRuns(1000, 2000, 10)
	//compareMultipleRuns(10000, 20000, 10)
	//compareMultipleRuns(100000, 200000, 10)
	//compareMultipleRuns(1000000, 2000000, 10)

	// #3
	//RelativelyPrimeProbability(1000, 2000, 500)

	// #4
	fmt.Println(HasRepeat([]int{1,2,3,4,6,7,8,8}))
	fmt.Println(HasRepeat([]int{1,2,3,4,6,7,8}))
	// #5

}


func PlayCrapsOnce() bool {
	round := 1
	point := 0
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
	// carry playing
	point = score
	round ++
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


func AddNoise(poll float64, marginOfError float64) float64 {
	x := rand.NormFloat64()
	x = x/2
	x = x*marginOfError
	return x+poll
}

func SimulateOneElection(polls map[string]float64, electoralVotes map[string]int, marginOfError float64) [2]int {
	votes1 := 0
	votes2 := 0
	for k, v := range polls {
		poll := v
		adjustedPoll := AddNoise(poll, marginOfError)
		if adjustedPoll >= 0.5 {
			votes1 += electoralVotes[k]
		} else {
			votes2 += electoralVotes[k]
		}
	}
	result := [2]int{votes1,votes2}
	return result
}


func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]int, numTrials int, marginOfError float64) [3]float64 {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0
	for i := 1; 1<= numTrials; i++ {
		electionVotes := SimulateOneElection(polls, electoralVotes, marginOfError)
		votes1 := electionVotes[0]
		votes2 := electionVotes[1]
		if votes1 > votes2 {
			winCount1 ++
			//fmt.Println("1 win")
		} else if votes2 > votes1 {
			winCount2 ++
		} else {
			tieCount ++
		}
	}
	probability1 := float64(winCount1)/float64(numTrials)
	probability2 := float64(winCount2)/float64(numTrials)
	probabilityTie := float64(tieCount)/float64(numTrials)
	fmt.Printf("Proba of 1 winning = %v; 2 winning = %v; tie = %v",probability1,	probability2, probabilityTie )
	return [3]float64{probability1, probability2, probabilityTie}
}


func WeightedDie() int {
	randomNumber := rand.Float32()
	if randomNumber < 1/10 {
		return 1
	} else if randomNumber < 2/10 {
		return 2
	} else if randomNumber < 7/10 {
		return 3
	} else if randomNumber < 8/10 {
		return 4
	} else if randomNumber < 9/10 {
		return 5
	} else {
		return 6
	}
	return -1
}

func TrivialGCD(x int, y int) int{
	// in the first lesson, I didn't actually write the trivial solution, so I am creating it from scratch
	var commonDividers []int
	for i := 1; i <= int(math.Min(float64(x),float64(y))); i++ {
		if x%i == 0 && y%i ==0{
			commonDividers = append(commonDividers, i)
		}
	}
	result := GCD.MaxArray(commonDividers)
	//fmt.Println(result)
	return result
}


func EuclidGCD(x int, y int) int {
	GCD := 1
	for x != y{
		if x > y {
			x = x-y
		} else {
			y = y-x
		}
	}
	GCD = x
	return GCD
}

func compareSingleRun(x int, y int) [2]float64{
	start1 := time.Now()
	EuclidGCD(x,y)
	elapsed1 := float64(time.Since(start1))
	start2 := time.Now()
	TrivialGCD(x,y)
	elapsed2 := float64(time.Since(start2))
	comp := [2]float64{elapsed1, elapsed2}
	//fmt.Println(comp)
	return comp
}

func avgRunTime(runs [][2]float64) [2]float64 {
	var EuclidGCDRuns []float64
	var TrivialGCDRuns []float64
	for i := range runs{
		EuclidGCDRuns = append(EuclidGCDRuns, runs[i][0])
		TrivialGCDRuns = append(TrivialGCDRuns, runs[i][1])
	}
	EuclidAVG := avgArray(EuclidGCDRuns)
	TrivialAVG := avgArray(TrivialGCDRuns)
	return [2]float64{EuclidAVG, TrivialAVG}
}

func avgArray(x []float64) float64 {
	var sum float64
	len := float64(len(x))
	for i := range x {
		sum += x[i]
	}
	return sum / len
}

func compareMultipleRuns(lower int, upper int, n int) [2]float64 {
	// takes n pairs of random numbers within the two bounds then find their GCD using the two methods. Record the avg speed of the n runs by each method
	fmt.Printf("Simulating %v runs for numbers between %v - %v \n", n, lower, upper)
	var runResults [][2]float64
	for run := 1; run <= n; run ++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(upper - lower) + lower
		rand.Seed(time.Now().UnixNano())
		y := rand.Intn(upper - lower) + lower
		singleRun := compareSingleRun(x,y)
		runResults = append(runResults, singleRun)
	}
	avgRunResults := avgRunTime(runResults)
	fmt.Printf("AVG Euclid run time: %v \nAVG Trivial run time: %v \n", avgRunResults[0], avgRunResults[1])
	return avgRunResults
}


func RelativelyPrimeProbability(lower int, upper int, n int) float64 {
	relativePrimeCount := 0
	for i := 1; i <= n; i++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(upper - lower) + lower
		rand.Seed(time.Now().UnixNano())
		y := rand.Intn(upper - lower) + lower
		if isRelativePrime(x, y) {
			relativePrimeCount ++
		}
	}
	var probRelativePrime float64 = float64(relativePrimeCount) / float64(n)
	fmt.Printf("Proba of a random pair of numbers in the range %v - %v being relatively prime: %v", lower, upper, probRelativePrime)
	return probRelativePrime
}

func isRelativePrime(x int, y int) bool {
	if EuclidGCD(x,y) == 1 {
		return true
	} else {
		return false
	}
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func HasRepeat(x []int) bool {
	var numbers []int
	for i := range x {
		// add one to the count if int already seen before
		if contains(numbers, x[i]) {
			return true
		} else {
			numbers = append(numbers, x[i])
		}
	}
	return false
}
