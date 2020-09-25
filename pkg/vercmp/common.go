// Package vercmp provides many different string-based version comparison
// algorithms, intended to be used a la `strings.Sort`.
//
// The package is designed with the goal in mind of processing hundreds of
// packages at a time, and therefore calling a given package system's version
// comparison algorithm lots of times, such as to sort them or resolve their
// dependencies.
//
// Strings are taken to be version numbers. No parsing or marshalling is done;
// versions are *strings*. No need to marshal a version string when there are
// hundreds of such strings to be processed.
//
// The library does not attempt to describe parts of versions, such as major or
// minor portions. The library simply focuses on comparing versions.
//
// If the version string is compliant, the algorithm is correct.
//
// If the version number is not exactly compliant with the version spec, the
// algorithm tries its best to compare the strings well anyway. Packages do not
// always have perfectly compliant version numbers, so it is important to be
// able to handle some amount of noise when dealing with hundreds of packages.
package vercmp

import (
	"regexp"
	"strings"
)

var findZeroes *regexp.Regexp = regexp.MustCompile("^0+")

func strIntCompare(a string, b string) int {
	aNormalized := findZeroes.ReplaceAllString(a, "")
	bNormalized := findZeroes.ReplaceAllString(b, "")
	if len(aNormalized) == 0 {
		if len(bNormalized) == 0 {
			return 0
		} else {
			return -1
		}
	} else if len(bNormalized) == 0 {
		return 1
	} else if len(aNormalized) == len(bNormalized) {
		return strings.Compare(aNormalized, bNormalized)
	} else {
		return len(aNormalized) - len(bNormalized)
	}
}
