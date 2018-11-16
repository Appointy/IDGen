# idgen

[![Go Report Card](https://goreportcard.com/badge/github.com/srikrsna/idgen)](https://goreportcard.com/report/github.com/srikrsna/idgen)
[![Build Status](https://travis-ci.org/srikrsna/idgen.svg?branch=master)](https://travis-ci.org/srikrsna/idgen)
[![Coverage](http://gocover.io/_badge/github.com/srikrsna/idgen)](http://gocover.io/github.com/srikrsna/idgen)
<a href="https://godoc.org/github.com/srikrsna/idgen"><img src="https://img.shields.io/badge/godoc-reference-blue.svg"></a>
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Generates url safe lexically sorted universally unique ids with a prefix in go. Inspired from Stripe.

## Usage

```go
package main

import (
    "fmt"

    "github.com/srikrsna/idgen"
)

func main()  {
    id := idgen.New("cus")
    fmt.Println(id)
    // Output: cus_0000XSNJG0MQJHBF4QX1EFD6Y3
}
```

### Depends on

[oklog/ulid](https://github.com/oklog/ulid)
