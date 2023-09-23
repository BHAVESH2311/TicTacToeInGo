package main

 

import (

"fmt"

"tictactoe/game"

"tictactoe/player"

)

 

func main() {

    defer func() {

        if r := recover(); r != nil {

            fmt.Println("An error occurred:", r)

        }

    }()

 

    player1 := player.NewHumanPlayer("X","Bhavesh")

    player2 := player.NewHumanPlayer("O","Swastik")

 

    ticTacToe := game.NewGame(player1, player2)

 

    for !ticTacToe.IsGameOver() {

        fmt.Println("Current Board:")

        ticTacToe.GetBoard().Display()

        fmt.Printf("Player %s's turn. Enter row (0-2) and column (0-2): ", ticTacToe.CurrentPlayer().GetName())

        var row, col int

        fmt.Scan(&row, &col)

 

        err := ticTacToe.Play(row, col)

        if err != nil {

            panic(err)

        }

    }

 

    fmt.Println("Final Board:")

    ticTacToe.GetBoard().Display()

 

    winner := ticTacToe.GetWinner()

    if winner != nil {

        fmt.Printf("Player %s wins!\n", winner.GetName())

    } else {

        fmt.Println("It's a draw!")

    }

}