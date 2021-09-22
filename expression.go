package repat

var _ Expression = CharacterCollection("")
var _ Expression = group{}

func MakeChar(c rune) Expression {
	return makeChar(c)
}

func makeChar(c rune) CharacterCollection {
	if ('a' <= c && c <= 'z') ||
		('A' <= c && c <= 'Z') ||
		('0' <= c && c <= '9') {
		return CharacterCollection(c)
	} else {
		return CharacterCollection(`\` + string(c))
	}
}

func MakeCharRange(from rune, to rune) CharacterCollection {
	cf := makeChar(from)
	ct := makeChar(to)
	return CharacterCollection(cf + "-" + ct)
}

// builtIn characters
const (
	Dot          CharacterCollection = `\.`
	Hyphen       CharacterCollection = `\-`
	Plus         CharacterCollection = `\+`
	QuestionMark CharacterCollection = `\?`
	Asterisk     CharacterCollection = `\*`
	Slash        CharacterCollection = `\\`
	BackSlash    CharacterCollection = `\/`
	Dollar       CharacterCollection = `\$`

	OpenSquareBracket  CharacterCollection = `\[`
	CloseSquareBracket CharacterCollection = `\]`
	OpenParenthesis    CharacterCollection = `\(`
	CloseParenthesis   CharacterCollection = `\)`
)

// character collection
type CharacterCollection string

func (cc CharacterCollection) isExpression() {}

func (cc CharacterCollection) ToString() string {

	if l := len(cc); l == 0 {
		return ""
	} else if l == 1 || (l == 2 && cc[0] == '\\') {
		return string(cc)
	}
	return "[" + string(cc) + "]"
}

// builtIn character collections
const (
	Any CharacterCollection = `.`

	UpperCase CharacterCollection = `A-Z`
	LowerCase CharacterCollection = `a-z`
	Letter    CharacterCollection = UpperCase + LowerCase

	Digit    CharacterCollection = `\d`
	NonDigit CharacterCollection = `\D`

	Word    CharacterCollection = `\w` // Letter + Digit + '_'
	NonWord CharacterCollection = `\W`

	WhiteSpace    CharacterCollection = `\s`
	NonWhiteSpace CharacterCollection = `\S`
	WordBoundry   CharacterCollection = `\b`
)

type group struct {
	pattern Pattern
	capture bool
}

func (g group) isExpression() {}

func (g group) ToString() string {
	pref := "("
	if !g.capture {
		pref += "?:"
	}
	return pref + g.pattern.ToString() + ")"
}

func MakeGroup(pattern Pattern, capture bool) Expression {
	return group{
		pattern: pattern,
		capture: capture,
	}
}
