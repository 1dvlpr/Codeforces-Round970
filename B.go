## Square or Not
package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func solve(n int64, s string) {
	var o, z int64
 
	for _, x := range s {
		if x == '0' {
			z++
		} else {
			o++
		}
	}
 
	for i := int64(1); i*i <= n; i++ {
		if n%i == 0 {
			r := i
			c := n / i
 
			if o == 2*(r+c-2) && z == (r*c-o) {
				if r == c {
					fmt.Println("YES")
					return
				}
			}
		}
	}
 
	fmt.Println("NO")
}
 
func main() {
	reader := bufio.NewReader(os.Stdin)
 
	const TASK = "bt"
 
	if _, err := os.Stat(TASK + ".INP"); err == nil {
		fin, _ := os.Open(TASK + ".INP")
		fout, _ := os.Create(TASK + ".OUT")
		defer fin.Close()
		defer fout.Close()
		reader = bufio.NewReader(fin)
		writer := bufio.NewWriter(fout)
		defer writer.Flush()
 
		var t int64
		fmt.Fscanf(reader, "%d\n", &t)
 
		for t > 0 {
			t--
 
			var n int64
			var s string
			fmt.Fscanf(reader, "%d %s\n", &n, &s)
 
			solve(n, s)
		}
	} else {
		var t int64
		fmt.Fscanf(reader, "%d\n", &t)
 
		for t > 0 {
			t--
 
			var n int64
			var s string
			fmt.Fscanf(reader, "%d %s\n", &n, &s)
 
			solve(n, s)
		}
	}
}
