package httplogs

//go:generate genopts --function Filter statusCodes:[]int
func Filter(recs []*Record, optss ...FilterOption) []*Record {
	opts := MakeFilterOptions(optss...)

	var res []*Record
	statusCodes := opts.StatusCodes()
	for _, rec := range recs {
		if len(statusCodes) > 0 {
			if !inInSlice(rec.StatusCode, statusCodes) {
				continue
			}
		}
		res = append(res, rec)
	}

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
