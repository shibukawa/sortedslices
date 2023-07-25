# sortedslices

This package contains complementary functions for sorted slices.

## Functions

* `Insert(sortedSlice, element)`
* `InsertFunc(sortedSlice, element, cmp(i, j element) int)`

Insert element into sorted slice.

* `Remove(sortedSlice, element)`
* `RemoveFunc(sortedSlice, element, cmp(i, j element) int)`

Remove element from sorted slice.

* `IterateOver(callback(element, sliceIndex), sortedSlice...)`
* `IterateOverFunc(cmp(i, j element) int), callback(element, sliceIndex), sortedSlice...)`

Sort specified slices' element and pass to callback function.

For example, three slices([1, 3], [2, 5], [0, 4]) are passed, callback receives element in the following order

```
0, 1, 2, 3, 4, 5
```

* `Union(sortedSlice...)`
* `UnionFunc(cmp(i, j element) int), sortedSlice...)`

Generates union set from passed slices.

* `Difference(sortedSlice1, sortedSlice2)`
* `DifferenceFunc(cmp(i, j element) int), sortedSlice1, sortedSlice2)`

Generates difference set from passed slices.

* `Intersection(sortedSlice...)`
* `IntersectionFunc(cmp(i, j element) int), sortedSlice...)`

Generates intersection set from passed slices.

** `Min(sortedSlice)`

Returns minimum value in the slice.

** `Max(sortedSlice)`

Returns maximum value in the slice.

## Author

Yoshiki Shibukawa

## License

Apache-2