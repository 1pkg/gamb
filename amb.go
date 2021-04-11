package gamb

type AmbVar []interface{}

func NewAmbVar(v ...interface{}) AmbVar {
	return AmbVar(v)
}

type AmbFunc func(vi ...interface{}) bool

func Amb(f AmbFunc, vars ...AmbVar) AmbVar {
	set, ok := try(f, vars)
	if !ok {
	backtrack:
		for i, v := range vars {
			if len(v) <= 1 {
				break
			}
			if tset := Amb(f, mutate(vars, i)...); tset != nil {
				set = tset
				break backtrack
			}
		}
	}
	return set
}

func mutate(vars []AmbVar, i int) []AmbVar {
	mvars := make([]AmbVar, len(vars))
	for i := range vars {
		mvars[i] = make(AmbVar, len(vars[i]))
		copy(mvars[i], vars[i])
	}
	if len(mvars[i]) > 1 {
		mvars[i] = mvars[i][1:]
	}
	return mvars
}

func try(f AmbFunc, vars []AmbVar) (AmbVar, bool) {
	set := make(AmbVar, 0, len(vars))
	for _, v := range vars {
		for _, vi := range v {
			set = append(set, vi)
			break
		}
	}
	ok := f(set...)
	if ok {
		return set, true
	}
	return nil, false
}
