package gamb

import (
	"reflect"
	"testing"
)

func TestAmb(t *testing.T) {
	table := map[string]struct {
		in  []Var
		fun Func
		out Var
	}{
		"amb operator should produce expected result on multiple vars": {
			in: []Var{
				NewVar(10, 20, 30),
				NewVar(1, 2, 3, 5, 10),
				NewVar(2, 3, 4),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)+v[1].(int)-v[2].(int) == 7 },
			out: NewVar(10, 1, 4),
		},
		"amb operator should produce expected result on multiple vars long sequence": {
			in: []Var{
				NewVar(10, 20, 30, 40, 55, 100, 99, 50),
				NewVar(1, 2, 3, 5, 11),
				NewVar(2, 3, 4, 0),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)*v[1].(int)+v[2].(int) == 550 },
			out: NewVar(50, 11, 0),
		},
		"amb operator should produce expected result on single var": {
			in: []Var{
				NewVar(11, 15, 21, 30),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)%5 == 0 },
			out: NewVar(15),
		},
		"amb operator should skip empty vars and produce expected result on multiple vars": {
			in: []Var{
				NewVar(100, 200, 300),
				NewVar(),
				NewVar(),
				NewVar(6),
				NewVar(2, 10),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)/v[2].(int) == 30 },
			out: NewVar(300, 6, 10),
		},
		"amb operator should produce expected result on multiple unequal vars": {
			in: []Var{
				NewVar("1"),
				NewVar("2", "3"),
				NewVar("4", "5", "6"),
				NewVar("7", "8", "9", "A"),
			},
			fun: func(v ...interface{}) bool { return v[3].(string)+v[2].(string)+v[1].(string) == "842" },
			out: NewVar("1", "2", "4", "8"),
		},
		"amb operator should produce empty result on if there is no match": {
			in: []Var{
				NewVar(10, 20, 30),
				NewVar(1, 2, 3),
				NewVar(10, 20, 30),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)*v[1].(int)*v[2].(int) == 101 },
			out: nil,
		},
		"amb should operator never panic": {
			in: []Var{
				NewVar("1"),
				NewVar("2", "3"),
				NewVar("4", "5", "6"),
				NewVar("7", "8", "9", "A"),
			},
			fun: func(v ...interface{}) bool { return v[10].(string) == "panic" },
			out: nil,
		},
	}
	for tname, tcase := range table {
		t.Run(tname, func(t *testing.T) {
			out := Amb(tcase.fun, tcase.in...)
			if !reflect.DeepEqual(tcase.out, out) {
				t.Fatalf("amb expected result %v but got %v", tcase.out, out)
			}
		})
	}
}

func TestAll(t *testing.T) {
	table := map[string]struct {
		in  []Var
		fun Func
		out Var
	}{
		"all operator should produce expected result on multiple vars": {
			in: []Var{
				NewVar(10, 20, 30),
				NewVar(1, 2, 3, 5, 10),
				NewVar(2, 3, 4),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)+v[1].(int)+v[2].(int) == 15 },
			out: NewVar(NewVar(10, 1, 4), NewVar(10, 2, 3), NewVar(10, 3, 2)),
		},
		"all operator should produce expected result on multiple vars with duplicates": {
			in: []Var{
				NewVar(10, 10, 10, 20, 20),
				NewVar(1, 2, 3, 5, 10),
				NewVar(1),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)*v[1].(int)*v[2].(int) == 20 },
			out: NewVar(NewVar(10, 2, 1), NewVar(10, 2, 1), NewVar(10, 2, 1), NewVar(20, 1, 1), NewVar(20, 1, 1)),
		},
		"all operator should produce expected result on single var": {
			in: []Var{
				NewVar(11, 15, 21, 30),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)%5 == 0 },
			out: NewVar(NewVar(15), NewVar(30)),
		},
		"all operator should skip empty vars and produce expected result on multiple vars": {
			in: []Var{
				NewVar(100, 200, 300),
				NewVar(),
				NewVar(),
				NewVar(6),
				NewVar(2, 10),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)/v[2].(int) == 30 },
			out: NewVar(NewVar(300, 6, 10)),
		},
		"all operator should produce expected result on multiple unequal vars": {
			in: []Var{
				NewVar("1"),
				NewVar("2", "3"),
				NewVar("4", "5", "6"),
				NewVar("7", "8", "9", "A"),
			},
			fun: func(v ...interface{}) bool { return v[3].(string)+v[2].(string)+v[1].(string) == "842" },
			out: NewVar(NewVar("1", "2", "4", "8")),
		},
		"all operator should produce empty result on if there is no match": {
			in: []Var{
				NewVar(10, 20, 30),
				NewVar(1, 2, 3),
				NewVar(10, 20, 30),
			},
			fun: func(v ...interface{}) bool { return v[0].(int)*v[1].(int)*v[2].(int) == 101 },
			out: nil,
		},
		"all should operator never panic": {
			in: []Var{
				NewVar("1"),
				NewVar("2", "3"),
				NewVar("4", "5", "6"),
				NewVar("7", "8", "9", "A"),
			},
			fun: func(v ...interface{}) bool { return v[10].(string) == "panic" },
			out: nil,
		},
	}
	for tname, tcase := range table {
		t.Run(tname, func(t *testing.T) {
			out := All(tcase.fun, tcase.in...)
			if !reflect.DeepEqual(tcase.out, out) {
				t.Fatalf("all expected result %v but got %v", tcase.out, out)
			}
		})
	}
}

func TestCombination(t *testing.T) {
	v := All(func(v ...interface{}) bool {
		return v[0].(int)%v[1].(int) == 0
	},
		Amb(func(v ...interface{}) bool {
			return v[0].(int)-v[1].(int) > 100
		},
			NewVar(100, 200, 300),
			NewVar(300, 200, 100),
		),
		NewVar(3, 5),
	)
	e := NewVar(NewVar(300, 3), NewVar(300, 5), NewVar(100, 5))
	if !reflect.DeepEqual(e, v) {
		t.Fatalf("all expected result %v but got %v", e, v)
	}
}
