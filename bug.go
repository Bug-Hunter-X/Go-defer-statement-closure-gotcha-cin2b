```go
func main() {

    var x int = 10

    defer func() {
        fmt.Println("defer func x:", x)
    }()

    x = 5

    fmt.Println("main func x:", x)
}
```