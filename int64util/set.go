package int64util

import "sync"

type Set struct {
	sync.Once

	Set map[int64]struct{}
}

func (s *Set) Add(vs ...int64) {
	if len(vs) == 0 {
		return
	}

	s.Do(func() {
		if s.Set == nil {
			s.Set = make(map[int64]struct{}, len(vs))
		}
	})

	for _, v := range vs {
		s.Set[v] = struct{}{}
	}
}

func (s *Set) Int64s() []int64 {
	if len(s.Set) == 0 {
		return nil
	}

	res := make([]int64, 0, len(s.Set))

	for v := range s.Set {
		res = append(res, v)
	}

	return res
}

func (s *Set) Has(v int64) bool {
	if len(s.Set) == 0 {
		return false
	}

	_, ok := s.Set[v]
	return ok
}

func (s *Set) Delete(vv ...int64) {
	for _, v := range vv {
		delete(s.Set, v)
	}
}
