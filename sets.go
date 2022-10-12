package sets

import (
	"encoding/json"
)

type void struct{}

type Comparable interface {
	string | int | int64 | uint | uint64 | bool
}

type Set[T Comparable] map[T]void

func (s Set[T]) Add(v T) bool {
	if _, ok := s[v]; ok {
		return false
	}
	s[v] = void{}
	return true
}

func (s Set[T]) Remove(v T) bool {
	if _, ok := s[v]; !ok {
		return false
	}
	delete(s, v)
	return true
}

func (s Set[T]) Slice() []T {
	sl := make([]T, 0, len(s))
	for key := range s {
		sl = append(sl, key)
	}
	return sl
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

func (s Set[T]) UnmarshalJSON(v []byte) error {
	r := []T{}
	if err := json.Unmarshal(v, &r); err != nil {
		return err
	}

	for _, item := range r {
		s.Add(item)
	}

	return nil
}
