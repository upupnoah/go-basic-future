package slice

import "github.com/upupnoah/go-basic-future/go-basic-camp/ekit/internal/slice"

func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	res, err := slice.Add(src, element, index)
	return res, err
}
