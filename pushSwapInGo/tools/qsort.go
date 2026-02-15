/*
* @Author: Tlekbai Ali
* @Date:   2020-01-04 15:34:40
* @Last Modified by:   Tlekbai Ali
* @Last Modified time: 2020-01-04 16:24:41
 */

package tools

// Qsort algorithm implementation
func Qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := left + (right-left)/2

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	Qsort(a[:left])
	Qsort(a[left+1:])

	return a
}
