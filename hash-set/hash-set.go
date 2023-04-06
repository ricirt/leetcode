package hash_set

type MyHashSet struct {
	value []int
}

func Constructor() *MyHashSet {
	return new(MyHashSet)
}

func (this *MyHashSet) Add(key int) {
	if this.FindIndex(key) > 0 {
		return
	}
	this.value = append(this.value, key)
}

func (this *MyHashSet) Remove(key int) {
	index := this.FindIndex(key)
	if index < 0 {
		panic("value not found")
	}
	Delete(this.value, index, index+1)
}

func (this *MyHashSet) Contains(key int) bool {

	if this.FindIndex(key) < 0 {
		return false
	}
	return true
}

func (this *MyHashSet) FindIndex(value int) int {
	for i := 0; i < len(this.value); i++ {
		if value == this.value[i] {
			return i
		}
	}
	return -1
}

// golang.org/x/exp/slices
func Delete[S ~[]E, E any](s S, i, j int) S {
	//_ = s[i:j] // bounds check

	return append(s[:i], s[j:]...)
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
