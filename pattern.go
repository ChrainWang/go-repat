package repat

import (
	"strconv"
)

const (
	SOL stringPattern = `^`  // start of line
	EOL stringPattern = `$`  // end of line
	SOS stringPattern = `\A` // start of string
	EOS stringPattern = `\z` // end of string

)

var (
	zeroOrOnce = qualifier{min: 0, max: 1} // ?
	zeroOrMore = qualifier{min: 0, max: 0} // *
	onceOrMore = qualifier{min: 1, max: 0} // +
)

type QualifierProvider func() *qualifier

func ZeroOrOnce() *qualifier {
	return &zeroOrOnce
}

func ZeroOrMore() *qualifier {
	return &zeroOrMore
}

func OnceOrMore() *qualifier {
	return &onceOrMore
}

func WithMatchCount(min, max uint64) QualifierProvider {
	return func() *qualifier {
		return &qualifier{min: min, max: max}
	}
}

func WithMaximumMatchCount(max uint64) QualifierProvider {
	return func() *qualifier {
		return &qualifier{max: max}
	}
}

func WithMinimumMatchCount(min uint64) QualifierProvider {
	return func() *qualifier {
		return &qualifier{min: min}
	}

}

type qualifier struct {
	min, max uint64
	lazyMode bool
}

func (q qualifier) ToString() (s string) {
	defer func() {
		if q.lazyMode && q.min != q.max {
			s += "?"
		}
	}()

	switch {
	case q.min == q.max:
		if q.min == 0 {
			s = "*"
		} else if q.min != 1 {
			s = "{" + strconv.FormatUint(q.min, 10) + "}"
		}
	case q.min == 0:
		if q.max == 1 {
			s = "?"
		} else {
			s = "{," + strconv.FormatUint(q.max, 10) + "}"
		}
	case q.max == 0:
		if q.min == 1 {
			s = "+"
		} else {
			s = "{" + strconv.FormatUint(q.min, 10) + ",}"
		}
	default:
		s = "{" + strconv.FormatUint(q.min, 10) + "," + strconv.FormatUint(q.max, 10) + "}"
	}
	return
}

func (q *qualifier) LazyMode() *qualifier {
	if q.lazyMode {
		return q
	}

	return &qualifier{
		min: q.min,
		max: q.max,
	}
}

type stringPattern string

func (p stringPattern) ToString() string {
	return string(p)
}

type patternImpl struct {
	re Expression
	q  *qualifier
}

func (p patternImpl) ToString() string {
	if p.q == nil {
		return p.re.ToString()
	}
	return p.re.ToString() + p.q.ToString()
}

func MakePattern(any interface{}, qp QualifierProvider) Pattern {
	var p Pattern
	switch any.(type) {
	case Pattern:
		p = any.(Pattern)
	case string:
		p = stringPattern(any.(string))
	default:
		panic("invalid value to make regular expression")
	}

	if qp == nil {
		return p
	}

	return &patternImpl{
		re: asExpression(p),
		q:  qp(),
	}
}

func asExpression(p Pattern) Expression {
	if re, ok := p.(Expression); ok {
		return re
	}
	return MakeGroup(p, false)
}

func Join(patterns ...Pattern) Pattern {
	var s string
	for _, p := range patterns {
		if p == nil {
			continue
		}
		s += p.ToString()
	}
	return stringPattern(s)
}

func Or(patterns ...Pattern) Expression {
	if len(patterns) == 1 {
		return asExpression(patterns[0])
	}

	var s string
	for _, pattern := range patterns {
		s += "|" + pattern.ToString()
	}
	s = s[1:]
	return MakeGroup(stringPattern(s), false)
}
