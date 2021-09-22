package repat

type Pattern interface {
	ToString() string
}

type Expression interface {
	Pattern
	isExpression()
}
