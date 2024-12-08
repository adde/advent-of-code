package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[T]) ToSlice() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

// Union returns a new set with elements from both sets
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := New[T]()
	for item := range s {
		result.Add(item)
	}
	for item := range other {
		result.Add(item)
	}
	return result
}

// Intersection returns a new set with common elements
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := New[T]()
	for item := range s {
		if other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// Difference returns a new set with elements in s but not in other
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := New[T]()
	for item := range s {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

// String provides a string representation of the set
func (s Set[T]) String() string {
	var items []string
	for item := range s {
		items = append(items, fmt.Sprintf("%v", item))
	}
	return "{" + strings.Join(items, ", ") + "}"
}
