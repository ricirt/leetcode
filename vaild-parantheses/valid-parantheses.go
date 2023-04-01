package valid_parantheses

type ValidParentheses struct {
}
type IValidParantheses interface {
	IsValid()
}

func (v *ValidParentheses) IsValid(parantheses string) bool {

	stack := []byte{}
	for i := 0; i < len(parantheses); i++ {
		if parantheses[i] == '(' || parantheses[i] == '{' || parantheses[i] == '[' {
			stack = append(stack, parantheses[i])
		} else if len(stack) == 0 {
			return false
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if parantheses[i] == ')' && top != '(' {
				return false
			} else if parantheses[i] == '}' && top != '{' {
				return false
			} else if parantheses[i] == ']' && top != '[' {
				return false
			}
		}
	}
	return len(stack) == 0
}

func NewValidParentheses() ValidParentheses {
	return ValidParentheses{}
}
