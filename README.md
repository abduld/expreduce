This software is experimental quality and is not currently intended for serious use. There are plenty of more mature open source computer algebra systems to use instead.

# Source generation
Generate source files from lex and yacc:
```
go generate
```

# Example
This must be done after running "go generate". To run the example CAS prompt:

```
cd example
go run calc.go
```

```
# go run calc.go

> D[Sin[x]/x,x]
In:  D[(Sin[x] * x^-1), x]
Out: ((Cos[x] * x^-1) + (Sin[x] * -1 * x^-2))

> Table[a^2,{a,1,10}]
In:  Table[a^2, {a, 1, 10}]
Out: {1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

> Sum[i, {i, 1, n}]
In:  Sum[i, {i, 1, n}]
Out: (2^-1 * n * (1 + n))

> (2^(-1) * n * (1 + n)) /. n->5
In:  (((2^(1 * -1) * n) * (1 + n))) /. ((n) -> (5))
Out: 15

> Total[Table[i,{i,1,5}]]
In:  Total[Table[i, {i, 1, 5}]]
Out: 15
```

## Development

To run the tests:
```
go test
```
