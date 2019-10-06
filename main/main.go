package main

import (
	"fmt"
	"net/http"

	"github.com/theTardigrade/ip"
)

func main() {
	r := &http.Request{
		RemoteAddr: "86.14::128.95:50431:43",
	}

	fmt.Println(ip.Get(r))
}
