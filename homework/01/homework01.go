package main

import "fmt"

func main() {
	//PatternCount("CGATATATCCATAG", "ATA")
	//FrequentWords("CGATATATCCATAG", 2)
	//FrequencyTable("CGATATATCCATAG", 3)
	BetterFrequentWords("CGATATATCCATAG", 3)
}

func PatternCount(x1 string, x2 string) int {
	matchCount := 0
	lenSubString := len(x2)
	for i := 0; i <= (len(x1)-lenSubString); i++ {
		subString := x1[i:i+lenSubString]
		if subString == x2 {
			matchCount ++
		}
	}
	//fmt.Println("Number of matches =", matchCount)
	return matchCount
}

func FrequentWords(x string, k int) []string {
	lenText := len(x)
	matchCount := make([]int, lenText-k+1)
	for i:= 0; i <= (lenText-k); i++ {
		pattern := x[i:i+k]
		matchCount[i] = PatternCount(x, pattern)
	}
	var freqPatterns []string
	maxCount := MaxArray(matchCount)
	for i:= 0; i <= (lenText-k); i++ {
		if matchCount[i] == maxCount {
			pattern := x[i:i+k]
			if Contains(freqPatterns, pattern) == false {
				freqPatterns = append(freqPatterns, pattern)
			}
		}
	}
	fmt.Println("Most frequent words are:", freqPatterns)
	return freqPatterns
}

func MaxArray(x []int) int {
	biggestInt := x[0]
	for i := range x {
		if biggestInt < x[i] {
			biggestInt = x[i]
		}
	}
	//fmt.Println("Biggest int is", biggestInt)
	return biggestInt
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func FrequencyTable(x string, k int) map[string]int {
	freqMap := make(map[string]int)
	lenText := len(x)
	for i := 0; i <= lenText-k; i++ {
		pattern := x[i:i+k]
		if _, found := freqMap[pattern]; found {
			freqMap[pattern]++
		} else {
			freqMap[pattern] = 1
		}
	}
	fmt.Println("Frequencies of words mapping:", freqMap)
	return freqMap
}

func MaxMap(x map[string]int) int {
	max := 0
	init := true
	for _, v := range x {
		if init || v > max {
			init = false
			max = v
		}
	}
	return max
}

func BetterFrequentWords(x string, k int) []string {
	var frequentPatterns []string
	freqMap := FrequencyTable(x, k)
	max := MaxMap(freqMap)
	for k, v := range freqMap {
		if v == max {
			frequentPatterns = append(frequentPatterns, k)
		}
	}
	fmt.Println(frequentPatterns)
	return frequentPatterns
}






