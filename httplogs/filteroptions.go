// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package httplogs

import "fmt"

type FilterOption struct {
	f func(*filterOptionImpl)
	s string
}

func (o FilterOption) String() string { return o.s }

type FilterOptions interface {
	NegPathFilter() string
	HasNegPathFilter() bool
	NegUserAgentFilter() string
	HasNegUserAgentFilter() bool
	PathFilter() string
	HasPathFilter() bool
	StatusCodes() []int
	HasStatusCodes() bool
	UserAgentFilter() string
	HasUserAgentFilter() bool
}

func FilterNegPathFilter(negPathFilter string) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		opts.has_negPathFilter = true
		opts.negPathFilter = negPathFilter
	}, fmt.Sprintf("httplogs.FilterNegPathFilter(string %+v)", negPathFilter)}
}
func FilterNegPathFilterFlag(negPathFilter *string) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		if negPathFilter == nil {
			return
		}
		opts.has_negPathFilter = true
		opts.negPathFilter = *negPathFilter
	}, fmt.Sprintf("httplogs.FilterNegPathFilter(string %+v)", negPathFilter)}
}

func FilterNegUserAgentFilter(negUserAgentFilter string) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		opts.has_negUserAgentFilter = true
		opts.negUserAgentFilter = negUserAgentFilter
	}, fmt.Sprintf("httplogs.FilterNegUserAgentFilter(string %+v)", negUserAgentFilter)}
}
func FilterNegUserAgentFilterFlag(negUserAgentFilter *string) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		if negUserAgentFilter == nil {
			return
		}
		opts.has_negUserAgentFilter = true
		opts.negUserAgentFilter = *negUserAgentFilter
	}, fmt.Sprintf("httplogs.FilterNegUserAgentFilter(string %+v)", negUserAgentFilter)}
}

func FilterPathFilter(pathFilter string) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		opts.has_pathFilter = true
		opts.pathFilter = pathFilter
	}, fmt.Sprintf("httplogs.FilterPathFilter(string %+v)", pathFilter)}
}
func FilterPathFilterFlag(pathFilter *string) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		if pathFilter == nil {
			return
		}
		opts.has_pathFilter = true
		opts.pathFilter = *pathFilter
	}, fmt.Sprintf("httplogs.FilterPathFilter(string %+v)", pathFilter)}
}

func FilterStatusCodes(statusCodes []int) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		opts.has_statusCodes = true
		opts.statusCodes = statusCodes
	}, fmt.Sprintf("httplogs.FilterStatusCodes([]int %+v)", statusCodes)}
}
func FilterStatusCodesFlag(statusCodes *[]int) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		if statusCodes == nil {
			return
		}
		opts.has_statusCodes = true
		opts.statusCodes = *statusCodes
	}, fmt.Sprintf("httplogs.FilterStatusCodes([]int %+v)", statusCodes)}
}

func FilterUserAgentFilter(userAgentFilter string) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		opts.has_userAgentFilter = true
		opts.userAgentFilter = userAgentFilter
	}, fmt.Sprintf("httplogs.FilterUserAgentFilter(string %+v)", userAgentFilter)}
}
func FilterUserAgentFilterFlag(userAgentFilter *string) FilterOption {
	return FilterOption{func(opts *filterOptionImpl) {
		if userAgentFilter == nil {
			return
		}
		opts.has_userAgentFilter = true
		opts.userAgentFilter = *userAgentFilter
	}, fmt.Sprintf("httplogs.FilterUserAgentFilter(string %+v)", userAgentFilter)}
}

type filterOptionImpl struct {
	negPathFilter          string
	has_negPathFilter      bool
	negUserAgentFilter     string
	has_negUserAgentFilter bool
	pathFilter             string
	has_pathFilter         bool
	statusCodes            []int
	has_statusCodes        bool
	userAgentFilter        string
	has_userAgentFilter    bool
}

func (f *filterOptionImpl) NegPathFilter() string       { return f.negPathFilter }
func (f *filterOptionImpl) HasNegPathFilter() bool      { return f.has_negPathFilter }
func (f *filterOptionImpl) NegUserAgentFilter() string  { return f.negUserAgentFilter }
func (f *filterOptionImpl) HasNegUserAgentFilter() bool { return f.has_negUserAgentFilter }
func (f *filterOptionImpl) PathFilter() string          { return f.pathFilter }
func (f *filterOptionImpl) HasPathFilter() bool         { return f.has_pathFilter }
func (f *filterOptionImpl) StatusCodes() []int          { return f.statusCodes }
func (f *filterOptionImpl) HasStatusCodes() bool        { return f.has_statusCodes }
func (f *filterOptionImpl) UserAgentFilter() string     { return f.userAgentFilter }
func (f *filterOptionImpl) HasUserAgentFilter() bool    { return f.has_userAgentFilter }

func makeFilterOptionImpl(opts ...FilterOption) *filterOptionImpl {
	res := &filterOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeFilterOptions(opts ...FilterOption) FilterOptions {
	return makeFilterOptionImpl(opts...)
}
