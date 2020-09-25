package vercmp

import (
	"regexp"
	"strings"
)

var naiveFindDigits *regexp.Regexp = regexp.MustCompile("^\\p{Nd}+$")

func naivePartCompare(a string, b string) int {
	aResults := naiveFindDigits.FindStringSubmatch(a)
	bResults := naiveFindDigits.FindStringSubmatch(b)
	if aResults != nil && bResults != nil {
		return strIntCompare(a, b)
	} else if aResults != nil {
		return 1
	} else if bResults != nil {
		return -1
	} else {
		return strings.Compare(a, b)
	}
}

var naiveFindPunct *regexp.Regexp = regexp.MustCompile("\\p{P}+")

// Compare two version numbers, separated into parts by punctuation. Compare
// each part in turn. If both parts are numeric, comare them as numbers. If one
// part is numeric, but the other alphanumeric, the first part is newer. if
// both are alphanumeric, compare lexically. Continue until one of the parts is
// newer than the other, and return a -1, 0, or 1 in the usual manner to
// indicate if the second argument is newer, if they are both equal, or if the
// first argument is newer, respectively. This is, in fact, the old rpmvercmp
// algorithm. It is not used in the modern version of rpm; however, it is still
// used in comparing the local part of the python version comparison algorithm
// (`vercmp.PythonCompare`).
func NaiveCompare(a string, b string) int {
	aParts := naiveFindPunct.Split(a, -1)
	bParts := naiveFindPunct.Split(b, -1)
	aLen := len(aParts)
	bLen := len(bParts)
	var minLength int
	if aLen < bLen {
		minLength = aLen
	} else {
		minLength = bLen
	}
	var partCompare int
	for i := 0; i < minLength; i++ {
		partCompare = naivePartCompare(aParts[i], bParts[i])
		if partCompare != 0 {
			return partCompare
		}
	}
	return aLen - bLen
}
