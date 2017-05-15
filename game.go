package main

import (
	"math/rand"

	tm "github.com/buger/goterm"
)

type GameBoard [22][10]int
type Piece struct {
	Coords PieceCoords
	Color  int
}
type PieceCoords [4][2]int

func doTheStuff(command Command,
	currentPiece Piece,
	gameBoard GameBoard,
	currentBag []int,
	currentScore int) (
	score int,
	gameOver bool,
	newPiece Piece,
	newGameBoard GameBoard,
	newBag []int) {
	var proposedPieceLocation Piece
	if command != TIMEDROP && command != DROP {

		if command == SOFTDROP {
			proposedPieceLocation = translatePiece(0, 1, currentPiece)
		}
		if command == LEFT {
			proposedPieceLocation = translatePiece(-1, 0, currentPiece)
		}
		if command == RIGHT {
			proposedPieceLocation = translatePiece(1, 0, currentPiece)
		}
		if command == CLOCKWISE || command == ANTICLOCKWISE {
			for i := 0; i < 4; i++ {
				proposedPieceLocation = rotatePiece(i, command == CLOCKWISE, currentPiece)
				if !checkForCollisions(proposedPieceLocation, gameBoard) {
					break
				}
			}
		}
		if !checkForCollisions(proposedPieceLocation, gameBoard) && !getHasLanded(gameBoard, proposedPieceLocation) {
			currentPiece = proposedPieceLocation
		}
	} else {
		proposedPieceLocation = translatePiece(0, 1, currentPiece)

		landed := getHasLanded(gameBoard, proposedPieceLocation)

		if command == DROP {
			for !landed {
				currentPiece = proposedPieceLocation
				proposedPieceLocation = translatePiece(0, 1, currentPiece)
				landed = getHasLanded(gameBoard, proposedPieceLocation)
			}
		}

		if landed {
			for i := 0; i < 4; i++ {
				if currentPiece.Coords[i][1] >= 0 {
					gameBoard[currentPiece.Coords[i][1]][currentPiece.Coords[i][0]] = currentPiece.Color
				}
			}
			//check if line has been made
			numberOfLinesMade := 0
			for i, v := range gameBoard {
				lineBeenMade := true
				for _, u := range v {
					lineBeenMade = lineBeenMade && u != 0
				}
				if lineBeenMade {
					numberOfLinesMade++
					for j := i; j > 0; j-- {
						for k := range gameBoard[j] {
							gameBoard[j][k] = gameBoard[j-1][k]
						}
					}
					for l := range gameBoard[0] {
						gameBoard[0][l] = 0
					}
				}
			}
			score = currentScore + [5]int{0, 100, 300, 550, 800}[numberOfLinesMade]
			if len(currentBag) == 3 {
				currentBag = append(currentBag, rand.Perm(7)...)
			}
			//check game is over
			for i := 0; i < 4; i++ {
				if currentPiece.Coords[i][1] < 2 {
					gameOver = true
				}
			}
			currentPiece, currentBag = getNextPieceAsAppliedToBoard(currentBag)
		} else {
			currentPiece = proposedPieceLocation
		}
	}
	if score == 0 {
		score = currentScore
	}
	newGameBoard = gameBoard
	newPiece = currentPiece
	newBag = currentBag
	return
}

func getNextPieceAsAppliedToBoard(currentBag []int) (piece Piece, newCurrentBag []int) {
	piece = PIECES[currentBag[0]]
	newCurrentBag = currentBag[1:]
	numberOfRotations := rand.Intn(4)
	i := 0
	for i < numberOfRotations {
		piece = rotatePiece(0, true, piece)
		i++
	}

	piece = translatePiece(3, 0, piece)
	return
}

func rotatePiece(around int, clockwise bool, currentPiece Piece) Piece {
	baseX := currentPiece.Coords[around][0]
	baseY := currentPiece.Coords[around][1]

	for i := 0; i < 4; i++ {
		deltaX := currentPiece.Coords[i][0] - baseX
		deltaY := currentPiece.Coords[i][1] - baseY
		if clockwise {
			currentPiece.Coords[i][0] = deltaY + baseX
			currentPiece.Coords[i][1] = baseY - deltaX
		} else {
			currentPiece.Coords[i][0] = baseX - deltaY
			currentPiece.Coords[i][1] = baseY + deltaX
		}
	}

	return currentPiece
}

func checkForCollisions(piece Piece, gameBoard GameBoard) bool {
	//checkforcollisions
	for i := 0; i < 4; i++ {
		if piece.Coords[i][1] < 0 {
			continue
		}
		if piece.Coords[i][0] < 0 || piece.Coords[i][0] > 9 {
			return true
		}
		if piece.Coords[i][1] < 0 || piece.Coords[i][1] > 21 {
			return true
		}
		if gameBoard[piece.Coords[i][1]][piece.Coords[i][0]] != 0 {
			return true
		}
	}
	return false
}

func translatePiece(x int, y int, piece Piece) Piece {
	for i := 0; i < 4; i++ {
		piece.Coords[i][0] += x
		piece.Coords[i][1] += y
	}
	return piece
}

func getHasLanded(gameBoard GameBoard, currentPiece Piece) bool {
	hasLanded := false

	for i := 0; i < 4; i++ {
		if currentPiece.Coords[i][1] < 0 {
			continue
		}
		if currentPiece.Coords[i][1] > 21 {
			return true
		}
		hasLanded = hasLanded || gameBoard[currentPiece.Coords[i][1]][currentPiece.Coords[i][0]] != 0
	}
	return hasLanded
}

var PIECES = [7]Piece{
	Piece{Coords: PieceCoords{{1, 0}, {2, 0}, {0, 0}, {3, 0}}, Color: tm.CYAN},
	Piece{Coords: PieceCoords{{0, 0}, {1, 0}, {2, 0}, {0, 1}}, Color: tm.BLUE},
	Piece{Coords: PieceCoords{{2, 0}, {1, 0}, {0, 0}, {2, 1}}, Color: tm.WHITE},
	Piece{Coords: PieceCoords{{1, 0}, {1, 1}, {0, 0}, {2, 1}}, Color: tm.GREEN},
	Piece{Coords: PieceCoords{{1, 0}, {1, 1}, {0, 1}, {2, 0}}, Color: tm.RED},
	Piece{Coords: PieceCoords{{0, 1}, {1, 1}, {1, 0}, {0, 0}}, Color: tm.YELLOW},
	Piece{Coords: PieceCoords{{1, 0}, {1, 1}, {0, 0}, {2, 0}}, Color: tm.MAGENTA},
}
