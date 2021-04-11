package gamb

import "reflect"

type Operator func(Func, ...Var) Var

var (
	Amb = wrap(amb)
	All = wrap(all)
)

type Var []interface{}

func NewVar(v ...interface{}) Var {
	return Var(v)
}

type Func func(vi ...interface{}) bool

func amb(f Func, vars ...Var) Var {
	vout, ok := try(f, vars)
	if !ok {
	backtrack:
		for i, v := range vars {
			if len(v) <= 1 {
				continue
			}
			if tv := amb(f, mutate(vars, i)...); tv != nil {
				vout = tv
				break backtrack
			}
		}
	}
	return vout
}

func all(f Func, vars ...Var) Var {
	var accum Var
	if v, ok := try(f, vars); ok {
		accum = append(accum, v)
	}
	for i, v := range vars {
		if len(v) <= 1 {
			continue
		}
		if v := all(f, mutate(vars, i)...); v != nil {
			accum = append(accum, v...)
		}
	}
	return accum
}

func mutate(vars []Var, i int) []Var {
	mvars := make([]Var, len(vars))
	for i := range vars {
		mvars[i] = make(Var, len(vars[i]))
		copy(mvars[i], vars[i])
	}
	mvars[i] = mvars[i][1:]
	return mvars
}

func try(f Func, vars []Var) (Var, bool) {
	set := make(Var, 0, len(vars))
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

func wrap(op Operator) Operator {
	return func(f Func, vars ...Var) Var {
		defer func() {
			_ = recover()
		}()
		v := op(f, vars...)
		if v == nil {
			return nil
		}
		vuniq := make(Var, 0, len(v))
		for _, el := range v {
			if i := find(vuniq, el); i == -1 {
				vuniq = append(vuniq, el)
			}
		}
		return vuniq
	}
}

func find(v Var, el interface{}) int {
	for i, vel := range v {
		if reflect.DeepEqual(el, vel) {
			return i
		}
	}
	return -1
}
