package minimumstringlengthafterremovingsubstrings

import "fmt"

// https://leetcode.com/problems/minimum-string-length-after-removing-substrings/description/?envType=daily-question&envId=2024-10-07
type MinString struct {
}
type MinStringInterface interface {
	IsValid(parantheses string) bool
}

func MinLength(str string) int {
	counter := 1
	str = "ABFCACDB"

	for 0 < counter {
		counter = 0
		for i := 0; i < len(str)-1; i++ {
			if (str[i] == 'A' && str[i+1] == 'B') || (str[i] == 'C' && str[i+1] == 'D') {
				str = str[:i] + str[i+2:]
				fmt.Println(str)
				counter++
			}
		}
	}

	return len(str)
}
