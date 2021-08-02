package listops

type IntList []int
type predFunc func(int) bool
type unaryFunc func(int) int
type binFunc func(int, int) int

func (list IntList) Length() int {
	length := 0
	for range list {
		length++
	}
	return length
}

func (list IntList) Append(arr []int) IntList {
	return append(list, arr...)
}

func (list IntList) Reverse() IntList {
	reversed := make([]int, len(list))
	for i := range list {
		reversed[len(list)-i-1] = list[i]
	}
	return reversed
}

func (list IntList) Concat(arrs []IntList) IntList {
	flatten := list
	for i := range arrs {
		flatten = append(flatten, arrs[i]...)
	}
	return flatten
}

func (list IntList) Filter(fn predFunc) IntList {
	filtered := make([]int, 0)
	for i := range list {
		if fn(list[i]) {
			filtered = append(filtered, list[i])
		}
	}
	return filtered
}

func (list IntList) Map(fn unaryFunc) IntList {
	for i := range list {
		list[i] = fn(list[i])
	}
	return list
}

func (list IntList) Foldr(fn binFunc, acc int) int {
	for i := len(list) - 1; i >= 0; i-- {
		acc = fn(list[i], acc)
	}
	return acc
}

func (list IntList) Foldl(fn binFunc, acc int) int {
	for i := range list {
		acc = fn(acc, list[i])
	}
	return acc
}
