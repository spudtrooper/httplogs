package httplogs

import (
	"log"
	"regexp"
)

type filterStats map[string]int

func (s filterStats) Add(by string) { s[by]++ }

//go:generate genopts --function Filter statusCodes:[]int pathFilter:string negPathFilter:string userAgentFilter:string negUserAgentFilter:string
func Filter(recs []Record, optss ...FilterOption) []Record {
	opts := MakeFilterOptions(optss...)

	statusCodes := opts.StatusCodes()
	var pathFilter *regexp.Regexp
	if opts.PathFilter() != "" {
		pathFilter = regexp.MustCompile(opts.PathFilter())
	}
	var negPathFilter *regexp.Regexp
	if opts.NegPathFilter() != "" {
		negPathFilter = regexp.MustCompile(opts.NegPathFilter())
	}
	var userAgentFilter *regexp.Regexp
	if opts.UserAgentFilter() != "" {
		userAgentFilter = regexp.MustCompile(opts.UserAgentFilter())
	}
	var negUserAgentFilter *regexp.Regexp
	if opts.NegUserAgentFilter() != "" {
		negUserAgentFilter = regexp.MustCompile(opts.NegUserAgentFilter())
	}

	var res []Record
	stats := filterStats{}
	for _, rec := range recs {
		if len(statusCodes) > 0 {
			if !inInSlice(rec.StatusCode, statusCodes) {
				stats.Add("statusCode")
				continue
			}
		}
		if pathFilter != nil {
			if !pathFilter.MatchString(rec.Path) {
				stats.Add("path")
				continue
			}
		}
		if negPathFilter != nil {
			if negPathFilter.MatchString(rec.Path) {
				stats.Add("negpath")
				continue
			}
		}
		if userAgentFilter != nil {
			if !userAgentFilter.MatchString(rec.UserAgent) {
				stats.Add("useragent")
				continue
			}
		}
		if negUserAgentFilter != nil {
			if negUserAgentFilter.MatchString(rec.UserAgent) {
				stats.Add("neguseragent")
				continue
			}
		}
		res = append(res, rec)
	}

	log.Printf("filter stats: %+v", stats)

	return res
}

func inInSlice(needle int, haystack []int) bool {
	for _, it := range haystack {
		if needle == it {
			return true
		}
	}
	return false
}
