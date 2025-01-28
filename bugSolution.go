```go
package main

import "fmt"

func main() {

    x := 10

    defer func(x int) {
        fmt.Println("defer func x:", x)
    }(x) // Pass the current value to the closure

    x = 5

    fmt.Println("main func x:", x)
}
```