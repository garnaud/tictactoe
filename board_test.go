package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	board := NewBoard()
	board.add("x", 1)
	board.add("o", 5)
	fmt.Println("board:%+v", board)

	if !stringEquals(board.game[0], []string{"-", "x", "-"}) {
		t.Fail()
	}
	if !boolEquals(board.remaining, []bool{true, false, true, true, true, false, true, true, true}) {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	// given
	original := NewBoard()
	original.add("x", 0)
	original.add("x", 8)
	original.add("o", 4)

	// test
	duplicate := original.copy()
	original.add("o", 2)
	duplicate.add("o", 5)

	// check
	// original should have just the 'o' in position 2
	same := original.same(Board{
		game: [][]string{
			[]string{"x", "-", "o"},
			[]string{"-", "o", "-"},
			[]string{"-", "-", "x"}},
		remaining: []bool{false, true, false, true, false, true, true, true, false},
	})
	if !same {
		t.Fail()
	}
	// duplicate should have just the 'o' in position 2
	same = duplicate.same(Board{
		game: [][]string{
			[]string{"x", "-", "-"},
			[]string{"-", "o", "o"},
			[]string{"-", "-", "x"}},
		remaining: []bool{false, true, true, true, false, false, true, true, false},
	})
	if !same {
		t.Fail()
	}
}

func TestLeafEmpty(t *testing.T) {
	board := NewBoard()
	fmt.Print(board.toString())

	isLeaf, winner := board.isLeaf()
	fmt.Printf("isLeaf: %v, winner: %s\n", isLeaf, winner)
	if isLeaf || winner != "" {
		t.Fail()
	}
}
func TestLeafDiagonal1(t *testing.T) {
	board := NewBoard()
	board.add("x", 0)
	board.add("x", 4)
	board.add("x", 8)
	board.add("o", 2)
	board.add("o", 1)
	fmt.Print(board.toString())

	isLeaf, winner := board.isLeaf()
	fmt.Printf("isLeaf: %v, winner: %s\n", isLeaf, winner)
	if !isLeaf || winner != "x" {
		t.Fail()
	}
}

func TestLeafDiagonal2(t *testing.T) {
	board := NewBoard()
	board.add("x", 2)
	board.add("x", 4)
	board.add("x", 6)
	board.add("o", 5)
	board.add("o", 1)
	fmt.Print(board.toString())

	isLeaf, winner := board.isLeaf()
	fmt.Printf("isLeaf: %v, winner: %s\n", isLeaf, winner)
	if !isLeaf || winner != "x" {
		t.Fail()
	}
}

func TestLeafColumn(t *testing.T) {
	board := NewBoard()
	board.add("x", 0)
	board.add("o", 1)
	board.add("x", 3)
	board.add("o", 4)
	board.add("x", 6)
	fmt.Print(board.toString())

	isLeaf, winner := board.isLeaf()
	fmt.Printf("isLeaf: %v, winner: %s\n", isLeaf, winner)
	if !isLeaf || winner != "x" {
		t.Fail()
	}
}

func TestLeafLine(t *testing.T) {
	board := NewBoard()
	board.add("x", 0)
	board.add("o", 3)
	board.add("x", 1)
	board.add("o", 4)
	board.add("x", 2)
	fmt.Print(board.toString())

	isLeaf, winner := board.isLeaf()
	fmt.Printf("isLeaf: %v, winner: %s\n", isLeaf, winner)
	if !isLeaf || winner != "x" {
		t.Fail()
	}
}

func boolEquals(a, b []bool) bool {
	for i, av := range a {
		if av != b[i] {
			fmt.Printf("%s is different of %s (a=%v, b=%v)", av, b[i], a, b)
			return false
		}
	}
	return true
}

func stringEquals(a, b []string) bool {
	for i, av := range a {
		if av != b[i] {
			fmt.Printf("%s is different of %s (a=%v, b=%v)", av, b[i], a, b)
			return false
		}
	}
	return true
}
