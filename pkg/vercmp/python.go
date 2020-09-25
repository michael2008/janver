package vercmp

import (
	"regexp"
	"strings"
)

var pythonFindEpoch *regexp.Regexp = regexp.MustCompile("^\\p{Nd}+!")
var pythonFindPost *regexp.Regexp = regexp.MustCompile("\\.post(\\p{Nd}+)")
var pythonFindAlpha *regexp.Regexp = regexp.MustCompile("(\\p{Nd}+)(\\p{P})?a(\\p{Nd}+)")
var pythonFindBeta *regexp.Regexp = regexp.MustCompile("(\\p{Nd}+)(\\p{P})?b(\\p{Nd}+)")
var pythonFindRc *regexp.Regexp = regexp.MustCompile("(\\p{Nd}+)(\\p{P})?rc(\\p{Nd}+)")
var pythonFindCrc *regexp.Regexp = regexp.MustCompile("(\\p{Nd}+)(\\p{P})?c(\\p{Nd}+)")
var pythonFindDev *regexp.Regexp = regexp.MustCompile("\\.dev(\\p{Nd}+)")

// Convert a python (pip) version number to a debian version number such that
// the function call
//     vercmp.DebianCompare(vercmp.PythonNormalize(a),
//                          vercmp.PythonNormalize(b))
// would compare two correct python version numbers correctly according to the
// PEP 440 standard ( https://www.python.org/dev/peps/pep-0440/ ), *for version
// numbers which do not contain python local version parts*. Version numbers
// which *do* contain local version parts are handled specially in
// PythonCompare.
func PythonNormalize(a string) string {
	a = strings.ToLower(a)
	a = pythonFindEpoch.ReplaceAllString(a, "$1:")
	a = pythonFindPost.ReplaceAllString(a, "!~$1")
	a = pythonFindAlpha.ReplaceAllString(a, "$1~a$3")
	a = pythonFindBeta.ReplaceAllString(a, "$1~b$3")
	a = pythonFindRc.ReplaceAllString(a, "$1~rc$3")
	a = pythonFindCrc.ReplaceAllString(a, "$1~rc$3")
	a = pythonFindDev.ReplaceAllString(a, "~~dev$1")
	return a
}

// Compare two python (pip) version numbers according to the PEP 440 standard
// ( https://www.python.org/dev/peps/pep-0440/ ).
//
// First, it calls `PythonNormalize` on each version’s public version identifier
// part and then returns the `DebianCompare` of the results, provided both
// version numbers do not contain a local version identifier.
//
// If a local version id is present on one of the version numbers, but both
// version numbers are otherwise identical, the version with a local version
// identifier is considered newer. If both have local version ids, but they are
// otherwise equal, the local version ids are compared using the `NaiveCompare`
// function.
func PythonCompare(a string, b string) int {
	aParts := strings.Split(a, "+")
	bParts := strings.Split(b, "+")
	debResult := DebianCompare(PythonNormalize(aParts[0]), PythonNormalize(bParts[0]))
	if debResult != 0 {
		return debResult
	}

	if len(aParts) > 1 && len(bParts) > 1 {
		return NaiveCompare(aParts[1], bParts[1])
	} else if len(aParts) > 1 {
		return 1
	} else if len(bParts) > 1 {
		return -1
	} else {
		return 0
	}
}
