package board
import "fmt"

type Board struct {

grid [3][3]string

}

 

func NewBoard() *Board {

return &Board{}

}

 

func (b *Board) IsCellEmpty(row, col int) bool {

return b.grid[row][col] == ""

}

 
func (b *Board) SetCell(row, col int, mark string) {

    if row < 0 || row >= 3 || col < 0 || col >= 3 {

        panic("Invalid cell coordinates")

    }

 

    if b.grid[row][col] != "" {

        panic("Cell already occupied")

    }

 

    b.grid[row][col] = mark

}

 

func (b *Board) IsBoardFull() bool {

for i := 0; i < 3; i++ {

for j := 0; j < 3; j++ {

if b.grid[i][j] == "" {

return false

}

}

}

return true

}

 
func (b *Board) IsGameWon() bool {

	
	for i := 0; i < 3; i++ {
	
	if b.grid[i][0] != "" && b.grid[i][0] == b.grid[i][1] && b.grid[i][0] == b.grid[i][2] {
	
	return true
	
	}
	
	}
	
	for i := 0; i < 3; i++ {
	
	if b.grid[0][i] != "" && b.grid[0][i] == b.grid[1][i] && b.grid[0][i] == b.grid[2][i] {
	
	return true
	
	}
	
	}
	
	
	if b.grid[0][0] != "" && b.grid[0][0] == b.grid[1][1] && b.grid[0][0] == b.grid[2][2] {
	
	return true 
	
	}
	
	if b.grid[0][2] != "" && b.grid[0][2] == b.grid[1][1] && b.grid[0][2] == b.grid[2][0] {
	
	return true 
	
	}
	
	 
	
	return false
	
	}

	func (b *Board) Display() {

		fmt.Println("   0 1 2")
		
		for i := 0; i < 3; i++ {
		
		fmt.Print(i, " ")
		
		for j := 0; j < 3; j++ {
		
		if j > 0 {
		
		fmt.Print("|")
		
		}
		
		fmt.Print(" ", b.grid[i][j], " ")
		
		}
		
		fmt.Println()
		
		if i < 2 {
		
		fmt.Println("  -------")
		
		}
		
		}
		
		}

		