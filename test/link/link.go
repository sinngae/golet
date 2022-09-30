package link

type Resource struct {
	name        string
	nextRscList []string
}

func Linkable(src *Resource, dest *Resource) bool {
	for _, item := range src.nextRscList {
		if item == dest.name {
			return true
		}
	}
	return false
}

func Link0(data []*[]*Resource) ([]*[]*Resource, []*[]*Resource) {
	size := len(data)
	if size <= 0 {
		return nil, nil
	}
	if size == 1 {
		resp := make([]*[]*Resource, 0)
		for _, item := range *data[0] {
			compose := make([]*Resource, 0)
			compose = append(compose, item)
			resp = append(resp, &compose)
		}
		return resp, nil
	}

	mid := 0
	if len(data)%2 == 0 {
		mid = len(data) / 2
	} else {
		mid = len(data)/2 + 1
	}

	mapA := data[0:mid]
	mapB := data[mid:]
	rspA, failA := Link0(mapA)
	rspB, failB := Link0(mapB)

	resp := make([]*[]*Resource, 0)
	fail := append(failA, failB...)
	for a := 0; a < len(rspA); a++ {
		for b := 0; b < len(rspB); b++ {
			if Linkable((*rspA[a])[len(*rspA[a])-1], (*rspB[b])[0]) {
				compose := append(*rspA[a], *rspB[b]...)
				resp = append(resp, &compose)
			} else {
				pair := []*Resource{(*rspA[a])[len(*rspA[a])-1], (*rspB[b])[0]}
				fail = append(fail, &pair)
			}
		}
	}
	return resp, fail
}

func Join(data []*Resource, split string) string {
	resp := ""
	for idx, item := range data {
		resp += item.name
		if idx+1 != len(data) {
			resp += split
		}
	}
	return resp
}

func Link1(data []*[]*Resource) ([]*[]*Resource, []*[]*Resource) {
	joinMap := transToComposeMap(data)
	res := joinMap[0]
	for idx := 1; idx < len(joinMap); idx++ {
		resTmp := make([]*[]*Resource, 0)
		for _, joinA := range *res {
			for _, joinB := range *joinMap[idx] {
				if (*joinA)[len(*joinA)-1].name == (*joinB)[0].name {
					compose := append(*joinA, (*joinB)[1:]...)
					resTmp = append(resTmp, &compose)
				}
			}
		}
		res = &resTmp
	}
	return *res, nil
}

func transToComposeMap(data []*[]*Resource) []*[]*[]*Resource {
	res := make([]*[]*[]*Resource, 0, len(data)-1)
	for idx := 1; idx < len(data); idx++ {
		posA := data[idx-1]
		posB := data[idx]
		join := make([]*[]*Resource, 0)
		for _, rsrA := range *posA {
			for _, rsrB := range *posB {
				if Linkable(rsrA, rsrB) {
					compose := []*Resource{rsrA, rsrB}
					join = append(join, &compose)
				}
			}
		}
		res = append(res, &join)
	}
	return res
}

func Link3(data *[]*[]*Resource) []*[]*Resource {
	type (
		join struct {
			rsr         *Resource
			composesIdx int
		}
	)
	stack := make([]*[]*Resource, 0)
	record := make(map[string]*join) //
	for idxB := 1; idxB < len(*data); idxB++ {
		idxA := idxB - 1
		compose := make([]*Resource, 0)
		for _, rsrA := range *(*data)[idxA] {
			for _, next := range rsrA.nextRscList {
				for _, rsrB := range *(*data)[idxB] {
					if next == rsrB.name {
						compose = append(compose, rsrA, rsrB)
						//record[]
						stack = append(stack, &compose)
					}
				}
			}
		}
	}
	return nil
}
