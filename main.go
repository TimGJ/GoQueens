package main

// Implementation of the n-queens problem in Go.

import (
	"fmt"
	"bytes"
)

const (
	EMPTY = iota
	THREATENED
	TAKEN
)

type Square int

func (s Square) String() string {
	switch s {
	case EMPTY:
		return " "
	case THREATENED:
		return "T"
	case TAKEN:
		return "Q"
	default:
		return "?"
	}
}

type Row []Square
type Board []Row

func NewBoard(size int, old Board) Board {
	board := make(Board, size)
	for row := 0 ; row < size ; row++ {
		board[row] = make(Row, size)
		if old != nil {
			for col := 0 ; col < size ; col++ {
				board[row][col] = old[row][col]
			}
		}
	}
	return board
}

func (b Board) String() string {
	var spacerbuf, boardbuf bytes.Buffer

	for colnum := 0 ; colnum < len(b) ; colnum++ {
		spacerbuf.WriteString("+---")
	}
	spacerbuf.WriteString("+\n")

	for rownum, row := range b {
		boardbuf.Write(spacerbuf.Bytes())
		for colnum := 0 ; colnum < len(row) ; colnum++ {
			boardbuf.WriteString(fmt.Sprintf("| %v ", b[rownum][colnum]))
		}
		boardbuf.WriteString("|\n")
	}
	boardbuf.Write(spacerbuf.Bytes())
	return boardbuf.String()

}

func (b Board) Available(row, col int) bool {
	return b[row][col] == EMPTY
}

func (b Board) Occupy(row, col int) {
	// Mark all vertical, horizontal and diagnoal sqaures
	// as threatened. Then set the appropriate square to taken
	size := len(b)
	for i := 0 ; i < size; i++ {
		if i != row {
			b[i][col] = THREATENED
		}
		if i != col {
			b[row][i] = THREATENED
		}
		if row-i >= 0 && col-i >= 0 {
			b[row-i][col-i] = THREATENED
		}
		if row-i >= 0 && col+i < size {
			b[row-i][col+i] = THREATENED
		}
		if row+i < size && col-i >= 0 {
			b[row+i][col-i] = THREATENED
		}
		if row+i < size && col+i < size {
			b[row+i][col+i] = THREATENED
		}
	}
	b[row][col] = TAKEN
}

func Play(b Board, row int) {
	if row < len(b) {
		for col := 0 ; col < len(b[row]) ; col++ {
			n := NewBoard(len(b), b)
			if b.Available(row, col) {
				n.Occupy(row, col)
				Play(n, row+1)
			}
		}
	} else {
		fmt.Println(b)
	}
}
func main() {

	const size = 8
	Play(NewBoard(size, nil),0)
}
