package main

import (
	"fmt"
	valid_parantheses "leetcode/vaild-parantheses"
)

func main() {
	s := valid_parantheses.NewValidParentheses()
	fmt.Print(s.IsValid("()"))
}
