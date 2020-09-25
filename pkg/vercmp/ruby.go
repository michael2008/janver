package vercmp

import (
	"regexp"
)

var findDigitAlpha *regexp.Regexp = regexp.MustCompile(
	"(\\p{Nd}+)(\\p{L}+)")
var findAlphaDigit *regexp.Regexp = regexp.MustCompile(
	"(\\p{L}+)(\\p{Nd}+)")
var findAlphaSuffix *regexp.Regexp = regexp.MustCompile(
	"(\\p{Nd}+(\\.\\p{Nd}+)*)\\.(\\p{L}.*)$")

// Convert a Ruby Gem ( https://rubygems.org ) version number to a debian
// version number such that the function call
//     vercmp.DebianCompare(vercmp.RubyNormalize(a),
//                          vercmp.RubyNormalize(b))
// would compare two correct ruby gem version numbers
// correctly.
//
// Reference: http://ruby-doc.org/stdlib-2.0.0/libdoc/rubygems/rdoc/Gem/Version.html
func RubyNormalize(a string) string {
	a = findDigitAlpha.ReplaceAllString(a, "$1.$2")
	a = findAlphaDigit.ReplaceAllString(a, "$1.$2")
	a = findAlphaSuffix.ReplaceAllString(a, "$1~$3")
	return a
}

// Compare two ruby gem version numbers. Simply calls `vercmp.RubyNormalize`
// on each version and then feeds the results to `vercmp.DebianCompare`.
//
// Reference: http://ruby-doc.org/stdlib-2.0.0/libdoc/rubygems/rdoc/Gem/Version.html
func RubyCompare(a string, b string) int {
	return DebianCompare(RubyNormalize(a), RubyNormalize(b))
}
