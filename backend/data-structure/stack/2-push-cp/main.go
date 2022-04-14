package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	Top int
	Data []int
}

func NewStack(size int) Stack {
	return Stack {
		Top: -1,
		Data: make([]int, size),
	}
}

func (s *Stack) Push(Elemen int) error {
	if s.Top == len(s.Data) {
		return ErrStackOverflow
	}

	s.Top++
	s.Data = append(s.Data, Elemen)

	return nil
}
