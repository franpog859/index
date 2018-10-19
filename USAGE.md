# Usage of index

To get this package and use it just type in your terminal `go get github.com/franpog859/index`. After that you just can simply import the package in your code `import "github.com/franpog859/index"`. 

Remember not to pass some complex slices like multidimensional slices or slices of maps. This package just does not deal with such complexity so remember to check the `err` value! 

## function index.GetAll(slice, item)
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

## function index.IsAny(slice, item)
```go
package main
import (
  "fmt"
  "github.com/franpog859/index"
)

func main() {
  slice := []int{1, 2, 3, 1}
  item := 1
  
  isAny, err := index.IsAny(slice, item)
  if err != nil {
    fmt.Println(err)
  }
    
  fmt.Println(isAny)
}
```
```
Output:
true
```

## function index.HowMany(slice, item)
```go
package main
import (
  "fmt"
  "github.com/franpog859/index"
)

func main() {
  slice := []int{1, 2, 3, 1}
  item := 1
  
  howMany, err := index.HowMany(slice, item)
  if err != nil {
    fmt.Println(err)
  }
    
  fmt.Println(howMany)
}
```
```
Output:
2
```

There are a lot of similar use cases in the [test](https://github.com/franpog859/index/blob/master/index_test.go) file. Feel free to check it out!