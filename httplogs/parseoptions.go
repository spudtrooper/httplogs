// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package httplogs

import "fmt"

type ParseOption struct {
	f func(*parseOptionImpl)
	s string
}

func (o ParseOption) String() string { return o.s }

type ParseOptions interface {
	ResolveIPs() bool
	HasResolveIPs() bool
}

func ParseResolveIPs(resolveIPs bool) ParseOption {
	return ParseOption{func(opts *parseOptionImpl) {
		opts.has_resolveIPs = true
		opts.resolveIPs = resolveIPs
	}, fmt.Sprintf("httplogs.ParseResolveIPs(bool %+v)", resolveIPs)}
}
func ParseResolveIPsFlag(resolveIPs *bool) ParseOption {
	return ParseOption{func(opts *parseOptionImpl) {
		if resolveIPs == nil {
			return
		}
		opts.has_resolveIPs = true
		opts.resolveIPs = *resolveIPs
	}, fmt.Sprintf("httplogs.ParseResolveIPs(bool %+v)", resolveIPs)}
}

type parseOptionImpl struct {
	resolveIPs     bool
	has_resolveIPs bool
}

func (p *parseOptionImpl) ResolveIPs() bool    { return p.resolveIPs }
func (p *parseOptionImpl) HasResolveIPs() bool { return p.has_resolveIPs }

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
