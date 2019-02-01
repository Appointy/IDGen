# idgen

[![Go Report Card](https://goreportcard.com/badge/github.com/appointy/idgen)](https://goreportcard.com/report/github.com/appointy/idgen)
[![Build Status](https://travis-ci.org/appointy/idgen.svg?branch=master)](https://travis-ci.org/appointy/idgen)
[![Coverage](http://gocover.io/_badge/github.com/appointy/idgen)](http://gocover.io/github.com/appointy/idgen)
<a href="https://godoc.org/github.com/appointy/idgen"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a>
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Generates url safe lexically sorted universally unique ids with a prefix in go. Inspired from Stripe.

## Usage

```go
package main

import (
    "fmt"

    "github.com/appointy/idgen"
)

func main()  {
    id := idgen.New("cus")
    fmt.Println(id)
    // Output: cus_0000XSNJG0MQJHBF4QX1EFD6Y3
}
```

### Depends on

[oklog/ulid](https://github.com/oklog/ulid)
