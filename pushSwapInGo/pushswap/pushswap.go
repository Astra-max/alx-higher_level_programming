/*
* @Author: Tlekbai Ali
* @Date:   2020-01-04 12:44:51
* @Last Modified by:   Tlekbai Ali
* @Last Modified time: 2020-01-04 16:28:11
 */

package pushswap

import (
	"fmt"

	node "../node"
	stack "../stack"
)

// Node alias
type Node = node.Node

// Stack alias
type Stack = stack.Stack

// PushSwap Data Structure for subject
type PushSwap struct {
	A, B   *Stack
	Array  []int
	Length int
	Ops    int
	Print  int
}

// Init of PushSwap Data Structure
// Main Job: to push back numbers from PushSwap.Array to Stack PushSwap.A
func (p *PushSwap) Init(length int) {
	p.Length = length
	p.Ops = 0
	p.Print = 0
	p.A = &Stack{List: nil, Length: 0}
	p.B = &Stack{List: nil, Length: 0}
	for i := 0; i < p.Length; i++ {
		p.A.PushBack(p.Array[i])
	}
}

// PrintInner prints Push Swap's Stack A, B and ops
func (p PushSwap) PrintInner() {
	fmt.Print("Stack A: ")
	for tmpNode := p.A.List; tmpNode != nil; tmpNode = tmpNode.Next {
		fmt.Printf("%d ", tmpNode.Number)
	}
	fmt.Printf("\nStack B: ")
	for tmpNode := p.B.List; tmpNode != nil; tmpNode = tmpNode.Next {
		fmt.Printf("%d ", tmpNode.Number)
	}
	fmt.Printf("\nops: %d\n", p.Ops)
}

// SyncedWithArray checks whether stack A has the same order as PushSwap.Array
// Old Name: true_sorted
func (p *PushSwap) SyncedWithArray() bool {
	i := 0

	for p.Array[i] != p.A.List.Number {
		i++
	}
	for tmpNode := p.A.List; tmpNode != nil; tmpNode = tmpNode.Next {
		if tmpNode.Number != p.Array[i] {
			return false
		}
		i++
	}
	return true
}

/*               Operations             */

/*             Swap Operations          */

// SwapStack swaps first two numbers of given a stack
// If swap cannot be done returns false
func SwapStack(stack *Stack) bool {
	var tmpNode *Node

	if stack.Length < 2 {
		return false
	}
	tmpNode = stack.List
	stack.List = tmpNode.Next

	tmpNode.Next = stack.List.Next
	stack.List.Next = tmpNode
	return true
}

// SwapA Stack
func (p *PushSwap) SwapA() {
	SwapStack(p.A)
	if p.Print == 1 {
		fmt.Println("sa")
	}
	p.Ops++
}

// SwapB Stack
func (p *PushSwap) SwapB() {
	SwapStack(p.B)
	if p.Print == 1 {
		fmt.Println("sb")
	}
	p.Ops++
}

// SwapAB swap both A and B Stacks
func (p *PushSwap) SwapAB() {
	SwapStack(p.A)
	SwapStack(p.B)
	if p.Print == 1 {
		fmt.Println("ss")
	}
	p.Ops++
}

/*             Push Operations          */

// PushA operation
// Will not push if Stack B is empty
func (p *PushSwap) PushA() {
	if p.Print == 1 {
		fmt.Println("pa")
	}
	if p.B.Length > 0 {
		numberPtr := p.B.Pop()
		p.A.PushFront(*numberPtr)
	}
	p.Ops++
}

// PushB operation
// Will not push if Stack A is empty
func (p *PushSwap) PushB() {
	if p.Print == 1 {
		fmt.Println("pb")
	}
	if p.A.Length > 0 {
		numberPtr := p.A.Pop()
		p.B.PushFront(*numberPtr)
	}
	p.Ops++
}

/*             Rotate Operations          */

// RotateStack rotates stack one node next
// Example:
// 0 1 2 3 4 5 6 - stack
// After RotateStack:
// 1 2 3 4 5 6 0
func RotateStack(stack *Stack) bool {
	var tmpNode, head *Node

	if stack.Length < 2 {
		return false
	}
	head = stack.List
	stack.List = stack.List.Next
	for tmpNode = stack.List; tmpNode.Next != nil; tmpNode = tmpNode.Next {
		// Move pointer tmpNode to the last Node of Stack
	}
	tmpNode.Next = head
	head.Next = nil
	return true
}

// RotateA operation
func (p *PushSwap) RotateA() {
	RotateStack(p.A)
	if p.Print == 1 {
		fmt.Println("ra")
	}
	p.Ops++
}

// RotateB operation
func (p *PushSwap) RotateB() {
	RotateStack(p.B)
	if p.Print == 1 {
		fmt.Println("rb")
	}
	p.Ops++
}

// RotateAB operation
func (p *PushSwap) RotateAB() {
	RotateStack(p.A)
	RotateStack(p.B)
	if p.Print == 1 {
		fmt.Println("rr")
	}
	p.Ops++
}

/*             Rev Rotate Operations          */

// RevRotateStack reverse-rotates stack one node next
// Example:
// 0 1 2 3 4 5 6 - stack
// After RevRotateStack:
// 6 0 1 2 3 4 5
func RevRotateStack(stack *Stack) bool {
	var tmpNode, penultimate, tail *Node

	if stack.Length < 2 {
		return false
	}
	for tmpNode = stack.List; tmpNode.Next.Next != nil; tmpNode = tmpNode.Next {
		// Move pointer tmpNode to the penultimate Node of Stack
	}
	penultimate = tmpNode
	tail = penultimate.Next

	penultimate.Next = nil
	tail.Next = stack.List
	stack.List = tail

	return true
}

// RevRotateA operation
func (p *PushSwap) RevRotateA() {
	RevRotateStack(p.A)
	if p.Print == 1 {
		fmt.Println("rra")
	}
	p.Ops++
}

// RevRotateB operation
func (p *PushSwap) RevRotateB() {
	RevRotateStack(p.B)
	if p.Print == 1 {
		fmt.Println("rrb")
	}
	p.Ops++
}

// RevRotateAB operation
func (p *PushSwap) RevRotateAB() {
	RevRotateStack(p.A)
	RevRotateStack(p.B)
	if p.Print == 1 {
		fmt.Println("rrr")
	}
	p.Ops++
}
