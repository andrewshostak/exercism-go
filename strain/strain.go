package strain

type Ints []int
type Lists [][]int
type Strings []string

func (in Ints) Keep(f func(int) bool) Ints {
	var n []int
	for i := range in {
		if f(in[i]) {
			n = append(n, in[i])
		}
	}
	return n
}

func (in Ints) Discard(f func(int) bool) Ints {
	var n Ints
	for i := range in {
		if !f(in[i]) {
			n = append(n, in[i])
		}
	}
	return n
}

func (l Lists) Keep(f func([]int) bool) Lists {
	var n Lists
	for i := range l {
		if f(l[i]) {
			n = append(n, l[i])
		}
	}
	return n
}

func (s Strings) Keep(f func(string) bool) Strings {
	var n Strings
	for i := range s {
		if f(s[i]) {
			n = append(n, s[i])
		}
	}
	return n
}
