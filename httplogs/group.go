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

type intHistItem struct {
	Key  int
	Recs []Record
}

type intHist []intHistItem

func (a intHist) Len() int           { return len(a) }
func (a intHist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intHist) Less(i, j int) bool { return len(a[i].Recs) < len(a[j].Recs) }

func toIntHist(m map[int][]Record) intHist {
	var res intHist
	for k, recs := range m {
		res = append(res, intHistItem{k, recs})
	}
	return res
}

type GroupResult struct {
	ByIP         hist
	ByPath       hist
	ByUserAgent  hist
	ByStatusCode intHist
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
	var byStatusCode intHist
	{
		r := map[int][]Record{}
		for _, rec := range recs {
			r[rec.StatusCode] = append(r[rec.StatusCode], *rec)
		}
		byStatusCode = toIntHist(r)
		sort.Sort(byStatusCode)
	}

	res := GroupResult{
		ByIP:         byIP,
		ByPath:       byPath,
		ByUserAgent:  byUserAgent,
		ByStatusCode: byStatusCode,
	}
	return res
}
