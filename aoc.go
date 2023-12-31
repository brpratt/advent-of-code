package aoc

import (
	"errors"
	"io"

	"github.com/brpratt/advent-of-code/y2019"
	"github.com/brpratt/advent-of-code/y2023"
)

type Solver interface {
	Solve(r io.Reader) (string, error)
}

type SolverFunc func(r io.Reader) (string, error)

func (s SolverFunc) Solve(r io.Reader) (string, error) {
	return s(r)
}

var _ Solver = SolverFunc(NotImplemented)

var (
	ErrNotImplemented = errors.New("not implemented")
)

func NotImplemented(r io.Reader) (string, error) {
	return "", ErrNotImplemented
}

func Get(year int, day int, part int) Solver {
	days, ok := years[year]
	if !ok {
		return SolverFunc(NotImplemented)
	}

	if day <= 0 || day > len(days) {
		return SolverFunc(NotImplemented)
	}

	parts := days[day-1]
	switch part {
	case 1:
		return parts.part01
	case 2:
		return parts.part02
	default:
		return SolverFunc(NotImplemented)
	}
}

type day struct {
	part01 SolverFunc
	part02 SolverFunc
}

type year []day

var years = map[int]year{
	2019: []day{
		{y2019.SolveD01P01, y2019.SolveD01P02},
		{y2019.SolveD02P01, y2019.SolveD02P02},
	},
	2023: []day{
		{y2023.SolveD01P01, y2023.SolveD01P02},
		{y2023.SolveD02P01, y2023.SolveD02P02},
		{y2023.SolveD03P01, y2023.SolveD03P02},
		{y2023.SolveD04P01, y2023.SolveD04P02},
		{y2023.SolveD05P01, y2023.SolveD05P02},
		{y2023.SolveD06P01, y2023.SolveD06P02},
		{y2023.SolveD07P01, y2023.SolveD07P02},
		{y2023.SolveD08P01, y2023.SolveD08P02},
		{y2023.SolveD09P01, y2023.SolveD09P02},
		{y2023.SolveD10P01, y2023.SolveD10P02},
		{y2023.SolveD11P01, y2023.SolveD11P02},
		{y2023.SolveD12P01, y2023.SolveD12P02},
		{y2023.SolveD13P01, y2023.SolveD13P02},
		{y2023.SolveD14P01, y2023.SolveD14P02},
	},
}
