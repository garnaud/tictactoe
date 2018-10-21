package main

import "fmt"
import "strings"

// Board type contains game positions and remaining position
type Board struct {
	game      [][]string
	remaining []bool
}

// NewBoard creates an empty board
func NewBoard() Board {
	return Board{
		game: [][]string{
			[]string{"-", "-", "-"},
			[]string{"-", "-", "-"},
			[]string{"-", "-", "-"}},
		remaining: []bool{true, true, true, true, true, true, true, true, true},
	}
}

// copy current board to a new one
func (b Board) copy() Board {
	// copy game
	game := make([][]string, 3, 3)
	for l, line := range b.game {
		game[l] = make([]string, 3)
		for c, item := range line {
			game[l][c] = item
		}
	}
	// copy remaining
	remaining := make([]bool, 9)
	for i, value := range b.remaining {
		remaining[i] = value
	}
	return Board{
		game:      game,
		remaining: remaining,
	}
}

// same check this board is equals to the other one
func (b Board) same(other Board) bool {
	// compare game
	for l, line := range b.game {
		for c, this := range line {
			if this != other.game[l][c] {
				return false
			}
		}
	}
	// compare remaining
	for i, value := range b.remaining {
		if value != other.remaining[i] {
			return false
		}
	}
	return true
}

// add a new position for a given player
func (b *Board) add(player string, pos int) {
	i := 0
	for j := 0; j < 3; j++ {
		for k := 0; k < 3; k++ {
			if i == pos {
				b.game[j][k] = player
				b.remaining[pos] = false
				return
			} else {
				i++
			}
		}
	}
}

// remainingPosition returns all remaining positions
func (b Board) remainingPosition() []int {
	remainingPositions := make([]int, 0)
	for pos, remaining := range b.remaining {
		if remaining {
			remainingPositions = append(remainingPositions, pos)
		}

	}
	return remainingPositions
}

// toString return a pretty string of this board
func (b Board) toString() string {
	var result string
	result = "game:\n%s\n%s\n%s\nRemaining:\n%v\n"
	return fmt.Sprintf(result, strings.Join(b.game[0], ","), strings.Join(b.game[1], ","), strings.Join(b.game[2], ","), b.remainingPosition())
}

// isLeaf returns true if there is a winner or there remains no more position
// and the winner if exists else ""
func (b Board) isLeaf() (bool, string) {
	if len(b.remainingPosition()) == 0 {
		return true, ""
	}
	// line and column (trick: interleaving tests)
	for column, line := range b.game {
		if line[0] != "-" && line[0] == line[1] && line[0] == line[1] && line[1] == line[2] {
			return true, line[0]
		}
		if b.game[0][column] != "-" && b.game[0][column] == b.game[1][column] && b.game[0][column] == b.game[1][column] && b.game[1][column] == b.game[2][column] {
			return true, b.game[0][column]
		}
	}

	// diagonals
	if b.game[0][0] != "-" && b.game[0][0] == b.game[1][1] && b.game[0][0] == b.game[2][2] && b.game[2][2] == b.game[1][1] {
		return true, b.game[0][0]
	}
	if b.game[0][2] != "-" && b.game[0][2] == b.game[1][1] && b.game[0][2] == b.game[2][0] && b.game[2][0] == b.game[1][1] {
		return true, b.game[0][2]
	}

	return false, ""
}
