package rank

import (
	. "recommend/model"
)

func tagscore(usertag map[string]float64, art *Article) float64 {
	var hs float64 = 0.0
	for k, v := range usertag {
		if k == art.GetLargeclass() {
			hs += v
			continue
		}
		if k == art.GetMediumclass() {
			hs += v
			continue
		}
		if k == art.GetSmallclass() {
			hs += v
			continue
		}
		for _, x := range art.GetTags() {
			if k == x {
				hs += v
				continue
			}
		}
	}
	return hs
}
