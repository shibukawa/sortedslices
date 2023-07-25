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

func TestInsertFunc(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("insert returns new sorted slices", prop.ForAll(func(input []int) bool {
		expected := make([]int, len(input))
		copy(expected, input)
		slices.SortFunc(expected, cmp.Compare[int])

		value := input[0]
		array := input[1:]
		slices.SortFunc(array, cmp.Compare[int])

		inserted := sortedslices.InsertFunc(array, value, cmp.Compare[int])

		return reflect.DeepEqual(expected, inserted)
	}, numSliceGenerator))

	properties.TestingRun(t)
}

func TestRemoveFunc(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("removes item of array", prop.ForAll(func(input []int) bool {
		value := input[0]
		slices.SortFunc(input, cmp.Compare[int])

		removedArray := sortedslices.RemoveFunc(input, value, cmp.Compare[int])

		if len(removedArray) != len(input)-1 {
			return false
		}
		_, found := slices.BinarySearchFunc(removedArray, value, cmp.Compare[int])
		return !found
	}, numSliceGenerator))

	properties.TestingRun(t)
}

func TestIterateOverFunc(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("iterate over item of arrays", prop.ForAll(func(input1, input2, input3 []int) bool {
		var result []int
		slices.SortFunc(input1, cmp.Compare[int])
		slices.SortFunc(input2, cmp.Compare[int])
		slices.SortFunc(input3, cmp.Compare[int])
		sortedslices.IterateOverFunc(cmp.Compare[int], func(v, index int) {
			result = append(result, v)
		}, input1, input2, input3)
		return slices.IsSortedFunc(result, cmp.Compare[int])
	}, numSliceGenerator, numSliceGenerator, numSliceGenerator))

	properties.TestingRun(t)
}

func TestUnionFunc(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("union item of arrays", prop.ForAll(func(input1, input2, input3 []int) bool {
		slices.SortFunc(input1, cmp.Compare[int])
		slices.SortFunc(input2, cmp.Compare[int])
		slices.SortFunc(input3, cmp.Compare[int])
		result := sortedslices.UnionFunc(cmp.Compare[int], input1, input2, input3)
		if !slices.IsSortedFunc(result, cmp.Compare[int]) {
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

func TestDifferenceFunc(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("difference between arrays", prop.ForAll(func(input1, input2 []int) bool {
		slices.SortFunc(input1, cmp.Compare[int])
		slices.SortFunc(input2, cmp.Compare[int])
		result := sortedslices.DifferenceFunc(cmp.Compare[int], input1, input2)
		if !slices.IsSortedFunc(result, cmp.Compare[int]) {
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

func TestIntersectionFunc(t *testing.T) {
	numberGenerator := gen.Int()
	numSliceGenerator := gen.SliceOfN(20, numberGenerator)

	properties := gopter.NewProperties(nil)

	properties.Property("union item of arrays", prop.ForAll(func(input1, input2, input3 []int) bool {
		slices.SortFunc(input1, cmp.Compare[int])
		slices.SortFunc(input2, cmp.Compare[int])
		slices.SortFunc(input3, cmp.Compare[int])
		result := sortedslices.IntersectionFunc(cmp.Compare[int], input1, input2, input3)
		if !slices.IsSortedFunc(result, cmp.Compare[int]) {
			return false
		}
		for _, e := range result {
			if _, found := slices.BinarySearchFunc(input1, e, cmp.Compare[int]); !found {
				return false
			}
			if _, found := slices.BinarySearchFunc(input2, e, cmp.Compare[int]); !found {
				return false
			}
			if _, found := slices.BinarySearchFunc(input3, e, cmp.Compare[int]); !found {
				return false
			}
		}
		return true
	}, numSliceGenerator, numSliceGenerator, numSliceGenerator))

	properties.TestingRun(t)
}
