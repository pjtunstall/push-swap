package main

// Just not working yet. It always passes, even when told to fail. Why?

// import (
// 	"fmt"
// 	"math/rand"
// 	"push-swap/ps"
// 	"strconv"
// 	"strings"
// 	"testing"
// 	"time"
// )

// func generateRandomNumberString() string {
// 	rand.Seed(time.Now().UnixNano())
// 	numbers := make([]string, 7)
// 	for i := range numbers {
// 		numbers[i] = strconv.Itoa(rand.Int())
// 	}
// 	return strings.Join(numbers, " ")
// }

// func TestOrionVTUrk(t *testing.T) {
// 	failCount := 0
// 	input := generateRandomNumberString()
// 	fmt.Println(input)

// 	// orion modifies the stacks.
// 	a, _ := ps.NewStack(input)
// 	b, _ := ps.NewStack("")
// 	orion := bucket3(&a, &b)

// 	// turk modifies the stack
// 	a, _ = ps.NewStack(input)
// 	b, _ = ps.NewStack("")
// 	turk := turk(&a, &b)

// 	if len(turk) < len(orion) {
// 		failCount++
// 		t.Errorf("Test failed for input: %s. Length of `turk` is %v, less than length of `orion` %v", input, len(turk), len(orion))
// 	}
// }

// func turk(a, b *ps.Stack) []string {
// 	result := []string{}
// 	ps.Px(b, a)
// 	ps.Px(b, a)
// 	result = append(result, "pb", "pb")

// 	nums := b.GetNumsSlice()
// 	if nums[0] < nums[1] {
// 		ps.Sx(b)
// 		result = append(result, "sb")
// 	}

// 	result = append(result, insert(a, b, 3, true)...)

// 	_, rotatable := three(a.Nums)
// 	if !rotatable {
// 		ps.Sx(a)
// 		result = append(result, "sa")

// 	}

// 	result = append(result, insert(b, a, 0, false)...)

// 	result = append(result, justRotate(*a)...)
// 	ps.Run(a, b, justRotate(*a))

// 	return result
// }
