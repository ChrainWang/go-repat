package repat

import "regexp"

type compileOptions struct {
	matchStart  bool
	matchEnd    bool
	ignoreCases bool
	multiLine   bool
}

type CompileOption func(*compileOptions)

func MatchStart() CompileOption {
	return func(opts *compileOptions) {
		opts.matchStart = true
	}
}

func MatchEnd() CompileOption {
	return func(opts *compileOptions) {
		opts.matchEnd = true
	}
}

func MatchEntier() CompileOption {
	return func(opts *compileOptions) {
		opts.matchStart = true
		opts.matchEnd = true
	}
}

func IgnoreCases() CompileOption {
	return func(opts *compileOptions) {
		opts.ignoreCases = true
	}
}

func MultipleLineMode() CompileOption {
	return func(opts *compileOptions) {
		opts.multiLine = true
	}
}

func MustCompile(pattern Pattern, options ...CompileOption) *regexp.Regexp {
	opts := &compileOptions{}
	for _, o := range options {
		o(opts)
	}

	if opts.matchStart {
		if opts.multiLine {
			pattern = Join(SOS, pattern)
		} else {
			pattern = Join(SOL, pattern)
		}
	}
	if opts.matchEnd {
		if opts.multiLine {
			pattern = Join(pattern, EOS)
		} else {
			pattern = Join(pattern, EOL)
		}
	}

	return regexp.MustCompile(pattern.ToString())
}
