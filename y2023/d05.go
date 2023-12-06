package y2023

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type mapID int

const (
	SeedToSoil mapID = iota
	SoilToFertilizer
	FertilizerToWater
	WaterToLight
	LightToTemperature
	TemperatureToHumidity
	HumidityToLocation
)

type maprange struct {
	dst int
	src int
	len int
}

type mapping []maprange

func (m mapping) from(n int) int {
	for _, mr := range m {
		if n >= mr.src && n < (mr.src+mr.len) {
			offset := n - mr.src
			return mr.dst + offset
		}
	}

	return n
}

func (m mapping) to(n int) int {
	for _, mr := range m {
		if n >= mr.dst && n < (mr.dst+mr.len) {
			offset := n - mr.dst
			return mr.src + offset
		}
	}

	return n
}

type almanac struct {
	seeds []int
	maps  map[mapID]mapping
}

func (alm *almanac) location(seed int) int {
	soil := alm.maps[SeedToSoil].from(seed)
	fert := alm.maps[SoilToFertilizer].from(soil)
	watr := alm.maps[FertilizerToWater].from(fert)
	lite := alm.maps[WaterToLight].from(watr)
	temp := alm.maps[LightToTemperature].from(lite)
	hmdt := alm.maps[TemperatureToHumidity].from(temp)
	return alm.maps[HumidityToLocation].from(hmdt)
}

func (alm *almanac) reverse(lctn int) int {
	hmdt := alm.maps[HumidityToLocation].to(lctn)
	temp := alm.maps[TemperatureToHumidity].to(hmdt)
	lite := alm.maps[LightToTemperature].to(temp)
	watr := alm.maps[WaterToLight].to(lite)
	fert := alm.maps[FertilizerToWater].to(watr)
	soil := alm.maps[SoilToFertilizer].to(fert)
	return alm.maps[SeedToSoil].to(soil)
}

func parseAlmanac(r io.Reader) (alm almanac) {
	alm.maps = make(map[mapID]mapping)

	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return
	}

	fs := strings.Fields(scanner.Text())
	for i := 1; i < len(fs); i++ {
		val, _ := strconv.Atoi(fs[i])
		alm.seeds = append(alm.seeds, val)
	}

	scanner.Scan()
	scanner.Scan()

	for id := SeedToSoil; id <= HumidityToLocation; id++ {
		if !scanner.Scan() {
			return
		}

		for {
			if scanner.Text() == "" {
				scanner.Scan()
				break
			}

			fs := strings.Fields(scanner.Text())
			dst, _ := strconv.Atoi(fs[0])
			src, _ := strconv.Atoi(fs[1])
			len, _ := strconv.Atoi(fs[2])
			maprange := maprange{
				dst: dst,
				src: src,
				len: len,
			}

			alm.maps[id] = append(alm.maps[id], maprange)

			if !scanner.Scan() {
				return
			}
		}
	}

	return
}

func SolveD05P01(r io.Reader) (string, error) {
	alm := parseAlmanac(r)

	var min int
	for i, seed := range alm.seeds {
		if i == 0 {
			min = alm.location(seed)
			continue
		}

		loc := alm.location(seed)
		if loc < min {
			min = loc
		}
	}

	return strconv.Itoa(min), nil
}

func SolveD05P02(r io.Reader) (string, error) {
	alm := parseAlmanac(r)

	var min int
	for {
		seed := alm.reverse(min)
		for i := 0; i < len(alm.seeds); i += 2 {
			seedStart := alm.seeds[i]
			seedEnd := alm.seeds[i] + alm.seeds[i+1]
			if seed >= seedStart && seed < seedEnd {
				return strconv.Itoa(min), nil
			}
		}
		min += 1
	}
}
