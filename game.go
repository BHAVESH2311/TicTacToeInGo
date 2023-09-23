package game
import (

"fmt"

"tictactoe/board"

"tictactoe/player"

)


type TicTacToe struct {

board  *board.Board

players []player.Player

currentPlayer int

}

 

func NewGame(player1, player2 player.Player) *TicTacToe {

return &TicTacToe{

board: board.NewBoard(),

players: []player.Player{player1, player2},

currentPlayer: 0,

}

}

 

func (g *TicTacToe) Play(row, col int) error {

if g.board.IsCellEmpty(row, col) {

mark := g.players[g.currentPlayer].GetMark()

g.board.SetCell(row, col, mark)

g.currentPlayer = 1 - g.currentPlayer

return nil

}

return fmt.Errorf("Cell already occupied")

}

 

func (g *TicTacToe) IsGameOver() bool {

return g.board.IsGameWon() || g.board.IsBoardFull()

}

 

func (g *TicTacToe) GetWinner() player.Player {

if g.board.IsGameWon() {

return g.players[1 - g.currentPlayer]

}

return nil

}

 

func (g *TicTacToe) GetBoard() *board.Board {

return g.board

}

func (g *TicTacToe) CurrentPlayer() player.Player {

	return g.players[g.currentPlayer]
	
}