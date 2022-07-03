package wordtransform

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceWithLongS(t *testing.T) {
	tt := []struct {
		Input  string
		Output string
	}{
		{Input: "satisfaction", Output: "ſatisfaction"},
		{Input: "offset", Output: "offset"},
		{Input: "his", Output: "his"},
		{Input: "success", Output: "ſucceſs"},
		{Input: "possess", Output: "poſseſs"},
		{Input: "satisfaction", Output: "ſatisfaction"},
		{Input: "substantive", Output: "ſubstantive"},
	}

	for _, test := range tt {
		res := replaceWithLongS(test.Input)
		fmt.Println(test.Input, ", ", res)
		assert.Equal(t, test.Output, res)
	}
}

func TestReplaceWithV(t *testing.T) {
	tt := []struct {
		Input  string
		Output string
	}{
		{Input: "Witch", Output: "VVitch"},
		{Input: "Warlock", Output: "VVarlock"},
	}

	for _, test := range tt {
		res := replaceWithV(test.Input)
		assert.Equal(t, test.Output, res)
	}
}

func TestReplaceWithU(t *testing.T) {
	tt := []struct {
		Input  string
		Output string
	}{
		{Input: "Law", Output: "Lauu"},
	}

	for _, test := range tt {
		res := replaceWithU(test.Input)
		assert.Equal(t, test.Output, res)
	}
}

func TestFindAllIndex(t *testing.T) {
	tt := []struct {
		Input  string
		Input2 string
		Output []int
	}{
		{Input: "Law", Input2: "v", Output: nil},
		{Input: "Test", Input2: "T", Output: []int{0}},
		{Input: "Success", Input2: "c", Output: []int{2, 3}},
		{Input: "Success", Input2: "S", Output: []int{0}},
		{Input: "Success", Input2: "s", Output: []int{5, 6}},
	}

	for _, test := range tt {
		res := findAllIndex(test.Input, test.Input2)
		assert.Equal(t, test.Output, res)
	}
}

func TestCheckBeforeAndAfter(t *testing.T) {
	tt := []struct {
		Input  string
		Input2 int
		Input3 byte
		Input4 Bitmask
		Output bool
	}{
		{Input: "Law", Input2: 2, Input3: byte('a'), Input4: CHECK_BEFORE | CHECK_AFTER, Output: true},
		{Input: "Law", Input2: 2, Input3: byte('b'), Input4: CHECK_BEFORE | CHECK_AFTER, Output: false},
		{Input: "Law", Input2: 1, Input3: byte('L'), Input4: CHECK_BEFORE | CHECK_AFTER, Output: true},
		{Input: "Law", Input2: 1, Input3: byte('w'), Input4: CHECK_BEFORE | CHECK_AFTER, Output: true},
		{Input: "clos'd", Input2: 3, Input3: byte('\''), Input4: CHECK_AFTER, Output: true},
		{Input: "cloe's", Input2: 5, Input3: byte('\''), Input4: CHECK_AFTER, Output: false},
		{Input: "cloe's", Input2: 5, Input3: byte('\''), Input4: CHECK_BEFORE, Output: true},
		{Input: "cloe's", Input2: 5, Input3: byte('\''), Input4: CHECK_BEFORE | CHECK_AFTER, Output: true},
	}

	for _, test := range tt {
		res := checkBeforeAndAfter(test.Input, test.Input2, test.Input3, test.Input4)
		assert.Equal(t, test.Output, res)
	}
}
