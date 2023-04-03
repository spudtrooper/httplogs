// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package httplogs

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type ResolveIPsOption struct {
	f func(*resolveIPsOptionImpl)
	s string
}

func (o ResolveIPsOption) String() string { return o.s }

type ResolveIPsOptions interface {
	ForceResolveEmpties() bool
	HasForceResolveEmpties() bool
	Threads() int
	HasThreads() bool
	UseCache() bool
	HasUseCache() bool
	Verbose() bool
	HasVerbose() bool
}

func ResolveIPsForceResolveEmpties(forceResolveEmpties bool) ResolveIPsOption {
	return ResolveIPsOption{func(opts *resolveIPsOptionImpl) {
		opts.has_forceResolveEmpties = true
		opts.forceResolveEmpties = forceResolveEmpties
	}, fmt.Sprintf("httplogs.ResolveIPsForceResolveEmpties(bool %+v)", forceResolveEmpties)}
}
func ResolveIPsForceResolveEmptiesFlag(forceResolveEmpties *bool) ResolveIPsOption {
	return ResolveIPsOption{func(opts *resolveIPsOptionImpl) {
		if forceResolveEmpties == nil {
			return
		}
		opts.has_forceResolveEmpties = true
		opts.forceResolveEmpties = *forceResolveEmpties
	}, fmt.Sprintf("httplogs.ResolveIPsForceResolveEmpties(bool %+v)", forceResolveEmpties)}
}

func ResolveIPsThreads(threads int) ResolveIPsOption {
	return ResolveIPsOption{func(opts *resolveIPsOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}, fmt.Sprintf("httplogs.ResolveIPsThreads(int %+v)", threads)}
}
func ResolveIPsThreadsFlag(threads *int) ResolveIPsOption {
	return ResolveIPsOption{func(opts *resolveIPsOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}, fmt.Sprintf("httplogs.ResolveIPsThreads(int %+v)", threads)}
}

func ResolveIPsUseCache(useCache bool) ResolveIPsOption {
	return ResolveIPsOption{func(opts *resolveIPsOptionImpl) {
		opts.has_useCache = true
		opts.useCache = useCache
	}, fmt.Sprintf("httplogs.ResolveIPsUseCache(bool %+v)", useCache)}
}
func ResolveIPsUseCacheFlag(useCache *bool) ResolveIPsOption {
	return ResolveIPsOption{func(opts *resolveIPsOptionImpl) {
		if useCache == nil {
			return
		}
		opts.has_useCache = true
		opts.useCache = *useCache
	}, fmt.Sprintf("httplogs.ResolveIPsUseCache(bool %+v)", useCache)}
}

func ResolveIPsVerbose(verbose bool) ResolveIPsOption {
	return ResolveIPsOption{func(opts *resolveIPsOptionImpl) {
		opts.has_verbose = true
		opts.verbose = verbose
	}, fmt.Sprintf("httplogs.ResolveIPsVerbose(bool %+v)", verbose)}
}
func ResolveIPsVerboseFlag(verbose *bool) ResolveIPsOption {
	return ResolveIPsOption{func(opts *resolveIPsOptionImpl) {
		if verbose == nil {
			return
		}
		opts.has_verbose = true
		opts.verbose = *verbose
	}, fmt.Sprintf("httplogs.ResolveIPsVerbose(bool %+v)", verbose)}
}

type resolveIPsOptionImpl struct {
	forceResolveEmpties     bool
	has_forceResolveEmpties bool
	threads                 int
	has_threads             bool
	useCache                bool
	has_useCache            bool
	verbose                 bool
	has_verbose             bool
}

func (r *resolveIPsOptionImpl) ForceResolveEmpties() bool    { return r.forceResolveEmpties }
func (r *resolveIPsOptionImpl) HasForceResolveEmpties() bool { return r.has_forceResolveEmpties }
func (r *resolveIPsOptionImpl) Threads() int                 { return or.Int(r.threads, 10) }
func (r *resolveIPsOptionImpl) HasThreads() bool             { return r.has_threads }
func (r *resolveIPsOptionImpl) UseCache() bool               { return r.useCache }
func (r *resolveIPsOptionImpl) HasUseCache() bool            { return r.has_useCache }
func (r *resolveIPsOptionImpl) Verbose() bool                { return r.verbose }
func (r *resolveIPsOptionImpl) HasVerbose() bool             { return r.has_verbose }

func makeResolveIPsOptionImpl(opts ...ResolveIPsOption) *resolveIPsOptionImpl {
	res := &resolveIPsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeResolveIPsOptions(opts ...ResolveIPsOption) ResolveIPsOptions {
	return makeResolveIPsOptionImpl(opts...)
}
