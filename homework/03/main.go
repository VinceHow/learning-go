package main

import (
	"fmt"
)

func main() {
	var b Board
	b = append(b, []bool{true, false})
	b = append(b, []bool{false, true})
	fmt.Printf("Gen 1 \n")
	PrintBoard(b)
	fmt.Printf("Gen 2 \n")
	PrintBoard(UpdateBoard(b))

	var b2 Board
	b2 = append(b2, []bool{false, false, false, false, false})
	b2 = append(b2, []bool{false, false, false, false, false})
	b2 = append(b2, []bool{false, true, true, true, false})
	b2 = append(b2, []bool{false, false, false, false, false})
	b2 = append(b2, []bool{false, false, false, false, false})
	fmt.Printf("Gen 1 of a basic oscillator \n")
	PrintBoard(b2)
	fmt.Printf("Gen 2 of a basic oscillator \n")
	PrintBoard(UpdateBoard(b2))

	PrintBoards(PlayGoL(b2, 4))
}

// Board
// We can think of a Board as an array of arrays (2-dimensional array), with each bool representing the cell's state.
// The Board has r rows and c columns
type Board [][]bool

func PlayGoL(initialBoard Board, numGens int) []Board {
	var boards []Board
	boards = append(boards, initialBoard)
	for i := 1; i <= numGens; i++ {
		b := UpdateBoard(boards[i-1])
		boards = append(boards, b)
	}
	return boards
}

func UpdateBoard(currentBoard Board) Board {
	numRows := CountRows(currentBoard)
	numCols := CountCols(currentBoard)
	newBoard := InitializeBoard(numRows, numCols)
	for rowIndex, row := range currentBoard { // loop through rows
		for cellIndex := range row { // loop through cells
			newBoard[rowIndex][cellIndex] = UpdateCell(currentBoard, rowIndex, cellIndex)
		}
	}
	return newBoard
}

func InitializeBoard(r int, c int) [][]bool {
	var b Board
	for i := 1; i <= r; i++ {
		var row []bool
		for i := 1; i <= c; i++ {
			row = append(row, false)
		}
		b = append(b, row)
	}
	return b
}

func CountRows(x Board) int {
	return len(x)
}

func CountCols(x Board) int {
	var cellsInRow []int
	for _, row := range x {
		cellsInRow = append(cellsInRow, len(row))
	}
	columns := cellsInRow[0]
	// for safety, we conduct a check to confirm every row has the same number of columns
	for _, i := range cellsInRow {
		if i != columns {
			panic("The board contains rows with uneven number of columns.")
		}
	}
	return columns
}

func UpdateCell(currentBoard Board, row int, col int) bool {
	numNeighbors := CountLiveNeighbors(currentBoard, row, col)
	// apply rules when current cell is alive
	if currentBoard[row][col] {
		if numNeighbors == 2 || numNeighbors == 3 { // Rule of propagation
			return true
		} else { // lack of mates / overpopulation the cell dies
			return false
		}
	} else {                   // the cell is currently dead
		if numNeighbors == 3 { // birth to new life
			return true
		} else { // remain dead
			return false
		}
	}
}

func CountLiveNeighbors(currentBoard Board, row int, col int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ { // we loop through every eligible row
		for c := col - 1; c <= col+1; c++ { // and eligible column
			if !(r == row && c == col) && InField(currentBoard, r, c) { // excluding the current cell, and the neighbor is on the board
				if currentBoard[r][c] { // neighbor is alive
					count++
				}
			}
		}
	}
	return count
}

func InField(currentBoard Board, row int, col int) bool {
	numRows := CountRows(currentBoard)
	numCols := CountCols(currentBoard)
	if row < 0 || row > (numRows-1) || col < 0 || col > (numCols-1) {
		return false
	} else {
		return true
	}
}

func PrintBoards(boards []Board) {
	for i := range boards {
		PrintBoard(boards[i])
	}
}

func PrintBoard(b Board) {
	for i := range b {
		PrintRow(b[i])
	}
	fmt.Printf("\n")
}

func PrintRow(row []bool) {
	for i := range row {
		PrintCell(row[i])
	}
	fmt.Printf("\n")
}

func PrintCell(v bool) {
	if v {
		fmt.Printf("⬜ ")
	} else {
		fmt.Printf("⬛ ")
	}
}
