package main

import "fmt"

func myFunc(A []string, B []string) ([]string, []string) {

	var union []string // A U B
	var sDiff []string // (A \ B) U (B \ A)

	union = A

	// A in B
	for i := 0; i < len(A); i++ {
		AinB := false
		for j := 0; j < len(A); j++ {
			if AinB == false && A[i] == B[j] {
				AinB = true
			}
		}

		if AinB == false {
			sDiff = append(sDiff, A[i])
		}
	}

	// B in A
	for i := 0; i < len(B); i++ {
		BinA := false
		for j := 0; j < len(A); j++ {
			if BinA == false && B[i] == A[j] {
				BinA = true
			}
		}

		if BinA == false {
			union = append(union, B[i])
			sDiff = append(sDiff, B[i])
		}
	}

	return union, sDiff
}

func main() {
	arrayA := []string{"a", "b", "c"}
	arrayB := []string{"b", "c", "d"}

	fmt.Println(myFunc(arrayA, arrayB))
}
