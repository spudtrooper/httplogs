package httplogs

import (
	"sort"
)

type histItem struct {
	Key  string
	Recs []Record
}

type hist []histItem

func (a hist) Len() int           { return len(a) }
func (a hist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hist) Less(i, j int) bool { return len(a[i].Recs) < len(a[j].Recs) }

func toHist(m map[string][]Record) hist {
	var res hist
	for k, recs := range m {
		res = append(res, histItem{k, recs})
	}
	return res
}

type GroupResult struct {
	ByIP        hist
	ByPath      hist
	ByUserAgent hist
}

func Group(recs []*Record) GroupResult {

	var byPath hist
	{
		r := map[string][]Record{}
		for _, rec := range recs {
			r[rec.Path] = append(r[rec.Path], *rec)
		}
		byPath = toHist(r)
		sort.Sort(byPath)
	}
	var byIP hist
	{
		r := map[string][]Record{}
		for _, rec := range recs {
			r[rec.IP] = append(r[rec.IP], *rec)
		}
		byIP = toHist(r)
		sort.Sort(byIP)
	}
	var byUserAgent hist
	{
		r := map[string][]Record{}
		for _, rec := range recs {
			r[rec.UserAgent] = append(r[rec.UserAgent], *rec)
		}
		byUserAgent = toHist(r)
		sort.Sort(byUserAgent)
	}

	res := GroupResult{
		ByIP:        byIP,
		ByPath:      byPath,
		ByUserAgent: byUserAgent,
	}
	return res
}
