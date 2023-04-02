package main

import (
	"fmt"
	binary_search "leetcode/binary-search"
	valid_parantheses "leetcode/vaild-parantheses"
)

func main() {
	//valid parantheses
	s := valid_parantheses.NewValidParentheses()
	fmt.Println(s.IsValid("()"))

	//binary search
	search := binary_search.NewBinarySearch()
	fmt.Println(search.FindIndex([]int{1, 2, 3, 5, 7, 9, 12}, 3))

}
