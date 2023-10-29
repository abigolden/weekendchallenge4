package algorithms

// Imagine you have an array of numbers. Write an algorithm to remove all numbers
// less than five from the array.
// Input: '[10,4,5,8,2,9]'
// Output: '[10,5,8,9]'

func FilterLt5(list []int) []int {
	ilteredList := []int{}

	for _, num := range list {
		if num >= 5 {
			filteredList = append(filteredList, num)
		}
	}

	return filteredList
}
