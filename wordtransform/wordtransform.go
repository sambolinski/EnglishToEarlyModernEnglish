package wordtransform

import (
	"strings"
)

type Bitmask int

const (
	REPLACE_S_WITH_LONG_S = 0x1
	REPLACE_W_WITH_V      = 0x2 //uppercase w
	REPLACE_W_WITH_U      = 0x2 //lowercase w
)
const (
	CHECK_BEFORE = 0x1
	CHECK_AFTER  = 0x2
)

func transformWord(word string, mask Bitmask) string {
	if (mask & REPLACE_S_WITH_LONG_S) != 0 {
		word = replaceWithLongS(word)
	}
	if (mask & REPLACE_W_WITH_V) != 0 {
		word = replaceWithLongS(word)
	}
	if (mask & REPLACE_W_WITH_U) != 0 {
		word = replaceWithLongS(word)
	}
	return word
}

func replaceWithLongS(word string) string {

	allS := findAllIndex(word, "s")

	for key, idx := range allS {
		//Check if final s
		if idx == len(word)-1 {
			delete(allS, key)
		}

		//Check if before apostraphe
		if checkBeforeAndAfter(word, idx, byte('\''), CHECK_AFTER) {
			delete(allS, key)
		}

		//Check if before or after b
		if checkBeforeAndAfter(word, idx, byte('b'), CHECK_BEFORE|CHECK_AFTER) {
			delete(allS, key)
		}

		//Check if before or after k
		if checkBeforeAndAfter(word, idx, byte('k'), CHECK_BEFORE|CHECK_AFTER) {
			delete(allS, key)
		}

		//Check if before is s (second s of souble s will be lower): TODO add mask to make this either ſs or ſſ
		if checkBeforeAndAfter(word, idx, byte('s'), CHECK_BEFORE) {
			delete(allS, key)
		}

		//Check if before and after f
		if checkBeforeAndAfter(word, idx, byte('f'), CHECK_BEFORE|CHECK_AFTER) {
			delete(allS, key)
		}
	}

	mutable := []rune(word)
	for _, val := range allS {
		mutable[val] = rune('ſ')
	}

	return string(mutable)
}

func checkBeforeAndAfter(word string, startIdx int, val byte, mask Bitmask) bool {
	before := false
	after := false

	if (mask & CHECK_BEFORE) != 0 {
		if startIdx-1 >= 0 {
			if word[startIdx-1] == val {
				before = true
			}
		}

	}

	if (mask & CHECK_AFTER) != 0 {
		if startIdx+1 < len(word) {
			if word[startIdx+1] == val {
				after = true
			}
		}
	}

	return before || after

}
func replaceWithV(word string) string {
	return strings.ReplaceAll(word, "W", "VV")
}
func replaceWithU(word string) string {
	return strings.ReplaceAll(word, "w", "uu")
}

//Returns all indices of substring
//TODO:improve algorithm to find substring instead of letter
func findAllIndex(word, letterToFind string) (indices map[int]int) {
	indices = map[int]int{}
	for idx, letter := range word {
		if string(letter) == letterToFind {
			indices[idx] = idx
		}
	}
	return
}
