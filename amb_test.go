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
	}{}
	for tname, tcase := range table {
		t.Run(tname, func(t *testing.T) {
			out := Amb(tcase.fun, tcase.in...)
			if !reflect.DeepEqual(tcase.out, out) {
				t.Fatalf("amb expected result %v but got %v", tcase.out, out)
			}
		})
	}
}
