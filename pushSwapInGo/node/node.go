package node

// Node Used for List Data Structure
type Node struct {
	Number int
	Next   *Node
}

// CreateNode returns pointer to created Node Data Structure
func CreateNode(number int) *Node {
	return &Node{number, nil}
}
