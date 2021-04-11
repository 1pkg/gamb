package gamb

type AmbVar []interface{}

func NewAmbVar(v ...interface{}) AmbVar {
	return AmbVar(v)
}

type AmbFunc func(vi ...interface{}) bool

func Amb(c int, f AmbFunc, vars ...AmbVar) AmbVar {
	set := make(AmbVar, 0, len(vars))
	for _, v := range vars {
		for _, vi := range v {
			set = append(set, vi)
			break
		}
	}
	if !f(set...) {
		set = nil
	backtrack:
		for i, v := range vars {
			if len(v) <= 1 {
				break
			}
			nwars := make([]AmbVar, len(vars))
			for i := range vars {
				nwars[i] = make(AmbVar, len(vars[i]))
				copy(nwars[i], vars[i])
			}
			nwars[i] = v[1:]
			if tset := Amb(c+1, f, nwars...); len(tset) != 0 {
				set = tset
				break backtrack
			}
		}
	}
	return set
}
