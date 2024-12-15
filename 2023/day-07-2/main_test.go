package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 5905 {
		t.Errorf("Expected 5905, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 253253225 {
		t.Errorf("Expected 253253225, but got %d", res)
	}
}

func TestGetHandType(t *testing.T) {
	data := []Hand{
		{
			value:    "AAAAA",
			handType: FiveOfAKind,
		},
		{
			value:    "AAAA1",
			handType: FourOfAKind,
		},
		{
			value:    "AAAAJ",
			handType: FiveOfAKind,
		},
		{
			value:    "AAA11",
			handType: FullHouse,
		},
		{
			value:    "AAA12",
			handType: ThreeOfAKind,
		},
		{
			value:    "AA122",
			handType: TwoPairs,
		},
		{
			value:    "AA123",
			handType: Pair,
		},
		{
			value:    "KK677",
			handType: TwoPairs,
		},
		{
			value:    "AKQJT",
			handType: Pair,
		},
		{
			value:    "AKQ9T",
			handType: HighCard,
		},
	}

	for _, d := range data {
		_, res := getHandType(d.value)
		if res != d.handType {
			t.Errorf("Expected hand type for %s to be '%d', but got '%d'", d.value, d.handType, res)
		}
	}
}

func TestCompareHands(t *testing.T) {
	handA := Hand{
		value:    "AAAA1",
		handType: FourOfAKind,
	}
	handB := Hand{
		value:    "AAAAA",
		handType: FiveOfAKind,
	}

	res := compareHands(handA, handB)

	if res != true {
		t.Errorf("Should be true")
	}

	handA = Hand{
		value:    "BBBBB",
		handType: FourOfAKind,
	}
	handB = Hand{
		value:    "AAAAA",
		handType: FiveOfAKind,
	}

	res = compareHands(handA, handB)

	if res != true {
		t.Errorf("Should be true")
	}
}
