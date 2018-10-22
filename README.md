# Dual

The `dual` package defines the `Dual` type to represent [dual numbers](https://en.wikipedia.org/wiki/Dual_number). For simple explaination, see the [blog post](http://simonstechblog.blogspot.com/2011/11/dual-number.html) by Simon.

In linear algebra, the dual numbers extend the real numbers by adjoining one new element `ε` with the property `ε^2 = 0` (ε is nilpotent), in a similar way that complex numbers adjoin the imaginary unit `i` with the property `i*i=-1`.

### Install
```bash
go get github.com/tamnd/dual
```

### Usage

Use the `dual.New()` to define the dual number `2 + 1*ɛ`:
```go
x := dual.New(2, 1)
```

Define a function `f(x) = x^2/(x+1)` as:
```go
func f(x dual.Dual) dual.Dual {
   return dual.Div(dual.Mul(x, x), dual.Add(x, dual.FromFloat(1.0))
}
```


If we want to find the first derivative of `f(x)` at `x=2`, i.e, `f'(2)`, we could find it by using dual number arithmetic:

```go
y := f(x)
fmt.Println("f(x) = x^2/(x+1)")
fmt.Println("f(2) = ", dual.Real(y))
fmt.Println("f'(2) = ", dual.Epsilon(y))
```

The complete example:

```go
package main

import "github.com/tamnd/dual"

func f(x dual.Dual) dual.Dual {
   return dual.Div(dual.Mul(x, x), dual.Add(x, dual.FromFloat(1.0))
}

func main() {
    x := dual.New(2, 1)
    y := f(x)
    fmt.Println("f(x) = x^2/(x+1)")
    fmt.Println("f(2) = ", dual.Real(y))
    fmt.Println("f'(2) = ", dual.Epsilon(y))
}
```
