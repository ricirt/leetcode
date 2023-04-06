package rotate_string

import (
	Helper "leetcode/Helpers"
)

type RotateString struct{}

type IRotateString interface {
	RotateString(string, string) bool
}

func (this *RotateString) RotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(goal); j++ {
			if s[i] == goal[j] {
				if IsEqual(i-j, s, goal) {
					return true
				} else {
					continue
				}
			}
		}
	}

	return false
}

func IsEqual(difference int, s string, goal string) bool {
	h := Helper.NewHelpers()

	for i := 0; i < len(s); i++ {
		ss := (i + (h.AbsVal(difference) % len(s))) % len(s)

		if s[i] == goal[ss] {
			continue
		} else {
			return false
		}
	}
	return true
}

func NewRotateString() IRotateString {
	return &RotateString{}
}
