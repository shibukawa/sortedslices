package sortedslices_test

import (
	"cmp"
	"reflect"
	"slices"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/shibukawa/sortedslices"
)

func TestInsert(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("insert returns new sorted slices", prop.ForAll(func(input []int) bool {
		expected := make([]int, len(input))
		copy(expected, input)
		slices.Sort(expected)

		value := input[0]
		array := input[1:]
		slices.Sort(array)

		inserted := sortedslices.Insert(array, value)

		return reflect.DeepEqual(expected, inserted)
	}, numSliceGenerator))

	properties.TestingRun(t)
}

func TestRemove(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("removes item of array", prop.ForAll(func(input []int) bool {
		value := input[0]
		slices.Sort(input)

		removedArray := sortedslices.Remove(input, value)

		if len(removedArray) != len(input)-1 {
			return false
		}
		_, found := slices.BinarySearch(removedArray, value)
		return !found
	}, numSliceGenerator))

	properties.TestingRun(t)
}

func TestIterateOver(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("iterate over item of arrays", prop.ForAll(func(input1, input2, input3 []int) bool {
		var result []int
		slices.SortFunc(input1, cmp.Compare[int])
		slices.SortFunc(input2, cmp.Compare[int])
		slices.SortFunc(input3, cmp.Compare[int])
		sortedslices.IterateOver(func(v, index int) {
			result = append(result, v)
		}, input1, input2, input3)
		return slices.IsSorted(result)
	}, numSliceGenerator, numSliceGenerator, numSliceGenerator))

	properties.TestingRun(t)
}

func TestUnion(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("union item of arrays", prop.ForAll(func(input1, input2, input3 []int) bool {
		slices.Sort(input1)
		slices.Sort(input2)
		slices.Sort(input3)
		result := sortedslices.Union(input1, input2, input3)
		if !slices.IsSorted(result) {
			return false
		}
		for _, e := range result {
			if _, found := slices.BinarySearchFunc(input1, e, cmp.Compare[int]); found {
				continue
			}
			if _, found := slices.BinarySearchFunc(input2, e, cmp.Compare[int]); found {
				continue
			}
			if _, found := slices.BinarySearchFunc(input3, e, cmp.Compare[int]); found {
				continue
			}
			return false
		}
		return true
	}, numSliceGenerator, numSliceGenerator, numSliceGenerator))

	properties.TestingRun(t)
}

func TestDifference(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("difference between arrays", prop.ForAll(func(input1, input2 []int) bool {
		slices.Sort(input1)
		slices.Sort(input2)
		result := sortedslices.Difference(input1, input2)
		if !slices.IsSorted(result) {
			return false
		}
		for _, e := range result {
			if _, found := slices.BinarySearchFunc(input1, e, cmp.Compare[int]); !found {
				return false
			}
			if _, found := slices.BinarySearchFunc(input2, e, cmp.Compare[int]); found {
				return false
			}
		}
		return true
	}, numSliceGenerator, numSliceGenerator))

	properties.TestingRun(t)
}

func TestIntersection(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("union item of arrays", prop.ForAll(func(input1, input2, input3 []int) bool {
		slices.Sort(input1)
		slices.Sort(input2)
		slices.Sort(input3)
		result := sortedslices.Intersection(input1, input2, input3)
		if !slices.IsSorted(result) {
			return false
		}
		for _, e := range result {
			if _, found := slices.BinarySearch(input1, e); !found {
				return false
			}
			if _, found := slices.BinarySearch(input2, e); !found {
				return false
			}
			if _, found := slices.BinarySearch(input3, e); !found {
				return false
			}
		}
		return true
	}, numSliceGenerator, numSliceGenerator, numSliceGenerator))

	properties.TestingRun(t)
}

func TestMin(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("min item of array", prop.ForAll(func(input []int) bool {
		slices.Sort(input)

		return sortedslices.Min(input) == slices.Min(input)
	}, numSliceGenerator))

	properties.TestingRun(t)
}

func TestMax(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("max item of array", prop.ForAll(func(input []int) bool {
		slices.Sort(input)

		return sortedslices.Max(input) == slices.Max(input)
	}, numSliceGenerator))

	properties.TestingRun(t)
}
