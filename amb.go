package gamb

// Var defines ambiguous variable type.
type Var []interface{}

// NewVar creates instance of new ambiguous variable from provided input.
func NewVar(v ...interface{}) Var {
	return Var(v)
}

// Func defines disambiguous predicate function type,
// that checks provided input against some condition.
type Func func(v ...interface{}) bool

// Amb is ambiguous operator implementation,
// that returns first ambiguous variable matching disambiguous predicate.
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
		// in case of empty var, we just need to go one level deeper with the same accumulator value.
		if len(v) == 0 {
			if vt := amb(accum, f, vars[1:]...); vt != nil {
				return vt
			}
		}
		// terminate parent loop unconditionally.
		// nolint
		return nil
	}
	// we can execute this check only for leaf most iterations.
	if ok := f(accum...); ok {
		// reslice accumulator here to prevent memory sharing with parent loops.
		return NewVar(accum...)
	}
	return nil
}

// All is ambiguous operator implementation,
// that returns all ambiguous variables matching disambiguous predicate.
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
		// in case of empty var, we just need to go one level deeper with the same accumulator value.
		if len(v) == 0 {
			if vt := all(accum, f, vars[1:]...); vt != nil {
				vout = append(vout, vt...)
			}
		}
		// terminate parent loop unconditionally.
		// nolint
		return vout
	}
	// we can execute this check only for leaf most iterations.
	if ok := f(accum...); ok {
		// reslice accumulator here to prevent memory sharing with parent loops.
		// also wrap it in var twice to follow slice of slices logic,
		// second wrapping will be consumed in parent loops.
		return NewVar(append(vout, accum...))
	}
	return nil
}
