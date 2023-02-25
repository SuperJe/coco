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
