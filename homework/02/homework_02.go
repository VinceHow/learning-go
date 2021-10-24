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
	//fmt.Println(HasRepeat([]int{1,2,3,4,6,7,8,8}))
	//fmt.Println(HasRepeat([]int{1,2,3,4,6,7,8}))

	// #5
	//BirthdayParadox(24, 1000)

	// #6
	//ComputePeriodLength([]int{1, 2, 3, 4, 5})

	// #7
	//CountNumDigits(-10)
	//SquareMiddle(4, 2)
	//GenerateMiddleSquareSequence(3792,4)
	//evaluateSeedPeriod(2345,4)
	seedPeriodMap := evaluateSeedPeriodInRange(1, 9999, 4)
	countPeriodLessThanTen := 0
	for _, value := range seedPeriodMap {
		if value <= 10 {
			countPeriodLessThanTen ++
		}
	}
	fmt.Printf("For seeds between 1 and 9999, there are %v number of seeds that have a period of less than or equal to 10. \n", countPeriodLessThanTen)
	fmt.Println("The Middle-Square approach is very bad PRNG approach. Once a seed reaches a repeating loop, the period of repeat is very short.")


}

func PlayCrapsOnce() bool {
	round := 1
	point := 0
	score := SumTwoDice()
	// player wins
	for _, v := range [2]int{7, 11} {
		if score == v {
			//fmt.Println("Player wins during round number", round-1)
			return true
		}
	}
	// house wins
	for _, v := range [3]int{2, 3, 12} {
		if score == v {
			//fmt.Println("House wins during round number", round-1)
			return false
		}
		// carry playing
		point = score
		round++
	}
	for round > 1 {
		round++
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

func ComputeHouseEdge(n int) float64 {
	countWins := 0
	for i := 1; i <= n; i++ {
		outcome := PlayCrapsOnce()
		if !outcome {
			countWins++
		} else {
			countWins--
		}
	}
	edge := float64(countWins) / float64(n)
	fmt.Println("House edge is", edge)
	return edge
}

func AddNoise(poll float64, marginOfError float64) float64 {
	x := rand.NormFloat64()
	x = x / 2
	x = x * marginOfError
	return x + poll
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
	result := [2]int{votes1, votes2}
	return result
}

func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]int, numTrials int, marginOfError float64) [3]float64 {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0
	for i := 1; 1 <= numTrials; i++ {
		electionVotes := SimulateOneElection(polls, electoralVotes, marginOfError)
		votes1 := electionVotes[0]
		votes2 := electionVotes[1]
		if votes1 > votes2 {
			winCount1++
			//fmt.Println("1 win")
		} else if votes2 > votes1 {
			winCount2++
		} else {
			tieCount++
		}
	}
	probability1 := float64(winCount1) / float64(numTrials)
	probability2 := float64(winCount2) / float64(numTrials)
	probabilityTie := float64(tieCount) / float64(numTrials)
	fmt.Printf("Proba of 1 winning = %v; 2 winning = %v; tie = %v", probability1, probability2, probabilityTie)
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

func TrivialGCD(x int, y int) int {
	// in the first lesson, I didn't actually write the trivial solution, so I am creating it from scratch
	var commonDividers []int
	for i := 1; i <= int(math.Min(float64(x), float64(y))); i++ {
		if x%i == 0 && y%i == 0 {
			commonDividers = append(commonDividers, i)
		}
	}
	result := GCD.MaxArray(commonDividers)
	//fmt.Println(result)
	return result
}

func EuclidGCD(x int, y int) int {
	GCD := 1
	for x != y {
		if x > y {
			x = x - y
		} else {
			y = y - x
		}
	}
	GCD = x
	return GCD
}

func compareSingleRun(x int, y int) [2]float64 {
	start1 := time.Now()
	EuclidGCD(x, y)
	elapsed1 := float64(time.Since(start1))
	start2 := time.Now()
	TrivialGCD(x, y)
	elapsed2 := float64(time.Since(start2))
	comp := [2]float64{elapsed1, elapsed2}
	//fmt.Println(comp)
	return comp
}

func avgRunTime(runs [][2]float64) [2]float64 {
	var EuclidGCDRuns []float64
	var TrivialGCDRuns []float64
	for i := range runs {
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
	for run := 1; run <= n; run++ {
		rand.Seed(time.Now().UnixNano())
		x := rand.Intn(upper-lower) + lower
		y := rand.Intn(upper-lower) + lower
		singleRun := compareSingleRun(x, y)
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
		x := rand.Intn(upper-lower) + lower
		y := rand.Intn(upper-lower) + lower
		if isRelativePrime(x, y) {
			relativePrimeCount++
		}
	}
	var probRelativePrime float64 = float64(relativePrimeCount) / float64(n)
	fmt.Printf("Proba of a random pair of numbers in the range %v - %v being relatively prime: %v", lower, upper, probRelativePrime)
	return probRelativePrime
}

func isRelativePrime(x int, y int) bool {
	if EuclidGCD(x, y) == 1 {
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
		// check if number has been seen before
		if contains(numbers, x[i]) {
			return true
		} else {
			numbers = append(numbers, x[i])
		}
	}
	return false
}

func BirthdayParadox(numPeople int, numTrials int) float64 {
	repeats := 0
	for run := 1; run <= numTrials; run++ {
		// generate n random numbers between 1 - 365
		var birthdays []int
		for i := 0; i < numPeople; i++ {
			rand.Seed(time.Now().UnixNano())
			birthday := rand.Intn(365) + 1
			birthdays = append(birthdays, birthday)
		}
		if HasRepeat(birthdays) {
			repeats++
		}
	}
	proba := float64(repeats) / float64(numTrials)
	fmt.Printf("With %v people in the room, the prob of having at least two people sharing the same birthday is %v", numPeople, proba)
	return proba
}

func RealComputePeriodLength(x []int) int {
	/* I'm not sure that I understood the goal correctly, but I'm assuming that it means:
	1. take a list, determine whether it had repeated numbers
	2. check if the repeated numbers form a repeated segment of number
	3. capture the repeated segment, and return its length

	I am basing the above on the definition here: https://www.expii.com/t/periodic-sequences-definition-examples-4348#:~:text=The%20period%20of%20a%20sequence,be%20a%20positive%20whole%20number.
	In order words:
	- This is a periodic sequence: 1,1,2,1,1,2 (period = 3)
	- This is NOT a periodic sequence: 1,2,2,2,2,2
	- This is NOT a periodic sequence: 1,2,3,1,2
	*/
	if HasRepeat(x) {
		/*
			1. take the first number and find the indices of all its repeats, these are the potential periods for us to test
			2. add the potential periods to the second number, to see if it's also a repeat
			3. repeat step 2 for all numbers captured in the period
			4. the full cycle must be observed continuously twice
		*/
		var potentialPeriods []int
		for i := range x {
			// check if number has been seen before
			if x[i] == x[0] && i != 0 {
				potentialPeriods = append(potentialPeriods, i)
			}
		}
		fmt.Println("Potential periods:", potentialPeriods)

		// test and eliminate potential periods
		validatedPeriods := potentialPeriods
		for i := range potentialPeriods {
			periodToTest := potentialPeriods[i]
			for numberInPattern := 0; numberInPattern < periodToTest; numberInPattern++ {
				var locationsToValidate []int
				numberToFind := x[numberInPattern]
				for multiple := 1; (numberInPattern)+(multiple*periodToTest) < len(x)-1; multiple++ {
					locationsToValidate = append(locationsToValidate, (numberInPattern+1)+(multiple*periodToTest))
				}
				fmt.Printf("Testing to see if %v exists at the following locations %v\n", numberToFind, locationsToValidate)
				for location := range locationsToValidate {
					if numberToFind != x[locationsToValidate[location]] {
						validatedPeriods[i] = -1
					}
				}
			}
		}

		// find the smallest validated period length and return
		var final []int
		for i := range validatedPeriods {
			if validatedPeriods[i] >= 1 && (validatedPeriods[i]*2) <= len(x) {
				// must have seen as least 2 cycles
				final = append(final, validatedPeriods[i])
			}
		}
		if len(final) != 0 {
			fmt.Println("All valid periods:", final, "\n")
			return GCD.MinArray(final)
		} else {
			fmt.Println("No valid period found")
		}
	}
	return 0 // 0 if no period was found
}

func ComputePeriodLength(x []int) int {
	// maps unique numbers to the first index that they were seen
	numbersIndex := make(map[int]int)
	for i := range x {
		// check if number has been added to map already
		if _, ok := numbersIndex[x[i]]; ok {
			// if number has been seen, compute period
			firstAppearance := numbersIndex[x[i]]
			period := i - firstAppearance
			//fmt.Println("First period found has length:", period)
			return period
		} else { // add number to map
			numbersIndex[x[i]] = i
		}
	}
	fmt.Println("No repeat found")
	return 0
}

func CountNumDigits(x int) int {
	count := 0
	num := x
	for num != 0 {
		num /= 10
		count += 1
	}
	//fmt.Printf("Digits count of %v is %v \n", x, count)
	return count
}

func Pow10(x int) int {
	return int(math.Pow10(x))
}

func SquareMiddle(seed int, numDigits int) int{
	/*
	https://en.wikipedia.org/wiki/Middle-square_method
	1. square x to get Y
	2. we force Y to be 2N-digit number (where N is the desired number of digits, numDigits)
		- if Y does not have 2N digits, we add leading zeroes to it to compensate (not literally, but effectively)
	2. get the middle N digits of Y
	 */
	N := CountNumDigits(seed)
	if N > numDigits {
		panic("Seed has more digits than desired output.")
	}
	if numDigits % 2 != 0 {
		panic("The desired number of digits in the output integer is not even.")
	}
	var Y = int(math.Pow(float64(seed), 2))
	// cross off leading N/2 numbers
	Y = Y % Pow10(numDigits+(numDigits/2))
	// cross off trailing N/2 numbers
	Y = Y / Pow10(numDigits/2)
	//fmt.Printf("The random number generated by the seed %v is %v \n", seed, Y)
	return Y
}

func GenerateMiddleSquareSequence(seed int, numDigits int) []int {
	var seq []int
	newSeed := seed
	for HasRepeat(seq) == false {
		newSeed = SquareMiddle(newSeed, numDigits)
		seq = append(seq, newSeed)
	}
	//fmt.Println(seq)
	return seq
}

func evaluateSeedPeriod(seed int, numDigits int) int {
	seq := GenerateMiddleSquareSequence(seed, numDigits)
	period := ComputePeriodLength(seq)
	//fmt.Printf("Seed %v has a period of %v", seed, period)
	return period
}

func evaluateSeedPeriodInRange(lower int, upper int, numDigits int) map[int]int {
	seedPeriodMap := make(map[int]int)
	for i := lower; i <= upper; i ++ {
		seedPeriodMap[i] = evaluateSeedPeriod(i, numDigits)
	}
	return seedPeriodMap
}