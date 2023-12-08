package adventofcode

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

type Strength int

const (
	_ Strength = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	AS
)

type HandType int

const (
	_ HandType = iota
	HIGHCARD
	ONEPAIR
	TWOPAIR
	THREEKIND
	FULLHOUSE
	FOURKIND
	FIVEKIND
)

var strengths = map[string]Strength{
	"A": AS,
	"K": KING,
	"Q": QUEEN,
	"J": JACK,
	"T": TEN,
	"9": NINE,
	"8": EIGHT,
	"7": SEVEN,
	"6": SIX,
	"5": FIVE,
	"4": FOUR,
	"3": THREE,
	"2": TWO,
}

type Hand struct {
	cards []string
	bid   int
}

func newHand(line string) *Hand {
	chuncks := strings.Fields(line)
	cardString, bid := chuncks[0], chuncks[1]
	return &Hand{strings.Split(cardString, ""), parseInt(bid)}
}

func (h *Hand) _type() HandType {
	counts := map[string]int{}
	for _, card := range h.cards {
		if count, ok := counts[card]; ok {
			counts[card] = count + 1
		} else {
			counts[card] = 1
		}
	}

	switch len(counts) {
	case 1:
		return FIVEKIND
	case 2:
		for _, v := range counts {
			if v == 4 {
				return FOURKIND
			}
		}
		return FULLHOUSE
	case 3:
		pairs := 0
		for _, v := range counts {
			if v == 2 {
				pairs++
			}
		}
		if pairs == 2 {
			return TWOPAIR
		}
		return THREEKIND
	case 4:
		return ONEPAIR
	default:
		return HIGHCARD
	}
}

func (h *Hand) less(o *Hand) bool {
	hType, oType := h._type(), o._type()

	if hType < oType {
		return true
	} else if hType > oType {
		return false
	}

	return h.compareByStrengths(o)
}

func (h *Hand) compareByStrengths(o *Hand) bool {
	var i, hStrength, oStrength int
	for i = 0; i < 5; i++ {
		hStrength = int(strengths[h.cards[i]])
		oStrength = int(strengths[o.cards[i]])

		if hStrength != oStrength {
			break
		}
	}
	return hStrength < oStrength
}

func day7(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	hands := []Hand{}

	for scanner.Scan() {
		line := scanner.Text()
		hands = append(hands, *newHand(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].less(&hands[j])
	})

	sum := 0

	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	return fmt.Sprintf("%v", sum)
}

func day7Ext(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += parseNumberFromLine(line, false)
	}

	return fmt.Sprintf("%v", sum)
}
