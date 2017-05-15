package main

import (
	"math/rand"
	"time"
)

func main() {
	initUI()
	rand.Seed(time.Now().UTC().UnixNano())

	var gameState GameState

	//var currentBag []int
	//score := 0
	//var currentPiece Piece
	//var gameBoard GameBoard
	//var gameOver = false
	gameState.currentBag = rand.Perm(7)
	gameState.currentPiece, gameState.currentBag = getNextPieceAsAppliedToBoard(gameState.currentBag)

	commandChannel := make(chan Command, 100)
	stateChannel := make(chan bool)
	replyChannel := make(chan bool)

	go listenForCommands(commandChannel, stateChannel, replyChannel)

	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			//tick through game
			//score, gameOver, currentPiece, gameBoard, currentBag = doTheStuff(TIMEDROP, currentPiece, gameBoard, currentBag, score)
			gameState = gameState.doTheStuff(TIMEDROP)
			drawUI(gameState.gameBoard, gameState.currentPiece, gameState.currentBag, gameState.score)
		case command := <-commandChannel:
			//score, gameOver, currentPiece, gameBoard, currentBag = doTheStuff(command, currentPiece, gameBoard, currentBag, score)
			gameState = gameState.doTheStuff(command)
			drawUI(gameState.gameBoard, gameState.currentPiece, gameState.currentBag, gameState.score)
		}
		if gameState.gameOver {
			drawGameOver()
			setCursorToEnd()
			ticker.Stop()
			stateChannel <- true
			<-replyChannel
			break
		}
	}
}
