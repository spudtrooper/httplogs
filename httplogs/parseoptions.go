// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package httplogs

import "fmt"

type ParseOption struct {
	f func(*parseOptionImpl)
	s string
}

func (o ParseOption) String() string { return o.s }

type ParseOptions interface {
	Verbose() bool
	HasVerbose() bool
}

func ParseVerbose(verbose bool) ParseOption {
	return ParseOption{func(opts *parseOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("httplogs.ParseVerbose(bool %+v)", verbose)}
}
func ParseVerboseFlag(verbose *bool) ParseOption {
	return ParseOption{func(opts *parseOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("httplogs.ParseVerbose(bool %+v)", verbose)}
}

type parseOptionImpl struct {
	verbose     bool
	has_verbose bool
}

func (p *parseOptionImpl) Verbose() bool    { return p.verbose }
func (p *parseOptionImpl) HasVerbose() bool { return p.has_verbose }

func makeParseOptionImpl(opts ...ParseOption) *parseOptionImpl {
	res := &parseOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeParseOptions(opts ...ParseOption) ParseOptions {
	return makeParseOptionImpl(opts...)
}
