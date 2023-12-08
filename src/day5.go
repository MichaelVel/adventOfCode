package adventofcode

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

type Range struct {
	start, end int
}

func (r *Range) check(val int) bool {
	return r.start <= val && val < r.end
}

func (r *Range) offset(val int) int {
	return val - r.start
}

func (r *Range) fromOffset(val int) int {
	return val + r.start
}

type AlmanacMap struct {
	sources []Range
	dests   []Range
}

func newAlmanacMap(chunck string) *AlmanacMap {
	localMap := &AlmanacMap{
		sources: []Range{},
		dests:   []Range{},
	}

	scanner := bufio.NewScanner(strings.NewReader(chunck))
	scanner.Scan() // discard header
	for scanner.Scan() {
		data := Map(strings.Fields(scanner.Text()), parseInt)
		dest, source, length := data[0], data[1], data[2]
		localMap.sources = append(localMap.sources, Range{source, source + length})
		localMap.dests = append(localMap.dests, Range{dest, dest + length})
	}

	return localMap
}

func (am *AlmanacMap) get(val int) int {
	for i, sourceRange := range am.sources {
		if sourceRange.check(val) {
			offset := sourceRange.offset(val)
			return am.dests[i].fromOffset(offset)
		}
	}

	return val
}

func seeds(chunck string) []int {
	chunck = strings.Split(chunck, ":")[1]
	seedsArr := strings.Fields(chunck)
	return Map(seedsArr, parseInt)
}

func seedsRange(chunck string) []Range {
	seedRangeArr := []Range{}
	seedsArray := seeds(chunck)

	for i := 0; i < len(seedsArray)-1; i += 2 {
		start, length := seedsArray[i], seedsArray[i+1]
		seedRangeArr = append(seedRangeArr, Range{start, start + length})
	}

	return seedRangeArr
}

func day5(input string) string {
	chuncks := strings.Split(input, "\n\n")

	almanacSeeds := seeds(chuncks[0])
	almanacMaps := []AlmanacMap{}
	for _, chunck := range chuncks[1:] {
		almanacMaps = append(almanacMaps, *newAlmanacMap(chunck))
	}

	for _, m := range almanacMaps {
		almanacSeeds = Map(almanacSeeds, func(n int) int { return m.get(n) })
	}

	return fmt.Sprintf("%v", slices.Min(almanacSeeds))
}

func day5Ext(input string) string {
	chuncks := strings.Split(input, "\n\n")

	almanacSeeds := seedsRange(chuncks[0])
	almanacMaps := []AlmanacMap{}
	for _, chunck := range chuncks[1:] {
		almanacMaps = append(almanacMaps, *newAlmanacMap(chunck))
	}

	minLoc := -1
	fmt.Printf("%v\n", almanacSeeds)
	for _, r := range almanacSeeds {
		for seed := r.start; seed < r.end; seed++ {
			location := seed
			for _, m := range almanacMaps {
				location = m.get(location)
			}

			if minLoc == -1 {
				minLoc = location
			} else {
				minLoc = min(minLoc, location)
			}
		}

	}

	return fmt.Sprintf("%v", minLoc)
}
