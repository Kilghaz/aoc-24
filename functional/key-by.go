package functional

import "github.com/samber/lo"

func KeyBy[T any, K comparable](input []T, iterator func(index int, item T) K) map[K][]T {
	result := make(map[K][]T)
	for index, item := range input {
		key := iterator(index, item)
		if !lo.HasKey(result, key) {
			result[key] = []T{item}
			continue
		}
		result[key] = append(result[key], item)
	}
	return result
}
