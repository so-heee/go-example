package unorderedoutput

import "fmt"

// Unordered は1,2,3の順でプリントします
func Unordered() {
	for _, v := range []int{1, 2, 3} {
		fmt.Println(v)
	}
}
