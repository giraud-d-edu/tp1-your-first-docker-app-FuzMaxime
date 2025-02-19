package groupieTracker

// Sorting system 
func Sorts(sort string, data []Artists, left int, right int) {
	if sort == "a-z" {
		if left < right {
			pivot := PartitionByName(sort, data, left, right)
			Sorts(sort, data, left, pivot-1)
			Sorts(sort, data, pivot+1, right)

		}
	} else if sort == "date" {
		if left < right {
			pivot := PartitionByDate(sort, data, left, right)
			Sorts(sort, data, left, pivot-1)
			Sorts(sort, data, pivot+1, right)
		}
	} else if sort == "members" {
		if left < right {
			pivot := PartitionByMembers(sort, data, left, right)
			Sorts(sort, data, left, pivot-1)
			Sorts(sort, data, pivot+1, right)
		}
	}
}

/* -------------- sort by name -------------- */

func PartitionByName(sort string, data []Artists, left int, right int) int {
	pivot := data[right]
	i := left
	for j := left; j < right; j++ {
		if data[j].Name < pivot.Name {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[right] = data[right], data[i]
	return i
}

/* -------------- sort by date -------------- */

func PartitionByDate(sort string, data []Artists, left int, right int) int {
	pivot := data[right]
	i := left
	for j := left; j < right; j++ {
		if data[j].CreationDate < pivot.CreationDate {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[right] = data[right], data[i]
	return i
}

/* -------------- sort by members -------------- */

func PartitionByMembers(sort string, data []Artists, left int, right int) int {
	pivot := data[right]
	i := left
	for j := left; j < right; j++ {
		if len(data[j].Members) < len(pivot.Members) {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[right] = data[right], data[i]
	return i
}
