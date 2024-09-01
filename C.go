## Longest Good Array

package main
 
import (
	"bufio"
	"fmt"
	"math"
	"os"
)
 
func main() {
	reader := bufio.NewReader(os.Stdin)
 
	var t int
	fmt.Fscanf(reader, "%d\n", &t)
 
	for t > 0 {
		t--
 
		var l, r int64
		fmt.Fscanf(reader, "%d %d\n", &l, &r)
 
		diff := r - l
		k := int64((1 + math.Sqrt(1+8*float64(diff))) / 2)
 
		fmt.Println(k)
	}
}
