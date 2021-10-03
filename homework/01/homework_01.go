package main

import (
	"fmt"
	"strings"
)

func main() {
	//PatternCount("CGATATATCCATAG", "ATA")
	//FrequentWords("CGATATATCCATAG", 2)
	//FrequencyTable("CGATATATCCATAG", 3)
	//BetterFrequentWords("CGATATATCCATAG", 3)
	//Reverse("abc")
	//Complement("CGATATATCCATAG")
	//ReverseComplement("ATGATCAAG")
	//StartingIndices("CGATATATCCATAG", "ATA")
	FindClumps("CGATATATCCATAG", 3,8,2)
}

func PatternCount(x1 string, x2 string) int {
	positions := StartingIndices(x1,x2)
	return len(positions)
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

func Reverse(x string) string {
	lenText := len(x)
	var newText string
	for i := lenText-1; i >= 0; i--{
		newText = newText + string(x[i])
	}
	return newText
}

func Complement(x string) string {
	var newText string
	complementMap := map[string]string{
		"a":"t",
		"t":"a",
		"c":"g",
		"g":"c",
	}
	for _,v := range x {
		newText = newText + complementMap[strings.ToLower(string(v))]
	}
	return newText
}

func ReverseComplement(x string) string{
	newText := Reverse(Complement(x))
	fmt.Println(newText)
	return newText
}

func StartingIndices(x string, k string) []int{
	var positionSlice []int
	lenText := len(x)
	for i := 0; i<= lenText-len(k); i++ {
		if x[i:i+len(k)] == k {
			positionSlice = append(positionSlice, i)
		}
	}
	fmt.Println(positionSlice)
	return positionSlice
}

func FindClumps(x string, k int, L int, t int) []string {
	/*
	x = original genome
	k = length of pattern to find
	L = sliding window of capture
	t = min number of occurrence of pattern string within the window
	 */
	var patternsFound []string
	n := len(x)
	for i := 0; i <= n-L; i++ {
		window := x[i:i+L]
		freqMap := FrequencyTable(window, k)
		for k, v := range freqMap {
			if v >= t && Contains(patternsFound, k) == false {
				patternsFound = append(patternsFound, k)
			}
		}
	}
	fmt.Println(patternsFound)
	return patternsFound
}




