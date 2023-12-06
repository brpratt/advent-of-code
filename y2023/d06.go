package y2023

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

/*
distance = held * (time - held)
distance = held * time - held^2

distance = -held^2 + time*held
0 = -held^2 + time*held - distance

       -time +- sqrt(time^2 - 4*distance)
held = ----------------------------------
	               -2
*/

func held(time, distance int) (fst float64, snd float64) {
	fst = (float64(time) - math.Sqrt(float64(time*time-4*distance))) / 2
	snd = (float64(time) + math.Sqrt(float64(time*time-4*distance))) / 2
	return
}

func nwin(time, distance int) int {
	fst, snd := held(time, distance)

	winFst := int(math.Ceil(fst))
	if fst == float64(int(fst)) {
		winFst++
	}

	winSnd := int(math.Floor(snd))
	if snd == float64(int(snd)) {
		winSnd--
	}

	return (winSnd - winFst) + 1
}

func parseRaceDoc(r io.Reader) (times []int, distances []int) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return
	}
	fs := strings.Fields(scanner.Text())
	for i := 1; i < len(fs); i++ {
		time, _ := strconv.Atoi(fs[i])
		times = append(times, time)
	}

	if !scanner.Scan() {
		return
	}

	fs = strings.Fields(scanner.Text())
	for i := 1; i < len(fs); i++ {
		distance, _ := strconv.Atoi(fs[i])
		distances = append(distances, distance)
	}

	return
}

func SolveD06P01(r io.Reader) (string, error) {
	times, distances := parseRaceDoc(r)
	count := 1

	for i := range times {
		nwin := nwin(times[i], distances[i])
		count *= nwin
	}

	return strconv.Itoa(count), nil
}

func SolveD06P02(r io.Reader) (string, error) {
	times, distances := parseRaceDoc(r)

	timeStr := ""
	distanceStr := ""
	for i := range times {
		timeStr += strconv.Itoa(times[i])
		distanceStr += strconv.Itoa(distances[i])
	}

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)

	return strconv.Itoa(nwin(time, distance)), nil
}
