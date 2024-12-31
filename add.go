// Package add contains a function for adding integers
package add

import (
	"golang.org/x/exp/constraints"
)

// Add takes two parameters of any integer or float type, adds them, and returns the sum
//
// for further details [see]
//
// [see]: https://www/mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

type Number interface {
	constraints.Integer | constraints.Float
}
