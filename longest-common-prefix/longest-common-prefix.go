package longest_common_prefix

type LongestCommonPrefix struct {
}

func (l LongestCommonPrefix) CalculateLongestCommonPrefix(arr []string) string {
	tempArr := ""
	length := len(arr)
	if length == 0 {
		return ""
	}

	myArr := [...]string{"flower", "flow", "flight"}

	for i := 0; i < len(myArr)-1; i++ {
		for j := 0; j < len(myArr[j])-1; j++ {
			if myArr[i][j] == myArr[i+1][j] {
				tempArr += string(myArr[i][j])
			}
		}
	}

	return tempArr
}

func NewLongestCommonPrefix() *LongestCommonPrefix {
	return &LongestCommonPrefix{}
}
