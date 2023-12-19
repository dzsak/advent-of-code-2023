package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.
// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.
func Test_calibrationValues(t *testing.T) {
	testInput := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	// The newly-improved calibration document consists of lines of text;
	// each line originally contained a specific calibration value that the Elves now need to recover.
	// On each line, the calibration value can be found by combining
	// the first digit and the last digit (in that order) to form a single two-digit number.
	res := calibrationValues(testInput)
	assert.Equal(t, 142, res)
}

// It looks like some of the digits are actually spelled out with letters:
// one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".
func Test_calibrationLetterValues(t *testing.T) {
	testInput := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	//In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76.
	// Adding these together produces 281.
	res := calibrationLetterValues(testInput)
	assert.Equal(t, 281, res)
}
