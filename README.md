<p align="center">
<img src="https://raw.githubusercontent.com/franpog859/index/master/logo_index.png">
</p>

[![Documentation](https://godoc.org/github.com/franpog859/index?status.svg)](http://godoc.org/github.com/franpog859/index)
[![Go Report Card](https://goreportcard.com/badge/github.com/franpog859/index)](https://goreportcard.com/report/github.com/franpog859/index)
[![CircleCI](https://circleci.com/gh/franpog859/index.svg?style=shield)](https://circleci.com/gh/franpog859/index)

## Description

Index is a little Go package. It allows you to check on which indexes your item is positioned in the slice. It works with every type of slices! There is a common Go problem to pass the slice with unknown type to the function. It is solved here with the [reflect](https://golang.org/pkg/reflect/) package and [this](https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces-in-go) idea turned out to be a very good solution.

## Usage

To get this package and use it just type in your terminal `go get github.com/franpog859/index`. After that you just can simply import the package in your code `import "github.com/franpog859/index"` and use it as follows:

```go
package main
import (
  "fmt"
  "github.com/franpog859/index"
)

func main() {
  slice := []int{1, 2, 3, 1}
  item := 1
  
  indexes, err := index.GetAll(slice, item)
  if err != nil {
    fmt.Println(err)
  }
    
  fmt.Println(indexes)
}
```
```
Output:
[0 3]
```
Remember not to pass some complex slices like multidimensional slices or slices of maps. This package just does not deal with such complexity so remember to check the `err` value! 

If you want to see examples of usage other functions go to [USAGE.md](https://github.com/franpog859/index/blob/master/USAGE.md) file.
## Contribution

If you want to make this package better just fork this repository and prepare your pull request! Remember to keep the code clean and to test your implementation :wink:

To test this package run in your terminal: 
- `git clone https://github.com/franpog859/index.git`
- `cd index` 
- `go test -cover ./...`

If you see some bug or bad habit feel free to tell me! 
