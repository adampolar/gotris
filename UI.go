package main

import (
	"strconv"

	tm "github.com/buger/goterm"
)

const TETRIMINO_CHAR = "□"

func printAt(x int, y int, toPrint string, color int) {

	tm.MoveCursor(x, y)
	tm.Print(tm.Color(toPrint, color))
	tm.Flush()
	//fmt.Fprintf(new(bytes.Buffer), "\033[%d;%dH%s", y, x, toPrint)
}
func initUI() {
	print("\033[H\033[2J")
	print(tetrisBoard)
}

func drawGameOver() {
	printAt(10, 10, "┌─────────┐", tm.RED)
	printAt(10, 11, "│GAME OVER│", tm.RED)
	printAt(10, 12, "└─────────┘", tm.RED)
}

func setCursorToEnd() {
	tm.MoveCursor(30, 0)
	tm.Flush()
}

func drawUI(board GameBoard, piece Piece, bag []int, score int) {
	for i, v := range board {
		if i < 2 {
			continue
		}
		for j, u := range v {
			if u != 0 {
				printAt(16+j, 2+i, TETRIMINO_CHAR, u)
			} else {
				printAt(16+j, 2+i, " ", u)
			}
		}
	}

	for k := 0; k < 4; k++ {
		if piece.Coords[k][1] > 1 {
			printAt(16+piece.Coords[k][0], 2+piece.Coords[k][1], TETRIMINO_CHAR, piece.Color)
		}
	}
	for j := 0; j < 8; j++ {
		printAt(7, j+16, "    ", 0)
	}

	for l := 0; l < 3; l++ {
		nextPiece := PIECES[bag[l]]
		for j := 0; j < 4; j++ {
			printAt(7+nextPiece.Coords[j][0], 22+nextPiece.Coords[j][1]-3*l, TETRIMINO_CHAR, nextPiece.Color)
		}
	}

	printAt(7, 6, strconv.Itoa(score), tm.WHITE)

}

var tetrisBoard = `┌────────────────────────────┐
│ GOTRIS                     │
│                            │
│  SCORE:     │          │   │
│             │          │   │
│             │          │   │
│             │          │   │
│             │          │   │
│             │          │   │
│             │          │   │
│             │          │   │
│             │          │   │
│             │          │   │
│             │          │   │
│   ┌─────┐   │          │   │
│   │     │   │          │   │
│   │     │   │          │   │
│   │     │   │          │   │
│   │     │   │          │   │
│   │     │   │          │   │
│   │     │   │          │   │
│   │     │   │          │   │
│   │     │   │          │   │
│   └─────┘   └──────────┘   │
│                            │
└────────────────────────────┘`
