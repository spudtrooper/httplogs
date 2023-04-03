// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package httplogs

import "fmt"

type FilterOption struct {
	f func(*filterOptionImpl)
	s string
}

func (o FilterOption) String() string { return o.s }

type FilterOptions interface {
	StatusCodes() []int
	HasStatusCodes() bool
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

type filterOptionImpl struct {
	statusCodes     []int
	has_statusCodes bool
}

func (f *filterOptionImpl) StatusCodes() []int   { return f.statusCodes }
func (f *filterOptionImpl) HasStatusCodes() bool { return f.has_statusCodes }

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
