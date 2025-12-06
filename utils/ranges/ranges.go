package ranges

type Range struct {
	Start int
	End   int
}

type Ranges []Range

func (r Ranges) Len() int {
	return len(r)
}

func (r *Ranges) Sort() {
	for i := 0; i < len(*r)-1; i++ {
		for j := i + 1; j < len(*r); j++ {
			if (*r)[i].Start > (*r)[j].Start {
				(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
			}
		}
	}
}

func (r Ranges) Merge() Ranges {
	merged := Ranges{}

	for _, current := range r {
		if len(merged) == 0 {
			merged = append(merged, current)
			continue
		}

		last := &merged[len(merged)-1]
		if current.Start <= last.End+1 {
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}
