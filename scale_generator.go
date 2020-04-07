package scale2

import (
	"strings"
)

type scaleArray [12]string

var sharpScale = scaleArray{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatScale = scaleArray{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var sharpUsers = scaleArray{"G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}
var pointer = 0

// Scale returns the generated musical scale
func Scale(tonic string, interval string) []string {
	var finalScale []string
	scale := findScale(tonic)
	pointer = findIndex(scale, strings.Title(strings.ToLower(tonic)))

	if interval == "" {
		interval = strings.Repeat("m", len(sharpScale))
	}

	for _, step := range interval {
		finalScale = append(finalScale, scale[pointer])
		pointerIncrementer(string(step))
	}

	return finalScale
}

func findIndex(src scaleArray, element string) int {
	for i, e := range src {
		if e == element {
			return i
		}
	}
	return -1
}

func findScale(tonic string) scaleArray {
	if tonic == "C" || tonic == "a" || findIndex(sharpUsers, tonic) != -1 {
		return sharpScale
	}
	return flatScale
}

func pointerIncrementer(step string) {
	if step == "m" {
		pointer++
	} else if step == "M" {
		pointer += 2
	} else {
		pointer += 3
	}

	if pointer > 11 {
		pointer = pointer - len(sharpScale)
	}
}
