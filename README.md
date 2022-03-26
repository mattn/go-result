# go-result

## Usage

```go
package main

import (
    "fmt"

    . "github.com/mattn/go-result"
)

func doSomething() Result2[int, bool] {
	return Ok2(123, true)
}

func main() {
	r := doSomething()
	v1, v2 := r.Unwrap()
    fmt.Println(v1, v2)
}
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
