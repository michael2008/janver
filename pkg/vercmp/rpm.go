package vercmp

import (
	"regexp"
)

var rpmFindNormalPunct *regexp.Regexp = regexp.MustCompile("[^\\p{L}\\p{N}~]+")
var rpmFindAlphaDigit *regexp.Regexp = regexp.MustCompile("(\\p{L}+)\\.(\\p{Nd}+)")
var rpmFindDigitAlpha *regexp.Regexp = regexp.MustCompile("(\\p{Nd}+)\\.(\\p{L}+)")

// Convert an rpm version number to a debian version number such that the
// function call
//     vercmp.DebianCompare(vercmp.RpmNormalize(a),
//                          vercmp.RpmNormalize(b))
// would compare two correct rpm version numbers correctly.
//
// Reference
// implementation:
// https://github.com/rpm-software-management/rpm/blob/master/lib/rpmvercmp.c
func RpmNormalize(a string) string {
	a = rpmFindNormalPunct.ReplaceAllString(a, ".")
	a = rpmFindAlphaDigit.ReplaceAllString(a, "$1$2")
	a = rpmFindDigitAlpha.ReplaceAllString(a, "$1$2")
	return a
}

// Compare two rpm version numbers. Simply calls `vercmp.RpmNormalize` on each
// version and then feeds the results to `vercmp.DebianCompare`.
//
// Reference
// implementation:
// https://github.com/rpm-software-management/rpm/blob/master/lib/rpmvercmp.c
func RpmCompare(a string, b string) int {
	return DebianCompare(RpmNormalize(a), RpmNormalize(b))
}
