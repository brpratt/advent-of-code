package y2023

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"
)

type handType int

const (
	fiveOfAKind handType = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	twoPair
	onePair
	highCard
)

type card rune

func (c card) value() int {
	vm := map[card]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
	}

	if val, ok := vm[c]; ok {
		return val
	}

	return int(c - '0')
}

type hand string

func (h hand) nonJ() (other []rune) {
	for _, c := range h {
		if c != 'J' {
			other = append(other, c)
		}
	}

	return
}

func (h hand) t() handType {
	counts := make(map[card]int)

	for _, c := range h {
		counts[card(c)]++
	}

	n3s := 0
	n2s := 0

	for _, count := range counts {
		switch count {
		case 5:
			return fiveOfAKind
		case 4:
			return fourOfAKind
		case 3:
			n3s++
		case 2:
			n2s++
		}
	}

	switch {
	case n3s == 1 && n2s == 1:
		return fullHouse
	case n3s == 1:
		return threeOfAKind
	case n2s == 2:
		return twoPair
	case n2s == 1:
		return onePair
	}

	return highCard
}

func (h hand) t2() handType {
	other := h.nonJ()
	if len(other) == 5 {
		return h.t()
	}
	if len(other) == 0 {
		return fiveOfAKind
	}

	maxType := highCard
	for _, c := range other {
		mapFunc := func(r rune) rune {
			if r == 'J' {
				return c
			}
			return r
		}

		newH := hand(strings.Map(mapFunc, string(h)))
		if newType := newH.t(); newType < maxType {
			maxType = newType
		}
	}

	return maxType
}

type handRound struct {
	h   hand
	bid int
}

func parseHandRound(s string) (hr handRound) {
	fs := strings.Fields(s)
	hr.h = hand(fs[0])
	hr.bid, _ = strconv.Atoi(fs[1])
	return
}

type byHand []handRound

func (a byHand) Len() int      { return len(a) }
func (a byHand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byHand) Less(i, j int) bool {
	iType := a[i].h.t()
	jType := a[j].h.t()

	if iType < jType {
		return false
	}
	if iType > jType {
		return true
	}

	cardsI := []card(a[i].h)
	cardsJ := []card(a[j].h)

	for i := range cardsI {
		cardI := cardsI[i].value()
		cardJ := cardsJ[i].value()

		if cardI < cardJ {
			return true
		}
		if cardI > cardJ {
			return false
		}
	}

	return false
}

type byHand2 []handRound

func (a byHand2) Len() int      { return len(a) }
func (a byHand2) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byHand2) Less(i, j int) bool {
	iType := a[i].h.t2()
	jType := a[j].h.t2()

	if iType < jType {
		return false
	}
	if iType > jType {
		return true
	}

	cardsI := []card(a[i].h)
	cardsJ := []card(a[j].h)

	for i := range cardsI {
		cardI := cardsI[i].value()
		cardJ := cardsJ[i].value()
		
		if cardsI[i] == 'J' {
			cardI = 0
		}

		if cardsJ[i] == 'J' {
			cardJ = 0
		}

		if cardI < cardJ {
			return true
		}
		if cardI > cardJ {
			return false
		}
	}

	return false
}

func SolveD07P01(r io.Reader) (string, error) {
	var rounds []handRound
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		rounds = append(rounds, parseHandRound(scanner.Text()))
	}

	sort.Sort(byHand(rounds))

	winnings := 0
	for i, hand := range rounds {
		winnings += hand.bid * (i + 1)
	}

	return strconv.Itoa(winnings), nil
}

func SolveD07P02(r io.Reader) (string, error) {
	var rounds []handRound
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		rounds = append(rounds, parseHandRound(scanner.Text()))
	}

	sort.Sort(byHand2(rounds))

	winnings := 0
	for i, hand := range rounds {
		winnings += hand.bid * (i + 1)
	}

	return strconv.Itoa(winnings), nil
}
