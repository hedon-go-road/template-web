package main

import (
	"github.com/hedon954/go-matcher/cmd"
	apihttp "github.com/hedon954/go-matcher/internal/api/httpapi"
)

func main() {
	defer cmd.StopSafe()
	apihttp.SetupHTTPServer()
}
