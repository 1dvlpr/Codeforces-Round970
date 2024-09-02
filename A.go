## Sakurako Exam


package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func solve(a, b int64) {
	// a 1's and b 2's
 
	if a%2 != 0 {
		fmt.Println("NO")
	} else {
		if a != 0 {
			fmt.Println("YES")
		} else {
			if b%2 == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
 
func main() {
	reader := bufio.NewReader(os.Stdin)
 
	var t int64
	fmt.Fscanf(reader, "%d\n", &t)
 
	for t > 0 {
		t--
 
		var a, b int64
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
 
		solve(a, b)
	}
}


// That's HOw the another way , you can do it: 
package main

import (
	"bufio"
	"fmt"
	"os"
)

// solve determines if it's possible to have a sequence of 'a' 1's and 'b' 2's
// that satisfies the given conditions. Specifically, it checks if 'a' is even,
// or if 'a' is zero and 'b' is even.
func solve(a, b int64) {
	// If 'a' is odd, it's impossible to satisfy the conditions, so return "NO"
	if a%2 != 0 {
		fmt.Println("NO")
	} else {
		// If 'a' is even and non-zero, return "YES"
		if a != 0 {
			fmt.Println("YES")
		} else {
			// If 'a' is zero, check if 'b' is even to return "YES", otherwise "NO"
			if b%2 == 0 {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int64
	// Read the number of test cases
	fmt.Fscanf(reader, "%d\n", &t)

	// Process each test case
	for t > 0 {
		t--

		var a, b int64
		// Read the values of 'a' and 'b' for each test case
		fmt.Fscanf(reader, "%d %d\n", &a, &b)

		// Call solve to determine the result for the current test case
		solve(a, b)
	}
}

