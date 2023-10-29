package algorithms

func SortBooks(arr []string) []string {
	for j := 1; j < len(arr); j++ {
		actual := arr[j]

		i := j - 1
		for i >= 0 && arr[i] > actual {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = actual
	}
	return arr
}
