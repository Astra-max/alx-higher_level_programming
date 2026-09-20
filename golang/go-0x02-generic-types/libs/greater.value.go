package main

import (
	"cmp"
	"fmt"
)

// GetDominantValue accepts a slice of any ordered type (int, float, string, or uintptr memory addresses).
// It returns the "greater value" or the one that comes first depending on the type.
func GetDominantValue[T cmp.Ordered](slice []T) (T, error) {
	var zero T
	if len(slice) == 0 {
		return zero, fmt.Errorf("cannot find value in an empty slice")
	}

	// Assume the first element is the target
	dominant := slice[0]

	for _, val := range slice {
		// cmp.Compare returns:
		// -1 if dominant < val
		//  0 if dominant == val
		// +1 if dominant > val
		if cmp.Compare(val, dominant) > 0 {
			dominant = val
		}
	}

	return dominant, nil
}

func main() {
	// 1. Testing with Integers
	intSlice := []int{5, 22, 14, 9}
	maxInt, _ := GetDominantValue(intSlice)
	fmt.Printf("Max Int: %v\n", maxInt) // Outputs: 22

	// 2. Testing with Floats
	floatSlice := []float64{3.14, 1.05, 7.89, 4.56}
	maxFloat, _ := GetDominantValue(floatSlice)
	fmt.Printf("Max Float: %v\n", maxFloat) // Outputs: 7.89

	// 3. Testing with Strings (Returns what comes LAST alphabetically, "greater" value)
	stringSlice := []string{"apple", "zebra", "banana"}
	maxStr, _ := GetDominantValue(stringSlice)
	fmt.Printf("Max String: %v\n", maxStr) // Outputs: zebra

	// 4. Testing with Memory Addresses (uintptr counts as cmp.Ordered)
	a, b := 1, 2
	addrSlice := []uintptr{uintptr(unsafe.Pointer(&a)), uintptr(unsafe.Pointer(&b))}
	maxAddr, _ := GetDominantValue(addrSlice)
	fmt.Printf("Highest Memory Address: %v\n", maxAddr)
}
