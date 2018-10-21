package main

import "fmt"
import "math/rand"

type Node struct {
	parent      *Node
	children    *[]Node
	wins        int
	simulations int
	board       Board
	player      string
	turn        string
	isLeaf      bool
	position    int
}

func (node *Node) createChild(board Board) {
	var newPlayer string
	if node.player == "x" {
		newPlayer = "o"
	} else {
		newPlayer = "x"
	}
	children := make([]Node, 0)
	child := Node{
		parent:      node,
		children:    &children,
		wins:        0,
		simulations: node.simulations,
		board:       board,
		player:      node.player,
		turn:        newPlayer,
	}
	*node.children = append(*node.children, child)
}

func NewRoot(player string) Node {
	children := make([]Node, 0)
	return Node{
		children:    &children,
		wins:        0,
		simulations: 0,
		board:       NewBoard(),
		player:      player,
		turn:        player,
		isLeaf:      false,
	}
}

func (node *Node) opponent() string {
	if node.turn == "x" {
		return "o"
	}
	return "x"
}

func (node *Node) simulate() (leaf, wins bool) {
	node.simulations = node.simulations + 1
	isLeaf, winner := node.board.isLeaf()
	if isLeaf {
		fmt.Println(node.board.toString())
		if winner == node.player {
			node.wins = node.wins + 1
			return true, true
		}
		return true, false
	}
	return false, false
}

func (node *Node) expansion() {
	for _, pos := range node.board.remainingPosition() {
		child := NewRoot(node.player)
		child.board = node.board.copy()
		child.turn = node.opponent()
		child.board.add(child.turn, pos)
		child.position = pos
		children := append(*node.children, child)
		node.children = &children
	}
}

func (node *Node) selection() bool {
	fmt.Printf("wins=%d,simulation=%d,position=%d\n", node.wins, node.simulations, node.position)
	// simulate and if node is a leaf then backpropagate to parent
	if leaf, wins := node.simulate(); leaf {
		return wins
	}
	// if has not child yet expands
	if len(*node.children) == 0 {
		node.expansion()
	}

	// choose a random child and continue selection
	childPosition := rand.Intn(len(*node.children))
	child := (*node.children)[childPosition]
	wins := child.selection()
	if wins {
		node.wins = node.wins + 1
	}
	// backpropagate result selection
	return wins
}

// returns the best choice
func (node Node) choice() int {
	best := 0
	bestPosition := 0
	fmt.Println("choices:")
	for choice, child := range *node.children {
		fmt.Printf("choice %d: wins=%d,simulation=%d,position=%d\n", choice, child.wins, child.simulations, child.position)
		if child.wins > best {
			bestPosition = child.position
			best = child.wins
		}
	}
	return bestPosition
}

func (node Node) string() {
	for _, child := range *node.children {
		fmt.Printf("simulations=%d,wins=%d,pos=%d,children=%d // ", child.simulations, child.wins, child.position, len(*child.children))
		child.string()
	}
	fmt.Println("end loop")
}

func (node Node) childString() {

}

func main() {
	rand.Seed(134141)
	root := NewRoot("x")
	root.turn = "o"
	for i := 0; i < 10; i++ {
		root.selection()
	}

	fmt.Printf("simulations: \n%+v\n", root.simulations)
	fmt.Printf("wins: \n%+v\n", root.wins)
	fmt.Printf("positions: %v", root.choice())
	root.string()
}
