package y2023

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type scratchcard struct {
	id      int
	winning map[int]struct{}
	have    map[int]struct{}
}

func (sc *scratchcard) matches() []int {
	var matches []int

	for num := range sc.have {
		if _, ok := sc.winning[num]; ok {
			matches = append(matches, num)
		}
	}

	return matches
}

func parseScratchcard(s string) (sc scratchcard) {
	header, body, _ := strings.Cut(s, ": ")
	cardId := strings.Fields(header)[1]
	sc.id, _ = strconv.Atoi(cardId)

	sc.winning = make(map[int]struct{})
	sc.have = make(map[int]struct{})
	winning, have, _ := strings.Cut(body, " | ")

	for _, numText := range strings.Fields(winning) {
		num, _ := strconv.Atoi(numText)
		sc.winning[num] = struct{}{}
	}

	for _, numText := range strings.Fields(have) {
		num, _ := strconv.Atoi(numText)
		sc.have[num] = struct{}{}
	}

	return
}

func SolveD04P01(r io.Reader) (string, error) {
	var sum int
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		sc := parseScratchcard(scanner.Text())
		matches := sc.matches()
		if len(matches) > 0 {
			sum += (1 << (len(matches) - 1))
		}
	}

	return strconv.Itoa(sum), nil
}

func SolveD04P02(r io.Reader) (string, error) {
	var sum int
	var cards []int
	matchmap := make(map[int][]int)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		sc := parseScratchcard(scanner.Text())
		matchmap[sc.id] = sc.matches()
		cards = append(cards, sc.id)
	}

	for {
		if len(cards) == 0 {
			break
		}

		sum += 1
		cardnum := cards[0]
		cards = cards[1:]

		matches := matchmap[cardnum]
		for i := 1; i <= len(matches); i++ {
			cards = append(cards, i+cardnum)
		}
	}

	return strconv.Itoa(sum), nil
}
