# golang-ip

This is a simple Golang package to access an IP address from an `http.Request` struct.

## Example

```golang
package main

import (
	"fmt"
	"net/http"

	ip "github.com/theTardigrade/golang-ip"
)

func main() {
	r := &http.Request{
		RemoteAddr: "192.158.1.38:80",
	}

	fmt.Println(ip.Get(r))
}
```
