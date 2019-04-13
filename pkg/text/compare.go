package text

func LeftMismatch(s1 []rune, first1 int, last1 int, s2 []rune, first2 int, last2 int) (int, int) {
	for first1 != last1 && first2 != last2 && s1[first1] == s2[first2] {
		first1++
		first2++
	}
	return first1, first2
}

func RightMismatch(s1 []rune, first1 int, last1 int, s2 []rune, first2 int, last2 int) (int, int) {
	first1 = len(s1) - 1 - first1
	first2 = len(s2) - 1 - first2
	last1 = len(s1) - 1 - last1
	last2 = len(s2) - 1 - last2
	for first1 != last1 && first2 != last2 && s1[first1] == s2[first2] {
		first1--
		first2--
	}
	return len(s1) - 1 - first1, len(s2) - 1 - first2
}
