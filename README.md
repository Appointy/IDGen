# idgen

Generates url safe lexically sorted random ids with a prefix in go. Inspired from Stripe.

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