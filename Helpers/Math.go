package Helpers

type IHelpers interface {
	AbsVal(int) int
}
type Helpers struct{}

func (h *Helpers) AbsVal(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func NewHelpers() IHelpers {
	return &Helpers{}
}
