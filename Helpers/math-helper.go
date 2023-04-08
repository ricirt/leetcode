package Helpers

type IMathHelper interface {
	AbsVal(int) int
}
type MathHelper struct{}

func (h *MathHelper) AbsVal(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func NewHelpers() IMathHelper {
	return &MathHelper{}
}
