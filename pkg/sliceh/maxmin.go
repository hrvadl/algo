package sliceh

import "slices"

type Orderable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float64 | ~float32
}

func MinFor2D[T Orderable](s [][]T) []T {
	mins := make([]T, len(s))
	for i, el := range s {
		mins[i] = slices.Min(el)
	}
	return mins
}

func MaxFor2D[T Orderable](s [][]T) []T {
	maxs := make([]T, len(s))
	for i, el := range s {
		maxs[i] = slices.Max(el)
	}
	return maxs
}

func MaxIdxs[T Orderable](s []T) []int {
	if len(s) == 0 {
		return nil
	}

	maxel := s[0]
	maxIndexes := make([]int, 1)
	for i := 1; i < len(s); i++ {
		el := s[i]
		if el > maxel {
			maxIndexes = []int{i}
			maxel = el
		} else if el == maxel {
			maxIndexes = append(maxIndexes, i)
		}
	}

	return maxIndexes
}

func MinIdxs[T Orderable](s []T) []int {
	if len(s) == 0 {
		return nil
	}

	minel := s[0]
	minIndexes := make([]int, 1)
	for i := 1; i < len(s); i++ {
		el := s[i]
		if el < minel {
			minIndexes = []int{i}
			minel = el
		} else if el == minel {
			minIndexes = append(minIndexes, i)
		}
	}

	return minIndexes
}
