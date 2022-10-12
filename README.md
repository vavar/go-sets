# go-sets

Simple Set Collection

## Install

```go
go get -u github.com/vavar/go-sets
```

## Usage

```go
package main

import (
 "fmt"

 "github.com/vavar/go-sets"
)

func main() {

 s := sets.Set[string]{}

 s.Add("11")
 s.Add("22")
 s.Add("33")

 s.Remove("22")

 fmt.Println(s.Slice()) // [11 22 33]

}
```
