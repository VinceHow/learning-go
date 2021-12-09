package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"math/rand"
	"time"
)

const (
	screenWidth  = 1000
	screenHeight = 700
	gridSize     = 20
)

func getSpeed() [8]float64 {
	possibleSpeeds := [8]float64{0, 0.25, 0.5, 1, 2, 3, 4, 5}
	return possibleSpeeds
}

type Cell struct {
	row		int
	col		int
}

type Game struct {
	board  					board
	liveCells 				[]Cell
	generations 			int64
	evolutionSpeed 			int // tiers of speeds taken, with 0 being first tier
	gamePaused				bool
}

func getLiveCells(b board) []Cell {
	// converts a board into a list of alive cell positions
	var liveCells []Cell
	for r, row := range b {
		for c, v := range row {
			if v {
				liveCells = append(liveCells, struct {
					row int
					col int
				}{row: r, col: c})
			}
		}
	}
	return liveCells
}

func (g *Game) Update() error {
	// process user input

	// 1. game speed controls
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.evolutionSpeed > 0 { // cannot reduce speed to below 0
			g.evolutionSpeed --
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.evolutionSpeed < len(getSpeed())-1 { // cannot increase above top speed
			g.evolutionSpeed ++
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) { // reset game
		g.reset()
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) { // pause game
		if g.gamePaused {
			g.gamePaused = false
		} else {
			g.gamePaused = true
		}
	}

	//// update board state
	if len(g.liveCells) >= 1 && !g.gamePaused{
		g.generations ++
		g.board = UpdateBoard(g.board)
		time.Sleep(time.Duration(int64(1000000000/getSpeed()[g.evolutionSpeed]))) // convert speed into wait time in nanoseconds

	}
	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *Game) reset() {
	// clear all live cells
	var b2 board
	b2 = append(b2, []bool{false,false,false,false,false})
	b2 = append(b2, []bool{false,false,false,false,false})
	b2 = append(b2, []bool{false,true,true,true,false})
	b2 = append(b2, []bool{false,false,false,false,false})
	b2 = append(b2, []bool{false,false,false,false,false})
	g.liveCells = getLiveCells(b2)
	g.generations = 0
	g.evolutionSpeed = 3
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draws the world of cells
	g.liveCells = getLiveCells(g.board)
	for _, v := range g.liveCells {
		ebitenutil.DrawRect(screen, float64(v.col*gridSize), float64(v.row*gridSize), gridSize, gridSize, color.RGBA{R: 0x80, G: 0xa0, B: 0xc0, A: 0xff})
	}
	// game information
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Number of cells alive: %v \nNumber of generations: %v \nCurrent speed: %v \n", len(g.liveCells), g.generations, getSpeed()[g.evolutionSpeed]))
}

func NewGame() *Game {
	var b2 board
	b2 = append(b2, []bool{false,false,false,false,false})
	b2 = append(b2, []bool{false,false,false,false,false})
	b2 = append(b2, []bool{false,true,true,true,false})
	b2 = append(b2, []bool{false,false,false,false,false})
	b2 = append(b2, []bool{false,false,false,false,false})
	g := &Game{
		generations:    0,
		evolutionSpeed: 3,
		gamePaused:     false,
		board: 			b2,
		liveCells:     	getLiveCells(b2),
	}
	return g
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
///////////



//////////////
func main() {
	/*
	Start of game:
		1. start with black screen with prompt for user to add live cells
		2. press space once finish adding cells to start game
	Once game is running:
		1. update screen for each generation
		2. up and down arrow to adjust speed of evolution
		3. press space to pause/unpause
	Game is terminated if:
		1. all cells are dead
		OR
		2. user presses Escape
	Game terminated:
		1. show number of generations ran
		2. show current cells alive
		3. prompt to start new round
	 */
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game of Life")
	ebiten.RunGame(NewGame())
}

