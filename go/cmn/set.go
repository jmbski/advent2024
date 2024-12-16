/*
Copyright 2024 Joseph Bochinski

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

********************************************************************************

	Package: cmn
	Title: set
	Description: Code for implementing a mathematical Set object
	Author: Joseph Bochinski
	Date: 2024-12-10

********************************************************************************
*/
package cmn

import (
	"sort"
	"sync"
)

// Comparable interface defines types that support comparison operations.
type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string
}

// Set is a custom data structure to represent a mathematical set
type Set[T Comparable] struct {
	elements map[T]struct{}
	mu       sync.RWMutex
}

// NewSet creates and returns a new TypedSet
func NewSet[T Comparable](items ...T) *Set[T] {
	set := &Set[T]{elements: make(map[T]struct{})}
	for _, item := range items {
		set.Add(item)
	}
	return set
}

// Add inserts an element into the set
func (s *Set[T]) Add(element T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elements[element] = struct{}{}
}

// Remove deletes an element from the set
func (s *Set[T]) Remove(element T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.elements, element)
}

// Contains checks if an element is in the set
func (s *Set[T]) Contains(element T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.elements[element]
	return exists
}

// Len returns the number of elements in the set
func (s *Set[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.elements)
}

// Values returns all elements in the set as a slice
func (s *Set[T]) Values() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

func (s *Set[T]) SortedValues() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func (s *Set[T]) EquivalentTo(other *Set[T]) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Len() != other.Len() {
		return false
	}
	for _, element := range s.Values() {
		if !other.Contains(element) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := NewSet[T]()

	// Add all elements from the first set
	for elem := range s.elements {
		result.Add(elem)
	}

	// Add all elements from the second set
	for elem := range other.elements {
		result.Add(elem)
	}

	return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := NewSet[T]()

	// Iterate over the smaller set for efficiency
	if s.Len() < other.Len() {
		for elem := range s.elements {
			if other.Contains(elem) {
				result.Add(elem)
			}
		}
	} else {
		for elem := range other.elements {
			if s.Contains(elem) {
				result.Add(elem)
			}
		}
	}

	return result
}

func (s *Set[T]) Complement(other *Set[T]) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := NewSet[T]()

	// Add elements from the first set that are not in the second set
	for elem := range s.elements {
		if !other.Contains(elem) {
			result.Add(elem)
		}
	}

	return result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Complement(other)
}
