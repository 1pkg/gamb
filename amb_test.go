package gamb

import (
	"reflect"
	"testing"
)

func TestAmb(t *testing.T) {
	table := map[string]struct {
		in  []AmbVar
		fun AmbFunc
		out AmbVar
	}{
		"amb operator should produce expected result on multiple vars": {
			in: []AmbVar{
				NewAmbVar(10, 20, 30),
				NewAmbVar(1, 2, 3, 5, 10),
				NewAmbVar(2, 3, 4),
			},
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)+vi[1].(int)-vi[2].(int) == 7
			},
			out: NewAmbVar(10, 1, 4),
		},
		"amb operator should produce expected result on single var": {
			in: []AmbVar{
				NewAmbVar(11, 15, 21, 30),
			},
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)%5 == 0
			},
			out: NewAmbVar(15),
		},
		"amb operator should skip empty vars and produce expected result on multiple vars": {
			in: []AmbVar{
				NewAmbVar(100, 200, 300),
				NewAmbVar(),
				NewAmbVar(),
				NewAmbVar(6),
				NewAmbVar(2, 10),
			},
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)/vi[2].(int) == 30
			},
			out: NewAmbVar(300, 6, 10),
		},
		"amb operator should produce expected result on multiple unequal vars": {
			in: []AmbVar{
				NewAmbVar("1"),
				NewAmbVar("2", "3"),
				NewAmbVar("4", "5", "6"),
				NewAmbVar("7", "8", "9", "A"),
			},
			fun: func(vi ...interface{}) bool {
				return vi[3].(string)+vi[2].(string)+vi[1].(string) == "842"
			},
			out: NewAmbVar("1", "2", "4", "8"),
		},
		"amb operator should produce empty result on if there is no match": {
			in: []AmbVar{
				NewAmbVar(10, 20, 30),
				NewAmbVar(1, 2, 3),
				NewAmbVar(10, 20, 30),
			},
			fun: func(vi ...interface{}) bool {
				return vi[0].(int)*vi[1].(int)*vi[2].(int) == 101
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
