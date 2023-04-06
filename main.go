package main

import (
	"fmt"
	binary_search "leetcode/binary-search"
	hash_set "leetcode/hash-set"
	rotate_string "leetcode/rotate-string"
	valid_parantheses "leetcode/vaild-parantheses"
)

func main() {
	//valid parantheses
	s := valid_parantheses.NewValidParentheses()
	fmt.Println(s.IsValid("()"))

	//binary search
	search := binary_search.NewBinarySearch()
	fmt.Println(search.FindIndex([]int{1, 2, 3, 5, 7, 9, 12}, 3))

	//RotateString
	rotateString := rotate_string.NewRotateString()
	fmt.Println(rotateString.RotateString("abcde", "cdeab"))

	obj := hash_set.Constructor()
	obj.Add(1)
	obj.Add(2)
	obj.Remove(1)
	param_3 := obj.Contains(2)
	fmt.Println(param_3)
}
