package main

import (
	"math/rand"
	"time"
)

func main() {
	initUI()
	rand.Seed(time.Now().UTC().UnixNano())

	var currentBag []int
	score := 0
	var currentPiece Piece
	var gameBoard GameBoard
	var gameOver = false
	currentBag = rand.Perm(7)
	currentPiece, currentBag = getNextPieceAsAppliedToBoard(currentBag)

	commandChannel := make(chan Command, 100)
	stateChannel := make(chan bool)
	replyChannel := make(chan bool)

	go listenForCommands(commandChannel, stateChannel, replyChannel)

	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			//tick through game
			score, gameOver, currentPiece, gameBoard, currentBag = doTheStuff(TIMEDROP, currentPiece, gameBoard, currentBag, score)
			//print(score)
			drawUI(gameBoard, currentPiece, currentBag, score)
		case command := <-commandChannel:
			score, gameOver, currentPiece, gameBoard, currentBag = doTheStuff(command, currentPiece, gameBoard, currentBag, score)
			drawUI(gameBoard, currentPiece, currentBag, score)
		}
		if gameOver {
			drawGameOver()
			setCursorToEnd()
			ticker.Stop()
			stateChannel <- true
			<-replyChannel
			break
		}
	}
}
