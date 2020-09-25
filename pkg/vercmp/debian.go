package vercmp

import (
	"regexp"
	"strings"
	"unicode"
)

func debianLexicalCompare(a string, b string) int {
	aRunes := []rune(a)
	bRunes := []rune(b)
	var i int
	for i = 0; i < len(aRunes) && i < len(bRunes); i++ {
		if aRunes[i] == '~' && bRunes[i] != '~' {
			return -1
		} else if aRunes[i] != '~' && bRunes[i] == '~' {
			return 1
		} else if unicode.IsLetter(aRunes[i]) && !unicode.IsLetter(bRunes[i]) {
			return -1
		} else if !unicode.IsLetter(aRunes[i]) && unicode.IsLetter(bRunes[i]) {
			return 1
		} else {
			difference := aRunes[i] - bRunes[i]
			if difference != 0 {
				return int(difference)
			}
		}
	}
	if len(aRunes) > len(bRunes) {
		if aRunes[i] == '~' {
			return -1
		} else {
			return 1
		}
	} else if len(bRunes) > len(aRunes) {
		if bRunes[i] == '~' {
			return 1
		} else {
			return -1
		}
	} else {
		return 0
	}
	return strings.Compare(a, b)
}

var findEpoch *regexp.Regexp = regexp.MustCompile("^(\\p{Nd}+):")

func debianEpoch(a string) (string, string) {
	aMatches := findEpoch.FindStringSubmatch(a)
	if len(aMatches) == 0 {
		return "0", a
	} else {
		aRest := findEpoch.ReplaceAllString(a, "")
		return aMatches[1], aRest
	}
}

var findDigits *regexp.Regexp = regexp.MustCompile("^(\\p{Nd}*)")
var findNonDigits *regexp.Regexp = regexp.MustCompile("^(\\P{Nd}*)")

// Compare two strings as if they were debian package version numbers, as
// outlined in the Debian Policy Manual:
// https://www.debian.org/doc/debian-policy/ch-controlfields.html#s-f-Version
// Return an integer less than 0 if version `a` is older than version `b`, a
// positive integer if version `a` compares as newer than version `b`, and 0
// otherwise.
//
// Epoch numbers, upstream versions, and debian revision version parts are
// fully supported.
//
// With the exceptions of `NaiveCompare` and `PythonCompare`, all other vercmp
// algorithms in vercmp are completely implemented in terms of this function.
func DebianCompare(a string, b string) int {
	if a == b {
		return 0
	}
	aEpoch, a := debianEpoch(a)
	bEpoch, b := debianEpoch(b)
	earlyCmp := strIntCompare(aEpoch, bEpoch)
	if earlyCmp != 0 {
		return earlyCmp
	}

	for len(a) > 0 || len(b) > 0 {
		aNonDigits := findNonDigits.FindStringSubmatch(a)
		bNonDigits := findNonDigits.FindStringSubmatch(b)
		nonDigitsResult := debianLexicalCompare(aNonDigits[1], bNonDigits[1])
		if nonDigitsResult != 0 {
			return nonDigitsResult
		}
		a = findNonDigits.ReplaceAllString(a, "")
		b = findNonDigits.ReplaceAllString(b, "")

		aDigits := findDigits.FindStringSubmatch(a)
		bDigits := findDigits.FindStringSubmatch(b)
		digitResult := strIntCompare(aDigits[1], bDigits[1])
		if digitResult != 0 {
			return digitResult
		}
		a = findDigits.ReplaceAllString(a, "")
		b = findDigits.ReplaceAllString(b, "")
	}
	return 0
}
