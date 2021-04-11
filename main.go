package amb

type AmbVar []interface{}

func NewAmbVar(v ...interface{}) AmbVar {
	return AmbVar(v)
}

type AmbFunc func(p ...interface{}) bool

func Amb(f AmbFunc, vars ...AmbVar) AmbVar {
	return nil
}
