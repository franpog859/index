package index

func getAll(slice []interface{}, _item interface{}) (indexes []int, ok bool) {
	indexes = []int{}
	for index, item := range slice {
		if item == _item {
			indexes = append(indexes, index)
		}
	}

	return indexes, len(indexes) != 0
}
