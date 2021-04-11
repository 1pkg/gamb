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
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)+vi[1].(int)-vi[2].(int) == 7
			},
			out: NewVar(10, 1, 4),
		},
		"amb operator should produce expected result on single var": {
			in: []Var{
				NewVar(11, 15, 21, 30),
			},
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)%5 == 0
			},
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
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)/vi[2].(int) == 30
			},
			out: NewVar(300, 6, 10),
		},
		"amb operator should produce expected result on multiple unequal vars": {
			in: []Var{
				NewVar("1"),
				NewVar("2", "3"),
				NewVar("4", "5", "6"),
				NewVar("7", "8", "9", "A"),
			},
			fun: func(vi ...interface{}) bool {
				return vi[3].(string)+vi[2].(string)+vi[1].(string) == "842"
			},
			out: NewVar("1", "2", "4", "8"),
		},
		"amb operator should produce empty result on if there is no match": {
			in: []Var{
				NewVar(10, 20, 30),
				NewVar(1, 2, 3),
				NewVar(10, 20, 30),
			},
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)*vi[1].(int)*vi[2].(int) == 101
			},
			out: nil,
		},
		"amb should operator never panic": {
			in: []Var{
				NewVar("1"),
				NewVar("2", "3"),
				NewVar("4", "5", "6"),
				NewVar("7", "8", "9", "A"),
			},
			fun: func(vi ...interface{}) bool {
				return vi[10].(string) == "panic"
			},
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
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)+vi[1].(int)+vi[2].(int) == 15
			},
			out: NewVar(NewVar(10, 3, 2), NewVar(10, 2, 3), NewVar(10, 1, 4)),
		},
		"all operator should produce expected result on single var": {
			in: []Var{
				NewVar(11, 15, 21, 30),
			},
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)%5 == 0
			},
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
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)/vi[2].(int) == 30
			},
			out: NewVar(NewVar(300, 6, 10)),
		},
		"all operator should produce expected result on multiple unequal vars": {
			in: []Var{
				NewVar("1"),
				NewVar("2", "3"),
				NewVar("4", "5", "6"),
				NewVar("7", "8", "9", "A"),
			},
			fun: func(vi ...interface{}) bool {
				return vi[3].(string)+vi[2].(string)+vi[1].(string) == "842"
			},
			out: NewVar(NewVar("1", "2", "4", "8")),
		},
		"all operator should produce empty result on if there is no match": {
			in: []Var{
				NewVar(10, 20, 30),
				NewVar(1, 2, 3),
				NewVar(10, 20, 30),
			},
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)*vi[1].(int)*vi[2].(int) == 101
			},
			out: nil,
		},
		"all should operator never panic": {
			in: []Var{
				NewVar("1"),
				NewVar("2", "3"),
				NewVar("4", "5", "6"),
				NewVar("7", "8", "9", "A"),
			},
			fun: func(vi ...interface{}) bool {
				return vi[10].(string) == "panic"
			},
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
