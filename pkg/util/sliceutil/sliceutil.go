package sliceutil

type StringSet map[string]struct{}

// ToStringSet 数组转set
func ToStringSet(src []string) StringSet {
	set := make(StringSet, 0)
	for _, s := range src {
		set[s] = struct{}{}
	}
	return set
}

// Contain src是否包含target
func Contain(src []string, target string) bool {
	for _, s := range src {
		if s == target {
			return true
		}
	}
	return false
}
