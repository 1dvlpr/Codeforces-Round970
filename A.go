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
