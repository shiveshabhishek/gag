package gag

func TwoSum(arrayElem []int, sum int) []int {
	for i := 1; i < len(arrayElem); i++ {
		for j := 0; j < i; j++ {
			if arrayElem[i]+arrayElem[j] == sum {
				return []int{j, i}
			}
		}
	}
	return nil
}
