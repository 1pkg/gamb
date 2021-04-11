# Gamb: go amb (ambiguous) operator implementation

[![lint](https://github.com/1pkg/gamb/workflows/lint/badge.svg)](https://github.com/1pkg/gamb/actions?query=workflow%3Alint+branch%3Amaster+)
[![build](https://github.com/1pkg/gamb/workflows/build/badge.svg)](https://github.com/1pkg/gamb/actions?query=workflow%3Abuild+branch%3Amaster+)
[![report](https://goreportcard.com/badge/github.com/1pkg/gamb)](https://goreportcard.com/report/github.com/1pkg/gamb)
[![version](https://img.shields.io/github/go-mod/go-version/1pkg/gamb)](https://github.com/1pkg/gamb/blob/master/go.mod)
[![license](https://img.shields.io/github/license/1pkg/gamb)](LICENSE)

`go get -u github.com/1pkg/gamb`

## Details

This package provides generic variadic implemention of [McCarthy's Ambiguous Operator](https://rosettacode.org/wiki/Amb) in go. Gamb exposes three ambiguous functions `Amb` to yield first variable matching ambiguous predicate; `All` to yield all variables matching ambiguous predicate; `Ord` to yield first strict ordered variable matching ambiguous predicate in single pass. Ambiguous predicate is defined as function `func(v ...interface{}) bool` that accepts ambiguous variable sets permutations and check some bolean condition against them.

```go
    out := Amb(
        func(v ...interface{}) bool {
            return v[0].(int)+v[1].(int)-v[2].(int) == 7
        }
        NewVar(10, 20, 30),
        NewVar(1, 2, 3, 5, 10),
        NewVar(2, 3, 4),
    )
    fmt.Println(out) // [10 1 4]
```

```go
    out := All(
        func(v ...interface{}) bool {
            return v[0].(int)+v[1].(int)-v[2].(int) == 7
        }
        NewVar(10, 20, 30),
        NewVar(1, 2, 3, 5, 10),
        NewVar(2, 3, 4),
    )
    fmt.Println(out) // [[10 1 4], [10, 2, 3], [10, 3, 2]]
```

```go
    out := Ord(
        func(v ...interface{}) bool {
            return v[0].(int)+v[1].(int)-v[2].(int) == 31
        }
        NewVar(10, 20, 30),
        NewVar(1, 2, 3, 5, 10),
        NewVar(2, 3, 4),
    )
    fmt.Println(out) // [30 5 4]
```

**Note:** ambiguous operator generally requires some form of backtracking with `n^m` operators, where `n` - size of ambiguous variable and `m` - number of ambiguous variables (for `ord` operator single pass is used reducing complexity to `n`). Therefore ambiguous operator is not efficient on processing big inputs.

## Licence

Gamb is licensed under the MIT License.  
See [LICENSE](LICENSE) for the full license text.
