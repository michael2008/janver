package vercmp

import (
	"regexp"
)

var findMetadata *regexp.Regexp = regexp.MustCompile("\\+.*$")
var findDashes *regexp.Regexp = regexp.MustCompile("-")

// Convert a SemVer version number to a debian version number such that the
// function call
//     vercmp.DebianCompare(vercmp.SemverNormalize(a),
//                          vercmp.SemverNormalize(b))
// would compare two correct SemVer version numbers correctly.
//
// Reference: http://semver.org/
func SemverNormalize(a string) string {
	a = findMetadata.ReplaceAllString(a, "")
	a = findDashes.ReplaceAllString(a, "~")
	return a
}

// Compare two SemVer version numbers. Simply calls `vercmp.SemverNormalize` on
// each version and then feeds the results to `vercmp.DebianCompare`.
//
// Reference: http://semver.org/
func SemverCompare(a string, b string) int {
	return DebianCompare(SemverNormalize(a), SemverNormalize(b))
}
