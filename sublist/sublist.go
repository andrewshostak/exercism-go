package sublist

type Relation string

const (
	equal     Relation = "equal"
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
	unequal   Relation = "unequal"
)

func Sublist(listOne, listTwo []int) Relation {
	switch {
	case len(listTwo) > len(listOne):
		diff := len(listTwo) - len(listOne)
		for i := 0; i <= diff; i++ {
			if compareSameLength(listOne, listTwo[i:(i+len(listOne))]) == equal {
				return sublist
			}
		}
	case len(listTwo) < len(listOne):
		diff := len(listOne) - len(listTwo)
		for i := 0; i <= diff; i++ {
			if compareSameLength(listTwo, listOne[i:(i+len(listTwo))]) == equal {
				return superlist
			}
		}
	}

	return compareSameLength(listOne, listTwo)
}

func compareSameLength(listOne, listTwo []int) Relation {
	for i := range listOne {
		if listOne[i] != listTwo[i] {
			return unequal
		}
	}
	return equal
}
