/*
* @Author: Tlekbai Ali
* @Date:   2020-01-04 14:43:58
* @Last Modified by:   Tlekbai Ali
* @Last Modified time: 2020-01-04 16:28:40
 */

package pushswap

import (
	tools "../tools"
)

// SortFirstTwoA sorts first two numbers of stack A
func (p *PushSwap) SortFirstTwoA() {
	if p.A.Length < 2 {
		return
	}
	if p.A.List.Number > p.A.List.Next.Number {
		p.SwapA()
	}
}

// SortFirstTwoB sorts first two numbers of stack B
func (p *PushSwap) SortFirstTwoB() {
	if p.B.Length < 2 {
		return
	}
	if p.B.List.Number < p.B.List.Next.Number {
		p.SwapB()
	}
}

// SortFirstTwoAB sorts first two numbers of stacks A and B
func (p *PushSwap) SortFirstTwoAB() {
	ARevSorted := p.A.IsRevSorted(2)
	BSorted := p.B.IsSorted(2)
	if ARevSorted && BSorted {
		p.SwapAB()
	} else if BSorted {
		p.SortFirstTwoB()
	} else if ARevSorted {
		p.SortFirstTwoA()
	}
}

// SortFirstThreeA sorts first three numbers of stack A
func (p *PushSwap) SortFirstThreeA() {
	var a, b, c, max int

	if p.A.Length < 3 || p.A.IsSorted(3) {
		return
	}
	a = p.A.List.Number
	b = p.A.List.Next.Number
	c = p.A.List.Next.Next.Number
	max = tools.Max(a, b, c)

	if max == c {
		p.SwapA()
		return
	}
	if p.A.Length == 3 {
		if max == a {
			p.RotateA()
		} else {
			p.RevRotateA()
		}
		if (max == a && b > c) || (max == b && c > a) {
			p.SwapA()
		}
		return
	}
	if max == a {
		p.SwapA()
	}
	p.RotateA()
	p.SwapA()
	p.RevRotateA()
	if (max == a && b > c) || (max == b && a > c) {
		p.SwapA()
	}
}

// SortFirstThreeB sorts first three numbers of stack B
func (p *PushSwap) SortFirstThreeB() {
	var a, b, c, min, max int

	if p.B.Length < 3 || p.B.IsRevSorted(3) {
		return
	}
	a = p.B.List.Number
	b = p.B.List.Next.Number
	c = p.B.List.Next.Next.Number
	min = tools.Min(a, b, c)
	max = tools.Max(a, b, c)

	if p.B.Length == 3 {
		if min == c {
			p.SwapB()
			return
		}
		if min == a {
			p.RotateB()
		} else {
			p.RevRotateB()
		}
		if (min == a && c > b) || (min == b && a > c) {
			p.SwapB()
		}
		return
	}
	if max == b {
		p.SwapB()
	}
	if max == b && a > c {
		return
	}
	p.RotateB()
	p.SwapB()
	p.RevRotateB()
	if max == c {
		p.SwapB()
	}
	if max == c && a < b {
		p.RotateB()
		p.SwapB()
		p.RevRotateB()
	}
}
