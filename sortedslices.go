package sortedslices

import (
	"cmp"
	"slices"
	"sort"
)

// InsertFunc inserts item in correct position and returns a sorted slice.
func Insert[S ~[]E, E cmp.Ordered](sorted S, item E) S {
	i, _ := slices.BinarySearch(sorted, item)
	if i == len(sorted)-1 && sorted[i] < item {
		return append(sorted, item)
	}
	return append(sorted[:i], append(S{item}, sorted[i:]...)...)
}

// RemoveFunc removes item in a sorted slice.
func Remove[S ~[]E, E cmp.Ordered](sorted S, item E) S {
	i, found := slices.BinarySearch(sorted, item)
	if found {
		return append(sorted[:i], sorted[i+1:]...)
	}
	return sorted
}

// IterateOver iterates over input sorted slices and calls callback with each items in ascendant order.
func IterateOver[S ~[]E, E cmp.Ordered](callback func(item E, srcIndex int), sorted ...S) {
	sourceSlices := make([]S, 0, len(sorted))
	for _, src := range sorted {
		if len(src) > 0 {
			sourceSlices = append(sourceSlices, src)
		}
	}
	sourceSliceCount := len(sourceSlices)
	if sourceSliceCount == 0 {
		return
	} else if sourceSliceCount == 1 {
		for i, value := range sourceSlices[0] {
			callback(value, i)
		}
		return
	}
	indexes := make([]int, sourceSliceCount)
	sliceIndex := make([]int, sourceSliceCount)
	for i := range sourceSlices {
		sliceIndex[i] = i
	}
	index := 0
	for {
		minSlice := 0
		minItem := sourceSlices[0][indexes[0]]
		for i := 1; i < sourceSliceCount; i++ {
			if sourceSlices[i][indexes[i]] < minItem {
				minSlice = i
				minItem = sourceSlices[i][indexes[i]]
			}
		}
		callback(minItem, sliceIndex[minSlice])
		index++
		indexes[minSlice]++
		if indexes[minSlice] == len(sourceSlices[minSlice]) {
			sourceSlices = append(sourceSlices[:minSlice], sourceSlices[minSlice+1:]...)
			indexes = append(indexes[:minSlice], indexes[minSlice+1:]...)
			sliceIndex = append(sliceIndex[:minSlice], sliceIndex[minSlice+1:]...)
			sourceSliceCount--
			if len(sourceSlices) == 1 {
				slice := sourceSlices[0]
				for i := indexes[0]; i < len(slice); i++ {
					callback(slice[i], sliceIndex[0])
				}
				return
			}
		}
	}
}

// UnionFunc unions sorted slices and returns new slices.
func Union[S ~[]E, E cmp.Ordered](sorted ...S) S {
	length := 0
	sourceSlices := make([]S, 0, len(sorted))
	for _, src := range sorted {
		if len(src) > 0 {
			length += len(src)
			sourceSlices = append(sourceSlices, src)
		}
	}
	if length == 0 {
		return nil
	} else if len(sourceSlices) == 1 {
		return sourceSlices[0]
	}
	result := make([]E, length)
	sourceSliceCount := len(sourceSlices)
	indexes := make([]int, sourceSliceCount)
	index := 0
	for {
		minSlice := 0
		minItem := sourceSlices[0][indexes[0]]
		for i := 1; i < sourceSliceCount; i++ {
			if sourceSlices[i][indexes[i]] < minItem {
				minSlice = i
				minItem = sourceSlices[i][indexes[i]]
			}
		}
		result[index] = minItem
		index++
		indexes[minSlice]++
		if indexes[minSlice] == len(sourceSlices[minSlice]) {
			sourceSlices = append(sourceSlices[:minSlice], sourceSlices[minSlice+1:]...)
			indexes = append(indexes[:minSlice], indexes[minSlice+1:]...)
			sourceSliceCount--
			if len(sourceSlices) == 1 {
				copy(result[index:], sourceSlices[0][indexes[0]:])
				return result
			}
		}
	}
}

// Difference creates difference group of sorted slices and returns.
func Difference[S ~[]E, E cmp.Ordered](sorted1, sorted2 S) S {
	var result S
	var i, j int
	for i < len(sorted1) && j < len(sorted2) {
		if sorted1[i] < sorted2[j] {
			result = append(result, sorted1[i])
			i++
		} else if sorted1[i] > sorted2[j] {
			j++
		} else {
			i++
			j++
		}
	}
	result = append(result, sorted1[i:]...)
	return result
}

// IntersectionFunc creates intersection group of sorted slices and returns.
func Intersection[S ~[]E, E cmp.Ordered](sorted ...S) S {
	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i]) < len(sorted[j])
	})
	var result S
	if len(sorted[0]) == 0 {
		return result
	}
	cursors := make([]int, len(sorted))
	terminate := false
	for _, value := range sorted[0] {
		needIncrement := false
		for i := 1; i < len(sorted); i++ {
			found := false
			for j := cursors[i]; j < len(sorted[i]); j++ {
				valueOfOtherSlice := sorted[i][cursors[i]]
				if valueOfOtherSlice < value {
					cursors[i] = j + 1
				} else if value < valueOfOtherSlice {
					needIncrement = true
					break
				} else {
					found = true
					break
				}
			}
			if needIncrement {
				break
			}
			if !found {
				terminate = true
				break
			}
		}
		if terminate {
			break
		}
		if !needIncrement {
			result = append(result, value)
		}
	}
	return result
}

func Min[S ~[]E, E cmp.Ordered](x S) E {
	if len(x) < 1 {
		panic("sortedslices.Min: empty list")
	}
	return x[0]
}

func Max[S ~[]E, E cmp.Ordered](x S) E {
	if len(x) < 1 {
		panic("sortedslices.Max: empty list")
	}
	return x[len(x)-1]
}
