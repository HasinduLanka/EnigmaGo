package main

import (
	"regexp"
	"strings"
)

var RegexGo *regexp.Regexp = regexp.MustCompile("< *go  *(.*?)>(.*)< */go *>.*?")
var RegexVar *regexp.Regexp = regexp.MustCompile(`var *= *"(.*?)"`)
var RegexFunc *regexp.Regexp = regexp.MustCompile(`func *= *"(.*?)"`)

var NewLineRemover *strings.Replacer = strings.NewReplacer("\n", "")

func Render(B string) string {
	var O string

	RegexGo.ReplaceAllStringFunc(B, RenderTag)

	return O
}

func RenderTag(T string) string {
	MatchesVar := RegexVar.FindStringSubmatch(T)
	if len(MatchesVar) > 1 {
		return `<go var="` + MatchesVar[1] + `"></go> ` + RenderVar(MatchesVar[1])
	}

	MatchesFunc := RegexFunc.FindStringSubmatch(T)
	if len(MatchesFunc) > 1 {
		return `<go var="` + MatchesFunc[1] + `"></go> ` + RenderFunc(MatchesFunc[1])
	}

	return T
}

func RenderVar(T string) string {
	return `<a varname="` + T + `"> `
}

func RenderFunc(T string) string {
	return ""
}
