package adventofcode

import "testing"

func TestDay1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	expected := "142"
	result := day1(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}

}

func TestDay1Ext(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
zone
aaatwoaa`

	expected := "314"
	result := day1Ext(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}

}

func TestDay2(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	expected := "8"
	result := day2(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay2Ext(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	expected := "2286"
	result := day2Ext(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay3(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
..........
.....*1000`

	expected := "5361"
	// expected := "4461"
	result := day3(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay3Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay4(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	expected := "13"
	result := day4(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay4Ext(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	expected := "30"
	result := day4Ext(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay5(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	expected := "35"
	result := day5(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay5Ext(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	expected := "46"
	result := day5Ext(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay6(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200
`

	expected := "288"
	result := day6(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay6Ext(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200
`

	expected := "71503"
	result := day6Ext(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay7(t *testing.T) {
	// 	input := `2345A 1
	// Q2KJJ 13
	// Q2Q2Q 19
	// T3T3J 17
	// T3Q33 11
	// 2345J 3
	// J345A 2
	// 32T3K 5
	// T55J5 29
	// KK677 7
	// KTJJT 34
	// QQQJA 31
	// JJJJJ 37
	// JAAAA 43
	// AAAAJ 59
	// AAAAA 61
	// 2AAAA 23
	// 2JJJJ 53
	// JJJJ2 41` // 6592
	input := `2JJJJ 50
JJJJ2 10 ` // 70

	expected := "70"
	result := day7(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay7Ext(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay8(t *testing.T) {
	inputTests := []struct {
		input    string
		expected string
	}{
		{
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			expected: "2",
		},
		{
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			expected: "6",
		},
	}

	for _, tst := range inputTests {
		result := day8(tst.input)
		if result != tst.expected {
			t.Fatalf("Failed. expected=%s got=%s", tst.expected, result)
		}
	}
}

func TestDay8Ext(t *testing.T) {
	inputTests := []struct {
		input    string
		expected string
	}{
		{
			input: `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`,
			expected: "6",
		},
	}

	for _, tst := range inputTests {
		result := day8Ext(tst.input)
		if result != tst.expected {
			t.Errorf("Failed. expected=%s got=%s", tst.expected, result)
		}
	}
}

func TestDay9(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay9Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay10(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay10Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay11(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay11Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay12(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay12Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay13(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay13Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay14(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay14Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay15(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay15Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay16(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay16Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay17(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay17Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay18(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay18Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay19(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay19Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay20(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay20Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay21(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay21Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay22(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay22Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay23(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay23Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay24(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay24Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay25(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}

func TestDay25Ext(t *testing.T) {
	input := ``

	expected := ""
	result := defaultFunction(input)

	if result != expected {
		t.Fatalf("Failed. expected=%s got=%s", expected, result)
	}
}
