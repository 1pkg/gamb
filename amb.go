package gamb

type Operator func(Func, ...Var) Var

type Var []interface{}

func NewVar(v ...interface{}) Var {
	return Var(v)
}

type Func func(v ...interface{}) bool

func Amb(f Func, vars ...Var) Var {
	defer func() { _ = recover() }()
	return amb(nil, f, vars...)
}

func amb(accum Var, f Func, vars ...Var) Var {
	for _, v := range vars {
		for _, i := range v {
			if vt := amb(append(accum, i), f, vars[1:]...); vt != nil {
				return vt
			}
		}
		if len(v) == 0 {
			if vt := amb(accum, f, vars[1:]...); vt != nil {
				return vt
			}
		}
		// nolint
		return nil
	}
	if ok := f(accum...); ok {
		return NewVar(accum...)
	}
	return nil
}

func All(f Func, vars ...Var) Var {
	defer func() { _ = recover() }()
	return all(nil, f, vars...)
}

func all(accum Var, f Func, vars ...Var) Var {
	var vout Var
	for _, v := range vars {
		for _, i := range v {
			if vt := all(append(accum, i), f, vars[1:]...); vt != nil {
				vout = append(vout, vt...)
			}
		}
		if len(v) == 0 {
			if vt := all(accum, f, vars[1:]...); vt != nil {
				vout = append(vout, vt...)
			}
		}
		// nolint
		return vout
	}
	if ok := f(accum...); ok {
		return NewVar(append(vout, accum...))
	}
	return nil
}
