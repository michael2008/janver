package vercmp

import (
	"regexp"
)

var mavenNormalPunct = "[^\\p{L}\\p{N}~]+"
var mavenFindAlpha *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"([Aa][Ll][Pp][Hh][Aa]|[Aa])")
var mavenFindBeta *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"([Bb][Ee][Tt][Aa]|[Bb])")
var mavenFindMilestone *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"([Mm][Ii][Ll][Ee][Ss][Tt][Oo][Nn][Ee]|[Mm])")
var mavenFindRc *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"([Rr][Cc]|[Cc][Rr])")

var mavenFindSnapshot *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"([Ss][Nn][Aa][Pp][Ss][Hh][Oo][Tt])")

var mavenFindGaRelease *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"([Gg][Aa])")

var mavenFindFinalRelease *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"([Ff][Ii][Nn][Aa][Ll])")

var mavenFindStableRelease *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"([Ss][Tt][Aa][Bb][Ll][Ee])")

var mavenFindLetters *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"(\\p{L}+)")

var mavenFindZeroes *regexp.Regexp = regexp.MustCompile(mavenNormalPunct +
	"0+\\b")

// Convert a maven version number to a debian version number such that the
// function call
//     vercmp.DebianCompare(vercmp.MavenNormalize(a),
//                          vercmp.MavenNormalize(b))
// would compare two correct maven version numbers correctly.
//
// Reference: https://maven.apache.org/ref/3.0.3/maven-artifact/apidocs/org/apache/maven/artifact/versioning/ComparableVersion.html
func MavenNormalize(a string) string {
	a = mavenFindAlpha.ReplaceAllString(a, "~~alpha")
	a = mavenFindBeta.ReplaceAllString(a, "~~beta")
	a = mavenFindMilestone.ReplaceAllString(a, "~~milestone")
	a = mavenFindRc.ReplaceAllString(a, "~~rc")
	a = mavenFindSnapshot.ReplaceAllString(a, "~~snapshot")
	a = mavenFindGaRelease.ReplaceAllString(a, "")
	a = mavenFindFinalRelease.ReplaceAllString(a, "")
	a = mavenFindStableRelease.ReplaceAllString(a, "")
	a = mavenFindLetters.ReplaceAllString(a, "~$1")
	a = mavenFindZeroes.ReplaceAllString(a, "")
	return a

}

// Compare two maven version numbers. Simply calls `vercmp.MavenNormalize`
// on the versions and then feeds the results to `vercmp.DebianCompare`.
//
// Reference: https://maven.apache.org/ref/3.0.3/maven-artifact/apidocs/org/apache/maven/artifact/versioning/ComparableVersion.html
func MavenCompare(a string, b string) int {
	return DebianCompare(MavenNormalize(a), MavenNormalize(b))
}
