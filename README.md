# go-betterbig

Tired of writing `new(big.Int)...`? This package provides a set of functions to make working with `big.Int` (and other `math/big` features) easier.

## Installation

```bash
go get -u github.com/slzhffktm/go-betterbig
```

## Why is this (probably) better?

On native big.Int, to do a+b-c:

```go
import "math/big"

func main() {
	a := big.NewInt(1)
	b := big.NewInt(2)
	c := big.NewInt(3)
	res := new(big.Int).Sub( 
		new(big.Int).Add(
		    a,
            b,
        ),
		c,
    )
	fmt.Println("a+b-c:", res)
}
```

On betterbig, it is a lot simplified and easier to read:

```go
import "github.com/slzhffktm/go-betterbig"

func main() {
	a := betterbig.NewIntFromInt(1)
	b := betterbig.NewIntFromInt(2)
	c := betterbig.NewIntFromInt(3)
	res := a.Add(b).Sub(c)
    fmt.Println("a+b-c:", res)
}
```

## Future Works

- Finish up all `big.Int` functions
- Add more `math/big` functionality, e.g. `big.Rat`, `big.Float`
