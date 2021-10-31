package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	size   int32 = 30
	rows   int32 = 20
	cols   int32 = 10
	width  int32 = cols * size
	height int32 = rows * size
)

var (
	fps     int32 = 60
	counter int32
	board   [rows][cols]int
	score   int
)

func drawBoard() {
	// draw grid

	for i := int32(0); i < rows; i++ {
		rl.DrawLine(0, i*size, width, i*size, rl.Black)
	}

	for j := int32(0); j < cols; j++ {
		rl.DrawLine(j*size, 0, j*size, height, rl.Black)
	}

	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			switch board[i][j] {
			case 0:
				break
			case 1:
				rl.DrawRectangle(j*size, i*size, size, size, rl.Gray)
			case 2:
				rl.DrawRectangle(j*size, i*size, size, size, rl.Red)
			default:
				panic("illegal block")
			}
		}
	}
}

func checkCollisionsDown() (bool, int32, int32) {
	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			if board[i][j] == 2 {
				if i == rows-1 || board[i+1][j] == 1 {
					return true, i, j
				}
			}
		}
	}
	return false, -1, -1
}

func checkCollisionsLeft() (bool, int32, int32) {
	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			if board[i][j] == 2 {
				if j == 0 || board[i][j-1] == 1 {
					return true, i, j
				}
			}
		}
	}
	return false, -1, -1
}

func checkCollisionsRight() (bool, int32, int32) {
	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			if board[i][j] == 2 {
				if j == cols-1 || board[i][j+1] == 1 {
					return true, i, j
				}
			}
		}
	}
	return false, -1, -1
}

func moveDown() {
	log.Println("move down")
	newBoard := board
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if board[i][j] == 2 {
				newBoard[i][j] = 0
				newBoard[i+1][j] = 2
			}
		}
	}
	board = newBoard
}

func moveLeft() {
	log.Println("move left")
	newBoard := board
	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			if board[i][j] == 2 {
				newBoard[i][j] = 0
				newBoard[i][j-1] = 2
			}
		}
	}
	board = newBoard
}

func moveRight() {
	log.Println("move right")
	newBoard := board
	for i := int32(0); i < rows; i++ {
		for j := cols - 1; j >= 0; j-- {
			if board[i][j] == 2 {
				newBoard[i][j] = 0
				newBoard[i][j+1] = 2
			}
		}
	}
	board = newBoard
}

func freeze() {
	log.Println("freeze")
	for i := int32(0); i < rows; i++ {
		for j := int32(0); j < cols; j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}

func updateBoard() {
	if yes, i, j := checkCollisionsDown(); yes {
		log.Println("collision", i, j)
		freeze()
		// place new block
		board[0][3] = 2
		board[1][3] = 2
		board[1][4] = 2
		board[1][5] = 2
	}
	if counter++; counter == fps/2 {
		moveDown()
		counter = 0
	}
}

func checkKeys() {
	if rl.IsKeyPressed(rl.KeyRight) {
		if yes, i, j := checkCollisionsRight(); yes {
			log.Println("cannot move right to", i, j)
		} else {
			moveRight()
		}

	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		if yes, i, j := checkCollisionsLeft(); yes {
			log.Println("cannot move left to", i, j)
		} else {
			moveLeft()
		}
	}
}

func main() {
	rl.InitWindow(width, height, "tetris")
	defer rl.CloseWindow()
	rl.SetTargetFPS(fps)

	board[0][3] = 2
	board[1][3] = 2
	board[1][4] = 2
	board[1][5] = 2

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawFPS(10, 10)

		checkKeys()
		drawBoard()
		updateBoard()

		//rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}

/*
const (
	BlockSize = 30
	Width     = 10 * BlockSize
	Height    = 20 * BlockSize
)

type Game struct {
	Score     int
	DropTime  int
	DropTimer int
	Board     [10][20]int // 10 columns, 20 rows
}

func (g Game) DrawBoard() {
	for i := range g.Board {
		rl.DrawLine(int32(i*BlockSize), 0, int32(i*BlockSize), int32(Height), rl.Black)
	}

	for j := range g.Board[0] {
		rl.DrawLine(0, int32(j*BlockSize), int32(Width), int32(j*BlockSize), rl.Black)
	}
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[0]); j++ {
			switch g.Board[i][j] {
			case 0:
				break
			case 1:
				rl.DrawRectangle(int32(i), int32(j), int32(BlockSize), int32(BlockSize), rl.Gray)
			case 2:
				rl.DrawRectangle(int32(i), int32(j), int32(BlockSize), int32(BlockSize), rl.Red)
			default:
				panic("illegal block")
			}
		}
	}
}

func (g *Game) UpdatePiece() {
	checkCollision := func() bool {
		for i := 0; i < len(g.Board); i++ {
			for j := 0; j < len(g.Board[0]); j++ {
				// check  if active piece
				if g.Board[i][j] == 2 {
					if j == Height {
						log.Println("collision", i, j)

						return true
					}
					if g.Board[i][j+1] == 1 {
						log.Println("collision", i, j)
						return true
					}

				}
			}
		}
		return false
	}

	if checkCollision() {
		// if collision, change all
		for i := 0; i < len(g.Board); i++ {
			for j := 0; j < len(g.Board[0]); j++ {
				if g.Board[i][j] == 2 {
					g.Board[i][j] = 1
				}
			}
		}
	} else {
		for i := 0; i < len(g.Board); i++ {
			for j := 0; j < len(g.Board[0])-1; j++ {
				if g.Board[i][j] == 2 {
					g.Board[i][j] = 0
					g.Board[i][j+1] = 2
				}
			}
		}
	}
}

func (g *Game) AddPiece(id int) {
	switch id {
	case 0:

	}
}

func (g *Game) Update() {
	g.DrawBoard()

	g.DropTimer++

	if g.DropTimer == g.DropTime {
		log.Println("drop")
		g.UpdatePiece()
		g.DropTimer = 0
	}
}

func main() {
	rl.InitWindow(Width, Height, "tetris")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var game Game
	game.DropTime = 60 // drop once per second
	game.Board[0][0] = 2

	fmt.Println(game.Board)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		game.Update()

		rl.DrawFPS(10, 10)

		//rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
*/
